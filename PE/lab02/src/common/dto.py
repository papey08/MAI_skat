from pydantic import BaseModel
from datetime import datetime

class CreateUserMessage(BaseModel):
    login: str
    first_name: str
    last_name: str
    email: str
    password: str

class CreateUserResponse(BaseModel):
    id: int
    first_name: str
    last_name: str
    login: str
    email: str
    is_admin: bool
    created_at: datetime

class GetUserByIdMessage(BaseModel):
    user_id: int

class GetUserByIdResponse(BaseModel):
    id: int
    first_name: str
    last_name: str
    login: str
    email: str
    is_admin: bool
    created_at: datetime

class GetUsersMessage(BaseModel):
    login: str
    first_name: str
    last_name: str
    limit: int
    offset: int

class GetUsersResponse(BaseModel):
    Users: list[GetUserByIdResponse]

class LoginMessage(BaseModel):
    login: str
    password: str

class LoginResponse(BaseModel):
    id: int

class CreateTokensMessage(BaseModel):
    user_id: int

class CreateTokensResponse(BaseModel):
    access: str
    refresh: str

class RefreshTokensMessage(BaseModel):
    refresh: str

class RefreshTokensResponse(BaseModel):
    access: str
    refresh: str

class CreateFolderMessage(BaseModel):
    user_id: int
    path: str

class CreateFolderResponse(BaseModel):
    name: str
    path: str

class CreateFileMessage(BaseModel):
    user_id: int
    path: str
    
    content: bytes

class CreateFileResponse(BaseModel):
    name: str
    path: str

class DeleteFileMessage(BaseModel):
    user_id: int
    path: str

class DeleteFileResponse(BaseModel):
    result: str = 'OK'

class DeleteFolderMessage(BaseModel):
    user_id: int
    path: str

class DeleteFolderResponse(BaseModel):
    result: str = 'OK'

class GetFileMessage(BaseModel):
    user_id: int
    path: str

class GetFileResponse(BaseModel):
    name: str
    path: str
    content: bytes

class GetFolderMessage(BaseModel):
    user_id: int
    path: str
    
    file_limit: int
    file_offset: int
    folder_limit: int
    folder_offset: int

class GetFolderResponse(BaseModel):
    name: str
    path: str
    folders: list[str] = None
    files: list[str] = None

class GetFilesMessage(BaseModel):
    user_id: int
    file_name: str

    limit: int
    offset: int

class GetFilesResponse(BaseModel):
    files: list[str]

class ErrorResponse(BaseModel):
    code: str
