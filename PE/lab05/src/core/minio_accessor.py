from minio import Minio
import io

class MinioAccessor:
    def __init__(self, endpoint: str, access_key: str, secret_key: str):
        self.client = Minio(endpoint, access_key=access_key, secret_key=secret_key, secure=False)

    def create_file(self, user_id: int, name: str, content: bytes = b''):
        self._init_storage_for_user(user_id)
        
        bucket_name = self._to_padded_string(user_id)
        self.client.put_object(bucket_name, name, io.BytesIO(content), len(content))

    def delete_file(self, user_id: int, name: str):
        self._init_storage_for_user(user_id)

        bucket_name = self._to_padded_string(user_id)
        self.client.remove_object(bucket_name, name)
        

    def get_file(self, user_id: int, name: str) -> bytes:
        self._init_storage_for_user(user_id)
        
        bucket_name = self._to_padded_string(user_id)
        response = self.client.get_object(bucket_name, name)
        content = response.read()
        response.close()
        response.release_conn()
        return content

    def _to_padded_string(self, n: int) -> str:
        return f"{n:03}"

    def _init_storage_for_user(self, user_id: int):
        bucket_name = self._to_padded_string(user_id)
        if not self.client.bucket_exists(bucket_name):
            self.client.make_bucket(bucket_name)
