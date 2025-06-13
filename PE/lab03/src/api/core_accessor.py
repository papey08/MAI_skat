import nats
from nats.aio.client import Client as NATSClient
from typing import Optional
from pydantic import ValidationError

import common.entities as entities
import common.dto as dto
import common.exceptions as exceptions
import common.codes as codes

class CoreAccessor:
    def __init__(self, nats_url: str):
        self.nats_url = nats_url
        self.nc: Optional[NATSClient] = None

    async def __aenter__(self):
        self.nc = await nats.connect(self.nats_url)
        return self
    
    async def __aexit__(self, exc_type, exc_val, exc_tb):
        if self.nc:
            await self.nc.close()

    async def handle_error(self, data: bytes):
        err = dto.ErrorResponse.model_validate_json(data)
        if err.code == codes.OBJECT_ALREADY_EXISTS:
            raise exceptions.ObjectAlreadyExistsException()
        if err.code == codes.OBJECT_NOT_FOUND:
            raise exceptions.ObjectNotFoundException()
        if err.code == codes.OBJECT_INVALID_NAME:
            raise exceptions.ObjectInvalidName()
        if err.code == codes.INTERNAL_ERROR:
            raise exceptions.InternalException()

    async def create_folder(self, user_id: int, path: str) -> entities.Folder:
        try:
            msg = dto.CreateFolderMessage(
                user_id=user_id,
                path=path
            )
            response = await self.nc.request('create_folder', msg.model_dump_json().encode(), timeout=30)
            res = dto.CreateFolderResponse.model_validate_json(response.data.decode())
            return entities.Folder(
                name=res.name,
                path=res.path
            )
        except ValidationError:
            await self.handle_error(response.data)
            
    async def create_file(self, user_id: int, path: str, content: bytes) -> entities.File:
        try:
            msg = dto.CreateFileMessage(
                user_id=user_id,
                path=path,
                content=content,
            )
            response = await self.nc.request('create_file', msg.model_dump_json().encode(), timeout=30)
            res = dto.CreateFileResponse.model_validate_json(response.data.decode())
            return entities.File(
                name=res.name,
                path=res.path
            )
        except ValidationError:
            await self.handle_error(response.data.decode())
            
    async def delete_folder(self, user_id: int, path: str) -> None:
        try:
            msg = dto.DeleteFolderMessage(
                user_id=user_id,
                path=path
            )
            response = await self.nc.request('delete_folder', msg.model_dump_json().encode(), timeout=30)
            _ = dto.DeleteFolderResponse.model_validate_json(response.data.decode())
            return None
        except ValidationError:
            await self.handle_error(response.data.decode())

    async def delete_file(self, user_id: int, path: str) -> None:
        try:
            msg = dto.DeleteFileMessage(
                user_id=user_id,
                path=path
            )
            response = await self.nc.request('delete_file', msg.model_dump_json().encode(), timeout=30)
            dto.DeleteFileResponse.model_validate_json(response.data.decode())
            return None
        except ValidationError:
            await self.handle_error(response.data.decode())
            
    async def get_file(self, user_id: int, path: str) -> entities.File:
        try:
            msg = dto.GetFileMessage(
                user_id=user_id,
                path=path,
            )
            response = await self.nc.request('get_file', msg.model_dump_json().encode(), timeout=30)
            res = dto.GetFileResponse.model_validate_json(response.data.decode())
            return entities.File(
                name=res.name,
                path=res.path,
                content=res.content)
        except ValidationError:
            await self.handle_error(response.data.decode())

    async def get_files(self, user_id: int, file_name: str, limit: int, offset: int) -> list[str]:
        try:
            msg = dto.GetFilesMessage(
                user_id=user_id,
                file_name=file_name,
                limit=limit,
                offset=offset
            )
            response = await self.nc.request('get_files', msg.model_dump_json().encode(), timeout=30)
            res = dto.GetFilesResponse.model_validate_json(response.data.decode())
            return res.files
        except ValidationError:
            await self.handle_error(response.data.decode())
            
    async def get_folder(self, user_id: int, path: str,
                         file_limit: int, file_offset: int, 
                         folder_limit: int, folder_offset: int) -> entities.Folder:
        try:
            msg = dto.GetFolderMessage(
                user_id=user_id,
                path=path,
                file_limit=file_limit,
                file_offset=file_offset,
                folder_limit=folder_limit,
                folder_offset=folder_offset
            )
            response = await self.nc.request('get_folder', msg.model_dump_json().encode(), timeout=30)
            res = dto.GetFolderResponse.model_validate_json(response.data.decode())
            folder = entities.Folder(
                name=res.name,
                path=res.path
            )
            for f in res.files:
                folder.files[f] = None
            for f in res.folders:
                folder.folders[f] = None
            return folder
        except ValidationError:
            await self.handle_error(response.data.decode())
