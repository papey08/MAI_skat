import os
from pymongo import MongoClient
import re

from common.entities import File, Folder
import common.exceptions as exceptions

class MongoAccessor:
    def __init__(self, conn_string: str):
        self.client = MongoClient(conn_string)
        self.db = self.client['storage_db']
        self.storages = self.db['storages']
        self.files = self.db['files']

    def _ensure_storage(self, user_id: int):
        if not self.storages.find_one({'user_id': user_id}):
            self.storages.insert_one({
                'user_id': user_id,
                'root': {'folders': {}, 'files': {}}
            })

    def _get_root(self, user_id: int):
        self._ensure_storage(user_id)
        return self.storages.find_one({'user_id': user_id})

    def _navigate_to_path(self, root, path_parts):
        curr = root
        for part in path_parts:
            if part not in curr['folders']:
                raise exceptions.ObjectNotFoundException()
            curr = curr['folders'][part]
        return curr

    def _update_storage(self, user_id: int, root: dict):
        self.storages.update_one({'user_id': user_id}, {'$set': {'root': root}})

    def create_folder(self, user_id: int, path: str) -> Folder:
        self._ensure_storage(user_id)
        root = self._get_root(user_id)['root']

        path_parts = [part for part in path.strip('/').split('/') if part]
        curr = root

        for part in path_parts[:-1]:
            if part not in curr['folders']:
                raise exceptions.ObjectNotFoundException()
            curr = curr['folders'][part]

        if path_parts[-1] in curr['folders']:
            raise exceptions.ObjectAlreadyExistsException()

        curr['folders'][path_parts[-1]] = {'name': path_parts[-1], 'path': path, 'files': {}, 'folders': {}}

        self._update_storage(user_id, root)
        return Folder(name=path_parts[-1] if path_parts else '', path=path)

    def create_file(self, user_id: int, path: str, content: str = '') -> tuple[File, str]:
        self._ensure_storage(user_id)
        root = self._get_root(user_id)['root']

        *folder_parts, file_name = [part for part in path.strip('/').split('/') if part]
        folder = self._navigate_to_path(root, folder_parts)

        if file_name in folder['files']:
            raise exceptions.ObjectAlreadyExistsException()

        folder['files'][file_name] = {'name': file_name, 'path': path}
        
        file_doc = {
            'name': file_name,
            'path': path,
            'user_id': user_id
        }
        result = self.files.insert_one(file_doc)
        file_id = result.inserted_id
        
        self._update_storage(user_id, root)
        
        return File(name=file_name, path=path), str(file_id)

    def delete_folder(self, user_id: int, path: str):
        root = self._get_root(user_id)['root']
        *parent_parts, folder_name = [part for part in path.strip('/').split('/') if part]
        parent = self._navigate_to_path(root, parent_parts)

        if folder_name not in parent['folders']:
            raise exceptions.ObjectNotFoundException()

        del parent['folders'][folder_name]
        self._update_storage(user_id, root)

    def delete_file(self, user_id: int, path: str):
        root = self._get_root(user_id)['root']
        *folder_parts, file_name = [part for part in path.strip('/').split('/') if part]
        folder = self._navigate_to_path(root, folder_parts)

        if file_name not in folder['files']:
            raise exceptions.ObjectNotFoundException()

        del folder['files'][file_name]
        
        self.files.delete_one({'path': path, 'user_id': user_id})

        self._update_storage(user_id, root)

    def get_folder(self, user_id: int, path: str) -> Folder:
        root = self._get_root(user_id)['root']
        path_parts = [part for part in path.strip('/').split('/') if part]
        folder = self._navigate_to_path(root, path_parts)
        res = Folder(folder['name'] if 'name' in folder else '/', folder['path'] if 'path' in folder else '/')
        res.folders = [f for f in folder['folders']]
        res.files = [f for f in folder['files']]
        return res

    def get_file(self, user_id: int, path: str) -> tuple[File, str]:
        root = self._get_root(user_id)['root']
        *folder_parts, file_name = [part for part in path.strip('/').split('/') if part]
        folder = self._navigate_to_path(root, folder_parts)

        file_data = folder['files'].get(file_name)
        if not file_data:
            raise exceptions.ObjectNotFoundException()

        file_doc = self.files.find_one({'path': path, 'user_id': user_id})
        if not file_doc:
            file_doc = self.files.find_one({'path': f"/{path}", 'user_id': user_id})
            if not file_doc:
                raise exceptions.ObjectNotFoundException()

        return File(name=file_doc['name'], path=file_doc['path']), str(file_doc['_id'])

    def get_files(self, user_id: int, file_name: str, limit: int, offset: int) -> list[File]:
        pattern = re.compile(file_name, re.IGNORECASE)
        cursor = self.files.find(
            {'name': pattern, 'user_id': user_id}
        ).skip(offset).limit(limit)
        
        result = [File(name=file['name'], path=file['path']) for file in cursor]
        return result
