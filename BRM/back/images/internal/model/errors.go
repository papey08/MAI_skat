package model

import "errors"

var (
	ErrInvalidInput   = errors.New("введены некорректные данные")
	ErrWrongFormat    = errors.New("некорректный формат изображения")
	ErrImageTooBig    = errors.New("слишком большое изображение")
	ErrImageNotExists = errors.New("запрашиваемое изображение не существует")

	ErrDatabaseError = errors.New("неполадки с базой данных изображений")
	ErrServiceError  = errors.New("неполадки с сервисом")
)
