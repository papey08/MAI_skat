from postgres_accessor import PostgresAccessor
from cache_accessor import CacheAccessor

from common.entities import User


class UserAccessor:
    def __init__(self, database_url: str, redis_url: str):
        self.postgres_accessor = PostgresAccessor(database_url)
        self.cache_accessor = CacheAccessor(redis_url)
    
    def create_user(self, 
                    login: str,
                    first_name: str,
                    last_name: str, 
                    email: str, 
                    password: str) -> User:
        user = self.postgres_accessor.create_user(login, first_name, last_name, email, password)
        self.cache_accessor.cache_user(user)
        return user
    
    def get_user_by_id(self, user_id: int) -> User:
        user = self.cache_accessor.get_user(user_id)
        if not user:
            user = self.postgres_accessor.get_user_by_id(user_id)
            self.cache_accessor.cache_user(user)
        return user
    
    def get_users(self, login: str, first_name: str, last_name: str, limit: int, offset: int) -> list[User]:
        return self.postgres_accessor.get_users(login, first_name, last_name, limit, offset)
    
    def login(self, login: str, password: str) -> User:
        return self.postgres_accessor.login(login, password)
