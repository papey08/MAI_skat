package model

import "errors"

var (
	ErrInvalidInput = errors.New("введены некорректные данные")

	ErrWrongPassword      = errors.New("введён неправильный пароль")
	ErrCreateAccessToken  = errors.New("невозможно создать токен доступа")
	ErrParsingAccessToken = errors.New("невозможно прочитать токен доступа")

	ErrAccessTokenNotExpired = errors.New("время действия токена доступа ещё не истекло")
	ErrTokensNotExist        = errors.New("токены не найдены")

	ErrEmployeeNotExists = errors.New("запрашиваемый сотрудник не существует")
	ErrEmailRegistered   = errors.New("данный email уже зарегистрирован")

	ErrAuthRepoError = errors.New("неполадки с базой данных токенов")
	ErrPassRepoError = errors.New("неполадки с базой данных паролей")

	ErrServiceError = errors.New("неполадки с сервисом")
)
