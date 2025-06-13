import nats
import asyncio
from typing import Optional

from minio_accessor import MinioAccessor
from common.entities import SubPagination
import common.exceptions as exceptions 
import common.dto as dto
import common.codes as codes

import logging
logging.basicConfig(
    level=logging.INFO
)

class NatsServer:
    def __init__(self, nats_url: str, endpoint: str, access_key: str, secret_key: str):
        self.nc: Optional[nats.NATS] = None
        self.nats_url = nats_url

        self.memory_accessor = MinioAccessor(endpoint, access_key, secret_key)

    async def connect(self):
        self.nc = await nats.connect(self.nats_url)

    def exception_handler(method):
        async def wrapper(self, msg):
            try:
                await method(self, msg)
            except exceptions.ObjectNotFoundException as e:
                logging.error(f'{e}')
                error = dto.ErrorResponse(
                    code=codes.OBJECT_NOT_FOUND
                )
                await msg.respond(error.model_dump_json().encode())
            except exceptions.ObjectAlreadyExistsException as e:
                logging.error(f'{e}')
                error = dto.ErrorResponse(
                    code=codes.OBJECT_ALREADY_EXISTS
                )
                await msg.respond(error.model_dump_json().encode())
            except exceptions.ObjectInvalidName as e:
                logging.error(f'{e}')
                error = dto.ErrorResponse(
                    code=codes.OBJECT_INVALID_NAME
                )
                await msg.respond(error.model_dump_json().encode())
            except exceptions.InternalException as e:
                logging.error(f'{e}')
                error = dto.ErrorResponse(
                    code=codes.INTERNAL_ERROR
                )
                await msg.respond(error.model_dump_json().encode())
            except RuntimeError or TimeoutError as e:
                logging.error(f'{e}')
                error = dto.ErrorResponse(
                    code=codes.INTERNAL_ERROR
                )
                await msg.respond(error.model_dump_json().encode())
        return wrapper

    @exception_handler
    async def handle_create_folder(self, msg):
        create_folder = dto.CreateFolderMessage.model_validate_json(msg.data)

        logging.info(f'creating folder {create_folder.path} for user {create_folder.user_id}')
        folder = self.memory_accessor.create_folder(create_folder.user_id, create_folder.path)
        logging.info(f'created folder {create_folder.path} for user {create_folder.user_id}')
        resp = dto.CreateFolderResponse(
            name=folder.name,
            path=folder.path
        ).model_dump_json()
        await msg.respond(resp.encode())

    @exception_handler
    async def handle_create_file(self, msg):
        create_file = dto.CreateFileMessage.model_validate_json(msg.data)

        logging.info(f'creating file {create_file.path} for user {create_file.user_id}')
        file = self.memory_accessor.create_file(create_file.user_id, create_file.path, create_file.content)
        logging.info(f'created file {create_file.path} for user {create_file.user_id}')
        await msg.respond(dto.CreateFileResponse(
            name=file.name,
            path=file.path
        ).model_dump_json().encode())
        
    @exception_handler
    async def handle_delete_file(self, msg):
        delete_file = dto.DeleteFileMessage.model_validate_json(msg.data)

        logging.info(f'deleting file {delete_file.path} for user {delete_file.user_id}')
        self.memory_accessor.delete_file(delete_file.user_id, delete_file.path)
        logging.info(f'deleted file {delete_file.path} for user {delete_file.user_id}')
        await msg.respond(dto.DeleteFileResponse().model_dump_json().encode())

    @exception_handler
    async def handle_delete_folder(self, msg):
        delete_folder = dto.DeleteFolderMessage.model_validate_json(msg.data)

        logging.info(f'deleting folder {delete_folder.path} for user {delete_folder.user_id}')
        self.memory_accessor.delete_folder(delete_folder.user_id, delete_folder.path)
        logging.info(f'deleted folder {delete_folder.path} for user {delete_folder.user_id}')
        await msg.respond(dto.DeleteFolderResponse().model_dump_json().encode())

    @exception_handler
    async def handle_get_file(self, msg):
        get_file = dto.GetFileMessage.model_validate_json(msg.data)

        logging.info(f'getting file {get_file.path} for user {get_file.user_id}')
        file = self.memory_accessor.get_file(get_file.user_id, get_file.path)
        logging.info(f'got file {get_file.path} for user {get_file.user_id}')
        await msg.respond(dto.GetFileResponse(
            name=file.name,
            path=file.path,
            content=file.content
        ).model_dump_json().encode())

    @exception_handler
    async def handle_get_files(self, msg):
        get_files = dto.GetFilesMessage.model_validate_json(msg.data)

        logging.info(f'getting files for user {get_files.user_id}')
        files = self.memory_accessor.get_files(get_files.user_id, get_files.file_name, get_files.limit, get_files.offset)
        logging.info(f'got files for user {get_files.user_id}')
        await msg.respond(dto.GetFilesResponse(
            files=[f.path for f in files]
        ).model_dump_json().encode())

    @exception_handler
    async def handle_get_folder(self, msg):
        get_folder = dto.GetFolderMessage.model_validate_json(msg.data)
        pagination = SubPagination(
            folder_limit=get_folder.folder_limit,
            folder_offset=get_folder.folder_offset,
            file_limit=get_folder.file_limit,
            file_offset=get_folder.file_offset
        )

        logging.info(f'getting folder {get_folder.path} for user {get_folder.user_id}')
        folder = self.memory_accessor.get_folder(get_folder.user_id, get_folder.path)
        logging.info(f'got folder {get_folder.path} for user {get_folder.user_id}')
        res = dto.GetFolderResponse(
            name=folder.name,
            path=folder.path,
            files=[f.name for f in folder.files.values()],
            folders=[f.name for f in folder.folders.values()]
        )
        res.files = res.files[pagination.file_offset : 
                pagination.file_offset + pagination.file_limit]
        res.folders = res.folders[pagination.folder_offset : 
                pagination.folder_offset + pagination.folder_limit]
        await msg.respond(res.model_dump_json().encode())

    async def run(self):
        await self.connect()
        await self.nc.subscribe('create_folder', cb=self.handle_create_folder)
        await self.nc.subscribe('create_file', cb=self.handle_create_file)
        await self.nc.subscribe('delete_file', cb=self.handle_delete_file)
        await self.nc.subscribe('delete_folder', cb=self.handle_delete_folder)
        await self.nc.subscribe('get_file', cb=self.handle_get_file)
        await self.nc.subscribe('get_files', cb=self.handle_get_files)
        await self.nc.subscribe('get_folder', cb=self.handle_get_folder)

        logging.info('started nats server')

        try:
            await asyncio.Future()
        except asyncio.CancelledError:
            pass
        finally:
            await self.nc.close()
