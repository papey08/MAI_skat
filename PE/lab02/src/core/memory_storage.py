from common.entities import File, Folder
import common.exceptions as exceptions

class MemoryStorage:
    def __init__(self):
        self.storage = {}

    def create_folder(self, user_id: int, path: str) -> Folder:
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.split('/')
        if paths == [''] or paths == []:
            return exceptions.ObjectInvalidName()
        folder_name = paths[-1]
        current_folder = self.storage[user_id]
        for p in paths[:-1]:
            current_folder = current_folder.folders.get(p)
            if current_folder is None:
                raise exceptions.ObjectNotFoundException()
        if folder_name in current_folder.folders:
            raise exceptions.ObjectAlreadyExistsException()
        new_folder = Folder(name=folder_name, path=path)
        current_folder.folders[folder_name] = new_folder
        return new_folder
    
    def create_file(self, user_id: int, path: str, content: str = '') -> File:
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.strip('/').split('/')
        if paths == [''] or paths == []:
            return exceptions.ObjectInvalidName()
        file_name = paths[-1]
        current_folder = self.storage[user_id]
        for p in paths[:-1]:
            current_folder = current_folder.folders.get(p)
            if current_folder is None:
                raise exceptions.ObjectNotFoundException()
        if file_name in current_folder.files:
            raise exceptions.ObjectAlreadyExistsException()
        new_file = File(name=file_name, path=path, content=content)
        current_folder.files[file_name] = new_file
        return new_file

    def delete_folder(self, user_id: int, path: str):
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.split('/')
        if paths == ['']:
            paths = []
        current_folder = self.storage[user_id]
        for p in paths[:-1]:
            current_folder = current_folder.folders.get(p)
            if current_folder is None:
                raise exceptions.ObjectNotFoundException()
        if current_folder.folders.get(paths[-1]) is None:
            raise exceptions.ObjectNotFoundException()
        del current_folder.folders[paths[-1]]

    def delete_file(self, user_id: int, path: str):
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.split('/')
        if paths == ['']:
            paths = []
        current_folder = self.storage[user_id]
        for p in paths[:-1]:
            current_folder = current_folder.folders.get(p)
            if current_folder is None:
                raise exceptions.ObjectNotFoundException()
        if current_folder.files[paths[-1]] is None:
            raise exceptions.ObjectNotFoundException()
        del current_folder.files[paths[-1]]

    def get_folder(self, user_id: int, path: str) -> Folder:
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.split('/')
        if paths == ['']:
            paths = []
        current_folder = self.storage[user_id]
        for p in paths:
            current_folder = current_folder.folders.get(p)
            if current_folder is None:
                raise exceptions.ObjectNotFoundException()
        res = Folder(
            name=current_folder.name,
            path=current_folder.path
        )
        res.files = current_folder.files
        res.folders = current_folder.folders
        return res
    
    def get_file(self, user_id: int, path: str) -> File:
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.split('/')
        if paths == ['']:
            paths = []
        current_folder = self.storage[user_id]
        for p in paths[:-1]:
            current_folder = current_folder.folders.get(p)
            if current_folder is None:
                raise exceptions.ObjectNotFoundException()
        if current_folder.files.get(paths[-1]) is None:
            raise exceptions.ObjectNotFoundException()
        return current_folder.files[paths[-1]]
    
    def get_files(self, user_id: int, file_name: str, limit: int, offset: int) -> list[File]:
        self._init_storage_for_user(user_id)

        res = []
        bfs = [self.storage[user_id]]
        for folder in bfs:
            for file in folder.files:
                if file_name.lower() in file.lower():
                    res.append(folder.files[file])
            bfs.extend(folder.folders.values())
        return res[
            offset :
            offset + limit
        ]

    def _init_storage_for_user(self, user_id: int):
        if user_id not in self.storage:
            self.storage[user_id] = Folder(name='/', path='')

    def _validate_path(self, path: str):
        if path == '':
            return
        if path.strip().strip('/') == '' or path.strip('/').strip() == '':
            raise exceptions.ObjectInvalidName()
        if '' in [p.strip() for p in path.strip('/').split('/')]:
            raise exceptions.ObjectInvalidName()
