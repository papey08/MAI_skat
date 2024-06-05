package model

import "errors"

var (
	ErrInvalidInput = errors.New("введены некорректные данные")

	ErrCompanyNotExists  = errors.New("запрашиваемая компания не существует")
	ErrEmployeeNotExists = errors.New("запрашиваемый сотрудник не существует")
	ErrContactNotExists  = errors.New("запрашиваемый контакт не существует")
	ErrIndustryNotExists = errors.New("запрашиваемая отрасль не существует")
	ErrLeadNotExists     = errors.New("запрашиваемая сделка не существует")
	ErrStatusNotExists   = errors.New("запрашиваемый этап не существует")
	ErrEmailRegistered   = errors.New("данный email уже зарегистрирован")
	ErrContactExist      = errors.New("создаваемый контакт уже существует")
	ErrSelfContact       = errors.New("нельзя добавить самого себя в контакты")
	ErrOwnerDeletion     = errors.New("владелец компании не может быть удалён")

	ErrAdNotExists = errors.New("запрашиваемое объявление не существует")
	ErrSameCompany = errors.New("невозможно откликнуться на объявление собственной компании")

	ErrPermissionDenied = errors.New("недостаточно прав для выполнения этой операции")
	ErrUnauthorized     = errors.New("необходима авторизация")

	ErrNotificationNotExists = errors.New("запрашиваемое уведомление не существует")
	ErrNotificationAnswered  = errors.New("на запрашиваемое уведомление уже произведён ответ")

	ErrAuthError            = errors.New("неполадки с сервисом авторизации")
	ErrCoreError            = errors.New("неполадки с сервисом core")
	ErrCoreUnknown          = errors.New("неизвестная ошибка от сервиса core")
	ErrAdsError             = errors.New("неполадки с сервисом объявлений")
	ErrAdsUnknown           = errors.New("неизвестная ошибка от сервиса объявлений")
	ErrLeadsError           = errors.New("неполадки с сервисом сделок")
	ErrStatsError           = errors.New("неполадки с сервисом статистики")
	ErrNotificationsError   = errors.New("неполадки с сервисом уведомлений")
	ErrNotificationsUnknown = errors.New("неизвестная ошибка от сервиса объявлений")
)
