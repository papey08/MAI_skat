import redis
from datetime import timedelta
from common.exceptions import ExpiredRefreshTokenException

class RedisAccessor:
    def __init__(self, redis_url: str):
        self.redis = redis.Redis.from_url(redis_url, decode_responses=True)

    def save_refresh_token(self, user_id: int, refresh_token: str):
        expiration = timedelta(days=30)
        self.redis.setex(refresh_token, expiration, user_id)

    def get_user_id_by_token(self, refresh_token: str) -> int:
        user_id = self.redis.get(refresh_token)
        if user_id is not None:
            self.redis.delete(refresh_token)
            return int(user_id)
        else:
            raise ExpiredRefreshTokenException()
