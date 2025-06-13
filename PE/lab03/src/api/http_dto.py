from pydantic import BaseModel
from datetime import datetime
from typing import List

import common.entities as entities

class UserCreate(BaseModel):
    login: str
    first_name: str
    last_name: str
    email: str
    password: str

class UserResponse(BaseModel):
    id: int
    first_name: str
    last_name: str
    login: str
    email: str
    is_admin: bool
    created_at: datetime

def user_to_response(user: entities.User) -> UserResponse:
    return UserResponse(
        id=user.id,
        first_name=user.first_name,
        last_name=user.last_name,
        login=user.login,
        email=user.email,
        is_admin=user.is_admin,
        created_at=user.created_at
    )

class LoginRequest(BaseModel):
    login: str
    password: str

class RefreshToken(BaseModel):
    refresh: str

class TokensResponse(BaseModel):
    access: str
    refresh: str

def tokens_to_response(tokens: entities.Tokens) -> TokensResponse:
    return TokensResponse(access=tokens.access, refresh=tokens.refresh)

class FolderResponse(BaseModel):
    name: str
    path: str = '/'
    files: List[str]
    folders: List[str]

def folder_to_response(folder: entities.Folder) -> FolderResponse:
    return FolderResponse(
        name=folder.name,
        path=folder.path,
        files=list(folder.files.keys()),
        folders=list(folder.folders.keys())
    )

def folder_to_response(folder: entities.Folder) -> FolderResponse:
    return FolderResponse(
        name=folder.name,
        path=folder.path,
        files=list(folder.files.keys()),
        folders=list(folder.folders.keys())
    )

class FileResponse(BaseModel):
    name: str
    path: str = '/'

def file_to_response(file: entities.File) -> FileResponse:
    return FileResponse(
        name=file.name,
        path=file.path
    )

class FileSearchResponse(BaseModel):
    files: List[str] = None
