from pydantic import BaseModel
from datetime import datetime

class User(BaseModel):
    id: int
    first_name: str
    last_name: str
    login: str
    email: str
    password: str
    is_admin: bool
    created_at: datetime
