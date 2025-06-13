import nats
import json
import asyncio
import jwt
from datetime import datetime, timezone, timedelta
import random
import string

from memory_accessor import MemoryAccessor
from common.exceptions import ExpiredRefreshTokenException
import common.dto as dto
import common.codes as codes

import logging
logging.basicConfig(
    level=logging.INFO
)

class NatsServer:
    def __init__(self, nats_url: str, secret: str):
        self.nc = None
        self.nats_url = nats_url
        self.memory_accessor = MemoryAccessor()

        self.secret = secret

    async def connect(self):
        self.nc = await nats.connect(self.nats_url)

    def exception_handler(method):
        async def wrapper(self, msg):
            try:
                await method(self, msg)
            except ExpiredRefreshTokenException as e:
                logging.error(f'{e}')
                error = dto.ErrorResponse(
                    code=codes.EXPIRED_REFRESH_TOKEN
                )
                await msg.respond(error.model_dump_json().encode())
            except RuntimeError as e:
                logging.error(f'{e}')
                error = dto.ErrorResponse(
                    code=codes.INTERNAL_ERROR
                )
                await msg.respond(error.model_dump_json().encode())
        return wrapper

    @exception_handler
    async def handle_create(self, msg):
        user_id = dto.CreateTokensMessage.model_validate_json(msg.data).user_id

        logging.info(f'creating tokens for user {user_id}')
        access = self.__create_access_token(user_id)
        refresh = self.__generate_refresh()
        self.memory_accessor.save_refresh_token(user_id, refresh)
        
        logging.info(f'created tokens for user {user_id}')
        await msg.respond(dto.CreateTokensResponse(access=access, refresh=refresh).model_dump_json().encode())


    @exception_handler
    async def handle_refresh_tokens(self, msg):
        refresh_token = dto.RefreshTokensMessage.model_validate_json(msg.data).refresh
        user_id = self.memory_accessor.get_user_id_by_token(refresh_token)

        logging.info(f'refreshing tokens for user {user_id}')
        access = self.__create_access_token(user_id)
        refresh = self.__generate_refresh()
        self.memory_accessor.save_refresh_token(user_id, refresh)
        
        logging.info(f'refreshed tokens for user {user_id}')
        await msg.respond(dto.RefreshTokensResponse(access=access, refresh=refresh).model_dump_json().encode())

    async def run(self):
        await self.connect()
        await self.nc.subscribe('create_tokens', cb=self.handle_create)
        await self.nc.subscribe('refresh_tokens', cb=self.handle_refresh_tokens)

        logging.info('started nats server')

        try:
            await asyncio.Future()
        except asyncio.CancelledError:
            pass
        finally:
            await self.nc.close()

    def __create_access_token(self, user_id: int, expires_after_hours=1):
        encode = {
            'user_id': user_id,
            'exp': datetime.now(timezone.utc) + timedelta(hours=expires_after_hours)
        }
        return jwt.encode(encode, self.secret, algorithm='HS256')

    def __generate_refresh(self, length=40) -> str:
        characters = string.ascii_letters + string.digits
        return ''.join(random.choices(characters, k=length))
