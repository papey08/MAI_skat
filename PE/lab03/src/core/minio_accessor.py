from minio import Minio
from minio.error import S3Error
import io

from common.entities import File, Folder
import common.exceptions as exceptions

class MinioAccessor:
    def __init__(self, endpoint: str, access_key: str, secret_key: str):
        self.client = Minio(endpoint, access_key=access_key, secret_key=secret_key, secure=False)

    def create_folder(self, user_id: int, path: str) -> Folder:
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.split('/')
        if paths == [''] or paths == []:
            return exceptions.ObjectInvalidName()

        bucket_name = self._to_padded_string(user_id)
        objects = list(self.client.list_objects(bucket_name, prefix=f"{path}/", recursive=False))
        if objects:
            raise exceptions.ObjectAlreadyExistsException()
        self.client.put_object(bucket_name, f"{path}/", io.BytesIO(b""), 0, content_type="application/x-directory")
        return Folder(name=path.split('/')[-1], path=path)

    def create_file(self, user_id: int, path: str, content: bytes = b'') -> File:
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        paths = path.strip('/').split('/')
        if paths == [''] or paths == []:
            return exceptions.ObjectInvalidName()

        bucket_name = self._to_padded_string(user_id)
        try:
            self.client.stat_object(bucket_name, path)
            raise exceptions.ObjectAlreadyExistsException()
        except S3Error as e:
            if e.code != "NoSuchKey":
                raise

        self.client.put_object(bucket_name, path, io.BytesIO(content), len(content))
        return File(name=path.split('/')[-1], path=path, content=content)

    def delete_folder(self, user_id: int, path: str):
        self._init_storage_for_user(user_id)
        self._validate_path(path)
        
        bucket_name = self._to_padded_string(user_id)
        objects = list(self.client.list_objects(bucket_name, prefix=path, recursive=True))
        if not objects:
            raise exceptions.ObjectNotFoundException()
        for obj in objects:
            self.client.remove_object(bucket_name, obj.object_name)

    def delete_file(self, user_id: int, path: str):
        self._init_storage_for_user(user_id)
        self._validate_path(path)

        bucket_name = self._to_padded_string(user_id)
        try:
            self.client.stat_object(bucket_name, path)
            self.client.remove_object(bucket_name, path)
        except S3Error as e:
            if e.code == "NoSuchKey":
                raise exceptions.ObjectNotFoundException()
            raise

    def get_file(self, user_id: int, path: str) -> File:
        self._init_storage_for_user(user_id)
        self._validate_path(path)
        
        bucket_name = self._to_padded_string(user_id)
        try:
            response = self.client.get_object(bucket_name, path)
            content = response.read()
            response.close()
            response.release_conn()
            return File(name=path.split('/')[-1], path=path, content=content)
        except S3Error as e:
            if e.code == "NoSuchKey":
                raise exceptions.ObjectNotFoundException()
            raise

    def get_files(self, user_id: int, file_name: str, limit: int, offset: int) -> list[File]:
        self._init_storage_for_user(user_id)
        
        bucket_name = self._to_padded_string(user_id)
        objects = list(self.client.list_objects(bucket_name, recursive=True))
        res = []
        for obj in objects:
            if file_name.lower() in obj.object_name.lower() and not obj.object_name.endswith('/'):
                res.append(obj)
        return [
            File(name=obj.object_name.split('/')[-1], path=obj.object_name)
            for obj in res[offset:offset + limit]
        ]

    def get_folder(self, user_id: int, path: str) -> Folder:
        self._init_storage_for_user(user_id)
        self._validate_path(path)
        
        path = path.strip('/')

        bucket_name = self._to_padded_string(user_id)
        if path != '' and not path.endswith('/'):
            path += '/'
        prefix = path if path else ""
        objects = list(self.client.list_objects(bucket_name, prefix=prefix, recursive=False))
        if not objects:
            raise exceptions.ObjectNotFoundException()
        folder = Folder(name=path.split('/')[-1], path=path)
        for obj in objects:
            relative_path = obj.object_name[len(path):] if path else obj.object_name
            if '/' in relative_path:
                folder.folders[obj.object_name] = Folder(name=relative_path.split('/')[0], path=obj.object_name)
            else:
                folder.files[obj.object_name] = File(name=obj.object_name.split('/')[-1], path=obj.object_name)
        return folder

    def _to_padded_string(self, n: int) -> str:
        return f"{n:03}"

    def _init_storage_for_user(self, user_id: int):
        bucket_name = self._to_padded_string(user_id)
        if not self.client.bucket_exists(bucket_name):
            self.client.make_bucket(bucket_name)

    def _validate_path(self, path: str):
        if path == '':
            return
        if path.strip().strip('/') == '' or path.strip('/').strip() == '':
            raise exceptions.ObjectInvalidName()
        if '' in [p.strip() for p in path.strip('/').split('/')]:
            raise exceptions.ObjectInvalidName()
