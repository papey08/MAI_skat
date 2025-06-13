from fastapi import FastAPI, Query, Depends, File, UploadFile
from fastapi.responses import JSONResponse, StreamingResponse
from fastapi.security import HTTPBearer
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
import hashlib
import jwt
import re
from typing import List
import io

from auth_accessor import AuthAccessor
from user_accessor import UserAccessor
from core_accessor import CoreAccessor
import http_dto as dto
import common.entities as entities
import common.exceptions as exceptions
import common.codes as codes

class HttpServer():
    PAGINATION_LIMIT = 100

    def __init__(self, auth_nats_url: str, user_nats_url: str, core_nats_url: str, password_salt: str, jwt_secret: str):
        self.auth_accessor = AuthAccessor(auth_nats_url)
        self.user_accessor = UserAccessor(user_nats_url)
        self.core_accessor = CoreAccessor(core_nats_url)
        
        self.password_salt = password_salt
        self.jwt_secret = jwt_secret

        async def lifespan(app: FastAPI):
            await self.auth_accessor.__aenter__()
            await self.user_accessor.__aenter__()
            await self.core_accessor.__aenter__()
            yield
            await self.auth_accessor.__aexit__(None, None, None)
            await self.user_accessor.__aexit__(None, None, None)
            await self.core_accessor.__aexit__(None, None, None)

        self.app = FastAPI(
            lifespan=lifespan,
            openapi_url='/api/v1/openapi.json',
            title='API приложения для хранения файлов',
            root_path='http://localhost:8099/',
            docs_url='/api/v1/docs', 
            redoc_url='/api/v1/redoc'
        )

        self.app.add_middleware(
            CORSMiddleware,
            allow_origins=["*"],
            allow_credentials=True,
            allow_methods=["*"],
            allow_headers=["*"]
        )
        
        self.auth_scheme = HTTPBearer()

        @self.app.post('/api/v1/register', tags=['user'], response_model=dto.UserResponse, responses={
            409: {'detail': codes.USER_NOT_FOUND},
            400: {'detail': codes.USER_INVALID},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def create_user(user: dto.UserCreate) -> dto.UserResponse:
            self._validate_user(user)
            user.password = self._hash_password(user.password)
            user = await self.user_accessor.create_user(user)
            return dto.user_to_response(user)
        
        @self.app.get('/api/v1/users/{user_id}', tags=['user'], response_model=dto.UserResponse, responses={
            401: {'detail': codes.UNAUTHORIZED},
            404: {'detail': codes.USER_NOT_FOUND},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def get_user(user_id: int, token = Depends(self.auth_scheme)) -> dto.UserResponse:
            self._verify_token(token.credentials)
            user = await self.user_accessor.get_user_by_id(user_id)
            return dto.user_to_response(user)
        
        @self.app.get('/api/v1/users', tags=['user'], response_model=list[dto.UserResponse], responses={
            401: {'detail': codes.UNAUTHORIZED},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def get_users(token = Depends(self.auth_scheme),
                            login: str = Query(''), 
                            first_name: str = Query(''), 
                            last_name: str = Query(''),
                            limit: int = Query(self.PAGINATION_LIMIT),
                            offset: int = Query(0)) -> List[dto.UserResponse]:
            self._verify_token(token.credentials)
            limit = max(0, min(limit, self.PAGINATION_LIMIT))
            users = await self.user_accessor.get_users(login, first_name, last_name, limit, offset)
            return [dto.user_to_response(user) for user in users]
        
        @self.app.post('/api/v1/login', tags=['auth'], response_model=dto.TokensResponse, responses={
            401: {'detail': codes.LOGIN_UNABLE},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def login(items: dto.LoginRequest) -> dto.TokensResponse:
            items.password = self._hash_password(items.password)
            user_id = await self.user_accessor.login(entities.LoginItems(login=items.login, password=items.password))
            tokens = await self.auth_accessor.create_tokens(user_id)
            return dto.tokens_to_response(tokens)
        
        @self.app.post('/api/v1/refresh', tags=['auth'], response_model=dto.TokensResponse, responses={
            401: {'detail': codes.EXPIRED_REFRESH_TOKEN},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def refresh(token: dto.RefreshToken) -> dto.TokensResponse:
            tokens = await self.auth_accessor.refresh_tokens(token.refresh)
            return dto.tokens_to_response(tokens)
        
        @self.app.post('/api/v1/users/{user_id}/folders/{folder_path:path}', tags=['core'], response_model=dto.FolderResponse, responses={
            404: {'detail': codes.OBJECT_NOT_FOUND},
            409: {'detail': codes.OBJECT_ALREADY_EXISTS},
            401: {'detail': codes.UNAUTHORIZED},
            400: {'detail': codes.OBJECT_INVALID_NAME},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def create_folder(user_id: int, folder_path: str, token = Depends(self.auth_scheme)) -> dto.FolderResponse:
            await self._check_strict_access(user_id, token.credentials)
            folder_path = self._fix_path(folder_path)
            folder = await self.core_accessor.create_folder(user_id, folder_path)
            return dto.folder_to_response(folder)
        
        @self.app.post('/api/v1/users/{user_id}/files/{file_path:path}', tags=['core'], response_model=dto.FileResponse, responses={
            404: {'detail': codes.OBJECT_NOT_FOUND},
            409: {'detail': codes.OBJECT_ALREADY_EXISTS},
            401: {'detail': codes.UNAUTHORIZED},
            400: {'detail': codes.OBJECT_INVALID_NAME},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def create_file(user_id: int, file_path: str, file: UploadFile = File(...), token = Depends(self.auth_scheme)) -> dto.FileResponse:
            await self._check_strict_access(user_id, token.credentials)
            file_path = '/'.join([file_path, file.filename])
            file_path = self._fix_path(file_path)
            content = await file.read()
            file = await self.core_accessor.create_file(user_id, file_path, content)
            return dto.file_to_response(file)
        
        @self.app.post('/api/v1/users/{user_id}/files', tags=['core'], response_model=dto.FileResponse, responses={
            404: {'detail': codes.OBJECT_NOT_FOUND},
            409: {'detail': codes.OBJECT_ALREADY_EXISTS},
            401: {'detail': codes.UNAUTHORIZED},
            400: {'detail': codes.OBJECT_INVALID_NAME},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def create_file_in_root_folder(user_id: int, file: UploadFile = File(...), token = Depends(self.auth_scheme)) -> dto.FileResponse:
            await self._check_strict_access(user_id, token.credentials)
            content = await file.read()
            file = await self.core_accessor.create_file(user_id, '/' + file.filename, content)
            return dto.file_to_response(file)

        @self.app.get('/api/v1/users/{user_id}/folders/{folder_path:path}', tags=['core'], response_model=dto.FolderResponse, responses={
            404: {'detail': codes.OBJECT_NOT_FOUND},
            401: {'detail': codes.UNAUTHORIZED},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def get_folder(user_id: int, folder_path: str,
                            folder_limit: int = Query(self.PAGINATION_LIMIT),
                            folder_offset: int = Query(0),
                            file_limit: int = Query(self.PAGINATION_LIMIT),
                            file_offset: int = Query(0), token = Depends(self.auth_scheme)) -> dto.FolderResponse:
            await self._check_strict_access(user_id, token.credentials)
            folder_path = self._fix_path(folder_path)
            folder_limit = max(0, min(self.PAGINATION_LIMIT, folder_limit))
            file_limit = max(0, min(self.PAGINATION_LIMIT, file_limit))
            folder = await self.core_accessor.get_folder(user_id, folder_path, file_limit, file_offset, folder_limit, folder_offset)
            if folder.path == '':
                folder.path = '/'
            return dto.folder_to_response(folder)

        @self.app.get('/api/v1/users/{user_id}/folders', tags=['core'], response_model=dto.FolderResponse, responses={
            401: {'detail': codes.UNAUTHORIZED},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def get_root_folder(user_id: int, token = Depends(self.auth_scheme),
                            folder_limit: int = Query(self.PAGINATION_LIMIT),
                            folder_offset: int = Query(0),
                            file_limit: int = Query(self.PAGINATION_LIMIT),
                            file_offset: int = Query(0)) -> dto.FolderResponse:
            await self._check_strict_access(user_id, token.credentials)
            folder_limit = max(0, min(self.PAGINATION_LIMIT, folder_limit))
            file_limit = max(0, min(self.PAGINATION_LIMIT, file_limit))
            folder = await self.core_accessor.get_folder(user_id, '', file_limit, file_offset, folder_limit, folder_offset)
            return dto.folder_to_response(folder)

        @self.app.get('/api/v1/users/{user_id}/files/{file_path:path}', tags=['core'], responses={
            404: {'detail': codes.OBJECT_NOT_FOUND},
            401: {'detail': codes.UNAUTHORIZED},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def get_file(user_id: int, file_path: str, token = Depends(self.auth_scheme)) -> StreamingResponse:
            await self._check_strict_access(user_id, token.credentials)
            file_path = self._fix_path(file_path)
            file = await self.core_accessor.get_file(user_id, file_path)
            return StreamingResponse(
                io.BytesIO(file.content),
                media_type="application/octet-stream",
            )
        
        @self.app.get('/api/v1/users/{user_id}/files_search', tags=['core'], response_model=dto.FileSearchResponse, responses={
            401: {'detail': codes.UNAUTHORIZED},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def get_files(user_id: int, token = Depends(self.auth_scheme), pattern: str = Query(''),
                            limit: int = Query(self.PAGINATION_LIMIT), offset: int = Query(0)) -> List[str]:
            await self._check_strict_access(user_id, token.credentials)
            limit = max(0, min(self.PAGINATION_LIMIT, limit))
            files = await self.core_accessor.get_files(user_id, pattern, limit, offset)
            return dto.FileSearchResponse(files=files)
        
        @self.app.delete('/api/v1/users/{user_id}/folders/{folder_path:path}', tags=['core'], responses={
            200: {'detail': 'OK'},
            404: {'detail': codes.OBJECT_NOT_FOUND},
            401: {'detail': codes.UNAUTHORIZED},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def delete_folder(user_id: int, folder_path: str, token = Depends(self.auth_scheme)) -> dto.FolderResponse:
            await self._check_strict_access(user_id, token.credentials)
            folder_path = self._fix_path(folder_path)
            await self.core_accessor.delete_folder(user_id, folder_path)
            return JSONResponse(status_code=200, content={'detail': 'OK'})
        
        @self.app.delete('/api/v1/users/{user_id}/files/{file_path:path}', tags=['core'], responses={
            200: {'detail': 'OK'},
            404: {'detail': codes.OBJECT_NOT_FOUND},
            401: {'detail': codes.UNAUTHORIZED},
            403: {'detail': codes.ACCESS_DENIED},
            500: {'detail': codes.INTERNAL_ERROR}
        })
        async def delete_file(user_id: int, file_path: str, token = Depends(self.auth_scheme)) -> dto.FileResponse:
            await self._check_strict_access(user_id, token.credentials)
            file_path = self._fix_path(file_path)
            await self.core_accessor.delete_file(user_id, file_path)
            return JSONResponse(status_code=200, content={'detail': 'OK'})
        

        @self.app.exception_handler(exceptions.UserAlreadyExistsException)
        async def user_already_exists_exception_handler(request, exc):
            return JSONResponse(status_code=409, content={'detail': codes.USER_ALREADY_EXISTS})
        
        @self.app.exception_handler(exceptions.UserInvalidException)
        async def user_invalid_exception_handler(request, exc):
            return JSONResponse(status_code=400, content={'detail': codes.USER_INVALID})
        
        @self.app.exception_handler(exceptions.UserNotFoundException)
        async def user_not_found_exception_handler(request, exc):
            return JSONResponse(status_code=404, content={'detail': codes.USER_NOT_FOUND})
        
        @self.app.exception_handler(exceptions.LoginUnableException)
        async def login_unable_exception_handler(request, exc):
            return JSONResponse(status_code=401, content={'detail': codes.LOGIN_UNABLE})
        
        @self.app.exception_handler(exceptions.UnauthorizedException)
        async def unauthorized_exception_handler(request, exc):
            return JSONResponse(status_code=401, content={'detail': codes.UNAUTHORIZED})

        @self.app.exception_handler(exceptions.ExpiredRefreshTokenException)
        async def expired_refresh_token_exception_handler(request, exc):
            return JSONResponse(status_code=401, content={'detail': codes.EXPIRED_REFRESH_TOKEN})
        
        @self.app.exception_handler(exceptions.ObjectNotFoundException)
        async def object_not_found_exception_handler(request, exc):
            return JSONResponse(status_code=404, content={'detail': codes.OBJECT_NOT_FOUND})
        
        @self.app.exception_handler(exceptions.ObjectAlreadyExistsException)
        async def object_already_exists_exception_handler(request, exc):
            return JSONResponse(status_code=400, content={'detail': codes.OBJECT_ALREADY_EXISTS})
        
        @self.app.exception_handler(exceptions.ObjectInvalidName)
        async def object_invalid_name_exception_handler(request, exc):
            return JSONResponse(status_code=400, content={'detail': codes.OBJECT_INVALID_NAME})
        
        @self.app.exception_handler(exceptions.AccessDeniedException)
        async def access_denied_exception_handler(request, exc):
            return JSONResponse(status_code=403, content={'detail': codes.ACCESS_DENIED})
        
        @self.app.exception_handler(exceptions.InternalException)
        async def internal_exception_handler(request, exc):
            return JSONResponse(status_code=500, content={'detail': codes.INTERNAL_ERROR})
        
        @self.app.exception_handler(RuntimeError)
        async def runtime_error_exception_handler(request, exc):
            return JSONResponse(status_code=500, content={'detail': codes.INTERNAL_ERROR})
        
        @self.app.exception_handler(TimeoutError)
        async def timeout_error_exception_handler(request, exc):
            return JSONResponse(status_code=500, content={'detail': codes.INTERNAL_ERROR})

    def run(self, host: str, port: int):
        uvicorn.run(self.app, host=host, port=port)
    
    def _verify_token(self, token: str) -> int:
        try:
            payload = jwt.decode(token, self.jwt_secret, algorithms=['HS256'])
            return int(payload['user_id'])
        except jwt.PyJWTError:
            raise exceptions.UnauthorizedException()

    async def _check_strict_access(self, user_id: int, token: str):
        _ = await self.user_accessor.get_user_by_id(user_id)
        actor_id = self._verify_token(token)
        if user_id == actor_id:
            return
        actor = await self.user_accessor.get_user_by_id(actor_id)
        if not actor.is_admin:
            raise exceptions.AccessDeniedException()

    async def _check_access(self, user_id: int, token: str):
        _ = await self.user_accessor.get_user_by_id(user_id)
        actor_id = self._verify_token(token)
        _ = await self.user_accessor.get_user_by_id(actor_id)
        return

    def _hash_password(self, password: str) -> str:
        password = self.password_salt + password
        sha256_hash = hashlib.sha256()
        sha256_hash.update(password.encode('utf-8'))
        return sha256_hash.hexdigest()

    def _validate_user(self, user: dto.UserCreate):
        if not self._validate_email(user.email) or \
            not self._validate_login(user.login) or \
            not self._validate_name(user.first_name) or \
            not self._validate_name(user.last_name) or \
            not self._validate_password(user.password):
            raise exceptions.UserInvalidException()
        
    def _validate_email(self, email: str) -> bool:
        return bool(re.match(r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$', email))
        
    def _validate_login(self, login: str) -> bool:
        return len(login) > 0 and len(login) <= 100

    def _validate_name(self, name: str) -> bool:
        return len(name) > 0 and len(name) <= 50

    def _validate_password(self, password: str) -> bool:
        return bool(re.match(r'^[a-zA-Z0-9]{8,20}$', password))

    def _fix_path(self, path: str) -> str:
        if path.startswith('/'):
            return path[1:]
        return path
