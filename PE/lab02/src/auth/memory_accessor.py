from datetime import datetime, timezone, timedelta

from common.exceptions import ExpiredRefreshTokenException

class MemoryAccessor:
    def __init__(self):
        self.tokens_db = {}

    def save_refresh_token(self, user_id: int, refresh_token: str):
        self.tokens_db[refresh_token] = (user_id, datetime.now(timezone.utc) + timedelta(days=30))

    def get_user_id_by_token(self, refresh_token: str) -> int:
        if refresh_token in self.tokens_db and self.tokens_db[refresh_token][1] > datetime.now(timezone.utc):
            user_id = self.tokens_db[refresh_token][0]
            del self.tokens_db[refresh_token]
            return user_id
        else:
            raise ExpiredRefreshTokenException()
