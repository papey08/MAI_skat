from common import codes

class UserAlreadyExistsException(Exception):
    def __init__(self):
        super().__init__(codes.USER_ALREADY_EXISTS)

class UserNotFoundException(Exception):
    def __init__(self):
        super().__init__(codes.USER_NOT_FOUND)

class UserInvalidException(Exception):
    def __init__(self):
        super().__init__(codes.USER_INVALID)

class LoginUnableException(Exception):
    def __init__(self):
        super().__init__(codes.LOGIN_UNABLE)

class ExpiredRefreshTokenException(Exception):
    def __init__(self):
        super().__init__(codes.EXPIRED_REFRESH_TOKEN)

class ObjectNotFoundException(Exception):
    def __init__(self):
        super().__init__(codes.OBJECT_NOT_FOUND)

class ObjectAlreadyExistsException(Exception):
    def __init__(self):
        super().__init__(codes.OBJECT_ALREADY_EXISTS)

class ObjectInvalidName(Exception):
    def __init__(self):
        super().__init__(codes.OBJECT_INVALID_NAME)

class UnauthorizedException(Exception):
    def __init__(self):
        super().__init__(codes.UNAUTHORIZED)

class AccessDeniedException(Exception):
    def __init__(self):
        super().__init__(codes.ACCESS_DENIED)

class InternalException(Exception):
    def __init__(self):
        super().__init__(codes.INTERNAL_ERROR)
