import redis
from datetime import timedelta

from postgres_dto import User as UserDb
from redis_dto import User as UserCache

class CacheAccessor:
    def __init__(self, redis_url: str):
        self.redis = redis.Redis.from_url(redis_url, decode_responses=True)

    def cache_user(self, user: UserDb):
        self.redis.setex(user.id, timedelta(hours=1), UserCache(
            id=user.id,
            first_name=user.first_name,
            last_name=user.last_name,
            login=user.login,
            email=user.email,
            password=user.password,
            is_admin=user.is_admin,
            created_at=user.created_at
            ).model_dump_json())

    def get_user(self, user_id: int) -> UserDb:
        user = self.redis.get(user_id)
        if not user:
            return None
        user_cache = UserCache.model_validate_json(user)
        return UserDb(
            id=user_cache.id,
            login=user_cache.login,
            first_name=user_cache.first_name,
            last_name=user_cache.last_name,
            email=user_cache.email,
            password=user_cache.password,
            is_admin=user_cache.is_admin,
            created_at=user_cache.created_at
        )
