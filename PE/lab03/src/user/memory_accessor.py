from datetime import datetime, timezone

from common.entities import User
import common.exceptions as exceptions

# deprecated
class MemoryAccessor:
    def __init__(self):
        self.users_db = {
            1: User(
                id=1,
                login='admin',
                first_name='Илья',
                last_name='Ирбитский',
                email='myhamster@mail.ru',
                password='4d4384524b0f5b3a4750c60df2e70b7d4592966b757c698ed4a592ba928d9517', # hashed 'secret'
                is_admin=True,
                created_at=datetime.now(timezone.utc)
            )
        }
        self.logins = set()
        self.emails = set()
        self.current_id = 2

    def create_user(self, 
                    login: str,
                    first_name: str,
                    last_name: str, 
                    email: str, 
                    password: str) -> User:
        if login in self.logins or email in self.emails:
            raise exceptions.UserAlreadyExistsException()
        id = self.current_id
        user = User(
            id=id, 
            login=login,
            first_name=first_name,
            last_name=last_name,
            email=email, 
            password=password,
            is_admin=False,
            created_at=datetime.now(timezone.utc)
        )
        self.users_db[user.id] = user
        self.current_id += 1
        self.logins.add(login)
        self.emails.add(email)
        return user
    
    def get_user_by_id(self, user_id: int) -> User:
        user = self.users_db.get(user_id)
        if user:
            return user
        raise exceptions.UserNotFoundException()
    
    def get_users(self, login: str, first_name: str, last_name: str, limit: int, offset: int) -> list[User]:
        res = []
        for user in self.users_db.values():
            if ((login != '' and user.login == login) or login == '') and \
                ((first_name != '' and first_name.lower() in user.first_name.lower()) or first_name == '') and \
                ((last_name != '' and last_name.lower() in user.last_name.lower()) or last_name == ''):
                res.append(user)
        return res[offset:offset+limit]
    
    def login(self, login: str, password: str) -> User:
        for user in self.users_db.values():
            if user.login == login and user.password == password:
                return user
        raise exceptions.LoginUnableException()
