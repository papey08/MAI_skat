from fastapi import FastAPI, Query, Request
from fastapi.responses import JSONResponse
from fastapi.middleware.cors import CORSMiddleware

from contextlib import asynccontextmanager
import uvicorn
import hashlib
import re
import jwt
from jwt.exceptions import ExpiredSignatureError
from typing import List

from auth_accessor import AuthAccessor
from user_accessor import UserAccessor
from core_accessor import CoreAccessor
import http_dto as dto
import common.entities as entities
import common.exceptions as exceptions
import common.codes as codes

class HttpServer:
    PAGINATION_LIMIT = 100

    def __init__(self, auth_nats_url: str, user_nats_url: str, core_nats_url: str, password_salt: str, jwt_secret: str):
        self.auth_accessor = AuthAccessor(auth_nats_url)
        self.user_accessor = UserAccessor(user_nats_url)
        self.core_accessor = CoreAccessor(core_nats_url)
        
        self.password_salt = password_salt
        self.jwt_secret = jwt_secret
        
        self.app = FastAPI(lifespan=self.lifespan)

        self.app.add_middleware(
            CORSMiddleware,
            allow_origins=["*"],
            allow_credentials=True,
            allow_methods=["*"],
            allow_headers=["*"]
        )

        self.app.post('/api/v1/register', response_model=dto.UserResponse)(self.create_user)
        self.app.get('/api/v1/users/{user_id}', response_model=dto.UserResponse)(self.get_user)
        self.app.get('/api/v1/users', response_model=list[dto.UserResponse])(self.get_users)

        self.app.post('/api/v1/login', response_model=dto.TokensResponse)(self.login)
        self.app.post('/api/v1/refresh', response_model=dto.TokensResponse)(self.refresh)

        self.app.post('/api/v1/users/{user_id}/folders/{folder_path:path}', response_model=dto.FolderResponse)(self.create_folder)
        self.app.post('/api/v1/users/{user_id}/files/{file_path:path}', response_model=dto.FileResponse)(self.create_file)
        self.app.delete('/api/v1/users/{user_id}/folders/{folder_path:path}')(self.delete_folder)
        self.app.delete('/api/v1/users/{user_id}/files/{file_path:path}')(self.delete_file)
        self.app.get('/api/v1/users/{user_id}/files/search/{file_name}', response_model=dto.FileSearchResponse)(self.get_files)
        self.app.get('/api/v1/users/{user_id}/folders/{folder_path:path}', response_model=dto.FolderResponse)(self.get_folder)
        self.app.get('/api/v1/users/{user_id}/folders', response_model=dto.FolderResponse)(self.get_root_folder)
        self.app.get('/api/v1/users/{user_id}/files/{file_path:path}', response_model=bytes)(self.get_file)

        self.app.middleware('http')(self.auth_middleware)

        self.setup_exception_handlers()

    @asynccontextmanager
    async def lifespan(self, app: FastAPI):
        await self.auth_accessor.__aenter__()
        await self.user_accessor.__aenter__()
        await self.core_accessor.__aenter__()
        yield
        await self.auth_accessor.__aexit__(None, None, None)
        await self.user_accessor.__aexit__(None, None, None)
        await self.core_accessor.__aexit__(None, None, None)

    async def create_user(self, user: dto.UserCreate) -> dto.UserResponse:
        self._validate_user(user)
        user.password = self._hash_password(user.password)
        user = await self.user_accessor.create_user(user)
        return dto.user_to_response(user)

    async def _check_access(self, user_id: int, actor_id: int) -> bool:
        if user_id == actor_id:
            return True
        _ = await self.user_accessor.get_user_by_id(user_id, actor_id)
        actor = await self.user_accessor.get_user_by_id(actor_id, actor_id)
        return actor.is_admin
    
    async def auth_middleware(self, request: Request, call_next):
        if not request.url.path.startswith('/api/v1/users'):
            return await call_next(request)

        auth_header = request.headers.get('Authorization')
        if auth_header is None or not auth_header.startswith('Bearer '):
            return JSONResponse(status_code=401, content={'detail': 'Unauthorized'})
            
        access_token = auth_header.split(' ')
        if len(access_token) != 2:
            return JSONResponse(status_code=401, content={'detail': 'Unauthorized'})
        
        try:
            payload = jwt.decode(
                access_token[1],
                self.jwt_secret,
                algorithms=['HS256']
            )
            request.state.actor_id = payload.get('user_id')
        except ExpiredSignatureError:
            return JSONResponse(status_code=401, content={'detail': 'Expired token'})
        except Exception:
            return JSONResponse(status_code=401, content={'detail': 'Unauthorized'})

        return await call_next(request)

    async def get_user(self, request: Request, user_id: int) -> dto.UserResponse:
        if not await self._check_access(user_id, request.state.actor_id):
            raise JSONResponse(status_code=403, detail='Forbidden')
        user = await self.user_accessor.get_user_by_id(user_id)
        return dto.user_to_response(user)
        
    async def get_users(self, request: Request, user_id: int,
                        login: str = Query(''), 
                        first_name: str = Query(''), 
                        last_name: str = Query(''),
                        limit: int = Query(PAGINATION_LIMIT),
                        offset: int = Query(0)) -> List[dto.UserResponse]:
        if not await self._check_access(user_id, request.state.actor_id):
            raise JSONResponse(status_code=403, detail='Forbidden')
        users = await self.user_accessor.get_users(login, first_name, last_name, limit, offset)
        return [dto.user_to_response(user) for user in users]
        
    async def login(self, items: dto.LoginRequest) -> dto.TokensResponse:
        items.password = self._hash_password(items.password)
        user_id = await self.user_accessor.login(entities.LoginItems(login=items.login, password=items.password))
        tokens = await self.auth_accessor.create_tokens(user_id)
        return dto.tokens_to_response(tokens)

    async def refresh(self, refresh: dto.RefreshToken) -> dto.TokensResponse:
        tokens = await self.auth_accessor.refresh_tokens(refresh.refresh)
        if tokens == None:
            return JSONResponse(status_code=401, content={'detail': 'expired tokens'})
        return dto.tokens_to_response(tokens)
    
    async def create_folder(self, request: Request, user_id: int, folder_path: str) -> dto.FolderResponse:
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        folder_path = self._fix_path(folder_path)
        folder = await self.core_accessor.create_folder(user_id, folder_path)
        return dto.folder_to_response(folder)
        
    async def create_file(self, request: Request, user_id: int, file_path: str) -> dto.FileResponse:
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        file_path = self._fix_path(file_path)
        content = await request.body()
        file = await self.core_accessor.create_file(user_id, file, content)
        return dto.file_to_response(file)
        
    async def delete_folder(self, request: Request, user_id: int, folder_path: str):
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        folder_path = self._fix_path(folder_path)
        await self.core_accessor.delete_folder(user_id, folder_path)
        return JSONResponse(status_code=200, content={'detail': 'OK'})
        
    async def delete_file(self, request: Request, user_id: int, file_path: str):
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        file_path = self._fix_path(file_path)
        await self.core_accessor.delete_file(user_id, file_path)
        return JSONResponse(status_code=200, content={'detail': 'OK'})

    async def get_file(self, request: Request, user_id: int, file_path: str) -> bytes:
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        file_path = self._fix_path(file_path)
        file = await self.core_accessor.get_file(user_id, file_path)
        return file.content
        
    async def get_folder(self, request: Request, user_id: int, folder_path: str,
                         folder_limit = Query(100),
                         folder_offset = Query(0),
                         file_limit = Query(100),
                         file_offset = Query(0)) -> dto.FolderResponse:
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        folder_path = self._fix_path(folder_path)
        folder = await self.core_accessor.get_folder(user_id, folder_path,
                                                     file_limit, file_offset, 
                                                     folder_limit, folder_offset)
        if folder.path == '':
            folder.path = '/'
        return dto.folder_to_response(folder)
        
    async def get_root_folder(self, request: Request, user_id: int,
                              folder_limit = Query(PAGINATION_LIMIT),
                              folder_offset = Query(0),
                              file_limit = Query(PAGINATION_LIMIT),
                              file_offset = Query(0)) -> dto.FolderResponse:
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        folder = await self.core_accessor.get_folder(user_id, '',
                                                     file_limit, file_offset, 
                                                     folder_limit, folder_offset)
        return dto.folder_to_response(folder)
        
    async def get_files(self, request: Request, user_id: int, file_name: str,
                        limit = Query(PAGINATION_LIMIT), offset = Query(0)) -> dto.FileSearchResponse:
        if not await self._check_access(user_id, request.state.actor_id):
            return JSONResponse(status_code=403, content={'detail': 'Forbidden'})
        files = await self.core_accessor.get_files(user_id, file_name, limit, offset)
        return dto.FileSearchResponse(files=files)

    def run(self, host: str, port: int):
        uvicorn.run(self.app, host=host, port=port)

    def _fix_path(self, path: str):
        if path.startswith('/'):
            return path[1:]
        return path

    def _hash_password(self, password: str) -> str:
        password = self.password_salt + password
        sha256_hash = hashlib.sha256()
        sha256_hash.update(password.encode('utf-8'))
        return sha256_hash.hexdigest()
    
    def _validate_user(self, user: dto.UserCreate):
        if not self._validate_email(user.email):
            return JSONResponse(status_code=400, content={'detail': 'Invalid email'})
        if not self._validate_login(user.login):
            return JSONResponse(status_code=400, content={'detail': 'login max length is 100'})
        if not self._validate_name(user.first_name):
            return JSONResponse(status_code=400, content={'detail': 'first name max length is 50'})
        if not self._validate_name(user.last_name):
            return JSONResponse(status_code=400, content={'detail': 'last name max length is 50'})
        if not self._validate_password(user.password):
            return JSONResponse(
                status_code=400, 
                content={'detail': 'password length must be between 8 and 20 and only latin letters and digits are allowed'})

    def _validate_email(self, email: str) -> bool:
        return bool(re.match(r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$', email))
    
    def _validate_login(self, login: str) -> bool:
        return len(login) <= 100
    
    def _validate_name(self, name: str) -> bool:
        return len(name) <= 50
    
    def _validate_password(self, password: str) -> bool:
        return bool(re.match(r'^[a-zA-Z0-9]{8,20}$', password))

    def setup_exception_handlers(self):
        @self.app.exception_handler(exceptions.UserAlreadyExistsException)
        async def user_already_exists_exception_handler(request, exc):
            return JSONResponse(status_code=409, content={'detail': codes.USER_ALREADY_EXISTS})
        
        @self.app.exception_handler(exceptions.UserNotFoundException)
        async def user_not_found_exception_handler(request, exc):
            return JSONResponse(status_code=404, content={'detail': codes.USER_NOT_FOUND})
        
        @self.app.exception_handler(exceptions.LoginUnableException)
        async def login_unable_exception_handler(request, exc):
            return JSONResponse(status_code=401, content={'detail': codes.LOGIN_UNABLE})
        
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
