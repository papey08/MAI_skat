from sqlalchemy import create_engine, select
from sqlalchemy.exc import IntegrityError
from sqlalchemy.orm import sessionmaker
from sqlalchemy.sql import func

from common.entities import User
from user.postgres_dto import User as UserDto
import common.exceptions as exceptions

class PostgresAccessor:
    def __init__(self, database_url: str):
        self.engine = create_engine(database_url)
        self.Session = sessionmaker(bind=self.engine)

    def create_user(self, 
                    login: str,
                    first_name: str,
                    last_name: str, 
                    email: str, 
                    password: str) -> User:
        session = self.Session()
        try:
            user = UserDto(
                login=login,
                first_name=first_name,
                last_name=last_name,
                email=email,
                password=password
            )
            session.add(user)
            session.commit()
            session.refresh(user)
            return User(
                id=user.id,
                login=user.login,
                first_name=user.first_name,
                last_name=user.last_name,
                email=user.email,
                is_admin=user.is_admin,
                created_at=user.created_at
            )
        except IntegrityError:
            session.rollback()
            raise exceptions.UserAlreadyExistsException()
        except Exception as e:
            session.rollback()
            raise exceptions.InternalException()
        finally:
            session.close()

    def get_user_by_id(self, user_id: int) -> User:
        session = self.Session()
        try:
            user = session.get(UserDto, user_id)
            if not user:
                raise exceptions.UserNotFoundException()
            return User(
                id=user.id,
                login=user.login,
                first_name=user.first_name,
                last_name=user.last_name,
                email=user.email,
                is_admin=user.is_admin,
                created_at=user.created_at
            )
        finally:
            session.close()

    def get_users(self, login: str, first_name: str, last_name: str, limit: int, offset: int) -> list[User]:
        session = self.Session()
        try:
            query = select(UserDto).where(
                (UserDto.login == login) if login else True,
                (func.lower(UserDto.first_name).like(f'%{first_name.lower()}%')) if first_name else True,
                (func.lower(UserDto.last_name).like(f'%{last_name.lower()}%')) if last_name else True
            ).limit(limit).offset(offset)
            return [
                User(
                    id=user.id,
                    login=user.login,
                    first_name=user.first_name, 
                    last_name=user.last_name,
                    email=user.email,
                    is_admin=user.is_admin,
                    created_at=user.created_at
                ) for user in session.scalars(query).all()
            ]
        finally:
            session.close()

    def login(self, login: str, password: str) -> User:
        session = self.Session()
        try:
            query = select(UserDto).where(
                UserDto.login == login,
                UserDto.password == password
            )
            user = session.scalars(query).first()
            if not user:
                raise exceptions.LoginUnableException()
            return User(
                id=user.id,
                login=user.login,
                first_name=user.first_name,
                last_name=user.last_name,
                email=user.email,
                is_admin=user.is_admin,
                created_at=user.created_at
            )
        finally:
            session.close()
