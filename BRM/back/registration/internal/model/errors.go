package model

import "errors"

var (
	ErrInvalidInput      = errors.New("введены некорректные данные")
	ErrIndustryNotExists = errors.New("запрашиваемая отрасль не существует")

	ErrCoreError   = errors.New("неполадки с сервисом core")
	ErrCoreUnknown = errors.New("неизвестная ошибка от сервиса core")
)
