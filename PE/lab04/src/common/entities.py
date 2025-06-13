from datetime import datetime
from typing import Dict

class User:
    def __init__(self, 
                 login: str, 
                 first_name: str, 
                 last_name: str, 
                 email: str,
                 is_admin: bool = False,
                 id: int = 0,
                 created_at: datetime = None,
                 password: str=''):
        self.id = id
        self.login = login
        self.first_name = first_name
        self.last_name = last_name
        self.email = email
        self.password = password
        self.is_admin = is_admin
        self.created_at = created_at

class LoginItems:
    def __init__(self, login: str, password: str):
        self.login = login
        self.password = password

class Tokens:
    def __init__(self, access: str, refresh: str):
        self.access = access
        self.refresh = refresh

class File:
    def __init__(self, name: str, path: str, content: bytes = ''):
        self.name = name
        self.path = path
        self.content = content

class Folder:
    def __init__(self, name: str, path: str):
        self.name = name
        self.path = path
        self.files: Dict[str, File] = {}
        self.folders: Dict[str, 'Folder'] = {}

class SubPagination:
    def __init__(self, folder_limit: int, folder_offset: int, file_limit: int, file_offset: int):
        self.folder_limit = folder_limit
        self.folder_offset = folder_offset
        self.file_limit = file_limit
        self.file_offset = file_offset

class RefreshToken:
    def __init__(self, refresh: str):
        self.refresh = refresh
