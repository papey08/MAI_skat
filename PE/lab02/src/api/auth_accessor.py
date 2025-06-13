import nats
from nats.aio.client import Client as NATSClient
from typing import Optional
from pydantic import ValidationError


from common.entities import Tokens
import common.dto as dto
import common.exceptions as exceptions
import common.codes as codes

class AuthAccessor:
    def __init__(self, nats_url: str):
        self.nats_url = nats_url
        self.nc: Optional[NATSClient] = None

    async def __aenter__(self):
        self.nc = await nats.connect(self.nats_url)
        return self
    
    async def __aexit__(self, exc_type, exc_val, exc_tb):
        if self.nc:
            await self.nc.close()

    async def create_tokens(self, user_id: int) -> Tokens:
        try:
            msg = dto.CreateTokensMessage(user_id=user_id)
            response = await self.nc.request('create_tokens', msg.model_dump_json().encode(), timeout=30)
            res = dto.CreateTokensResponse.model_validate_json(response.data)
            return Tokens(access=res.access, refresh=res.refresh)
        except ValidationError:
            err = dto.ErrorResponse.model_validate_json(response.data)
            if err.code == codes.INTERNAL_ERROR:
                raise exceptions.InternalException()
    
    async def refresh_tokens(self, refresh_token: str) -> Tokens:
        try:
            msg = dto.RefreshTokensMessage(refresh=refresh_token)
            response = await self.nc.request('refresh_tokens', msg.model_dump_json().encode(), timeout=30)
            res = dto.RefreshTokensResponse.model_validate_json(response.data)
            return Tokens(access=res.access, refresh=res.refresh)
        except ValidationError:
            err = dto.ErrorResponse.model_validate_json(response.data)
            if err.code == codes.EXPIRED_REFRESH_TOKEN:
                raise exceptions.ExpiredRefreshTokenException()
            if err.code == codes.INTERNAL_ERROR:
                raise exceptions.InternalException()
