package contacts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
	"transport-api/internal/ports/httpserver/middleware"
)

// @Summary		Добавление нового контакта
// @Description	Добавляет новый контакт в список контактов сотрудника
// @Tags			core/contacts
// @Security		ApiKeyAuth
// @Accept			json
// @Produce		json
// @Param			input	body		addContactRequest	true	"id сотрудника, которого добавляют в контакты"
// @Success		200		{object}	contactResponse		"Успешное добавление контакта"
// @Failure		500		{object}	contactResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	contactResponse		"Неверный формат входных данных"
// @Failure		404		{object}	contactResponse		"Добавляемый сотрудник не найден"
// @Failure		401		{object}	contactResponse		"Ошибка авторизации"
// @Failure		403		{object}	contactResponse		"Нет прав для выполнения операции"
// @Router			/contacts [post]
func AddContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		var req addContactRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		contact, err := a.CreateContact(c, ownerId, req.EmployeeId)

		switch {
		case err == nil:
			data := contactToContactData(contact)
			c.JSON(http.StatusOK, contactResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrContactExist):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(model.ErrContactExist))
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrSelfContact):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(model.ErrSelfContact))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}

// @Summary		Получение списка контактов
// @Description	Получает список контактов сотрудника с использованием фильтрации и пагинации
// @Tags			core/contacts
// @Security		ApiKeyAuth
// @Produce		json
// @Param			limit	query		int				true	"Limit"
// @Param			offset	query		int				true	"Offset"
// @Success		200		{object}	contactResponse	"Успешное получение контактов"
// @Failure		500		{object}	contactResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	contactResponse	"Неверный формат входных данных"
// @Failure		401		{object}	contactResponse	"Ошибка авторизации"
// @Failure		403		{object}	contactResponse	"Нет прав для выполнения операции"
// @Router			/contacts [get]
func GetContactsList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		contacts, amount, err := a.GetContacts(c,
			ownerId,
			core.GetContacts{
				Limit:  uint(limit),
				Offset: uint(offset),
			})
		switch {
		case err == nil:
			data := contactsToContactDataList(contacts)
			c.JSON(http.StatusOK, сontactListResponse{
				Data: contactListData{
					Contacts: data,
					Amount:   amount,
				},
				Err: nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}

// @Summary		Получение контакта
// @Description	Получает контакт по id
// @Tags			core/contacts
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int				true	"id контакта"
// @Success		200	{object}	contactResponse	"Успешное получение контакта"
// @Failure		500	{object}	contactResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	contactResponse	"Неверный формат входных данных"
// @Failure		404	{object}	contactResponse	"Контакт не найден"
// @Failure		401	{object}	contactResponse	"Ошибка авторизации"
// @Failure		403	{object}	contactResponse	"Нет прав для выполнения операции"
// @Router			/contacts/{id} [get]
func GetContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		contact, err := a.GetContactById(c, ownerId, id)
		switch {
		case err == nil:
			data := contactToContactData(contact)
			c.JSON(http.StatusOK, contactResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}

// @Summary		Редактирование контакта
// @Description	Изменяет одно или несколько полей контакта
// @Tags			core/contacts
// @Security		ApiKeyAuth
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"id контакта"
// @Param			input	body		updateContactRequest	true	"Новые поля"
// @Success		200		{object}	contactResponse			"Успешное обновление данных"
// @Failure		500		{object}	contactResponse			"Проблемы на стороне сервера"
// @Failure		400		{object}	contactResponse			"Неверный формат входных данных"
// @Failure		404		{object}	contactResponse			"Контакт не найден"
// @Failure		401		{object}	contactResponse			"Ошибка авторизации"
// @Failure		403		{object}	contactResponse			"Нет прав для выполнения операции"
// @Router			/contacts/{id} [put]
func UpdateContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		var req updateContactRequest
		if err = c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		contact, err := a.UpdateContact(c,
			ownerId,
			id,
			core.UpdateContact{
				Notes: req.Notes,
			})
		switch {
		case err == nil:
			data := contactToContactData(contact)
			c.JSON(http.StatusOK, contactResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}

// @Summary		Удаление контакта
// @Description	Безвозвратно удаляет контакт и все его поля
// @Tags			core/contacts
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int				true	"id контакта"
// @Success		200	{object}	contactResponse	"Успешное удаление контакта"
// @Failure		500	{object}	contactResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	contactResponse	"Неверный формат входных данных"
// @Failure		404	{object}	contactResponse	"Контакт не найден"
// @Failure		401	{object}	contactResponse	"Ошибка авторизации"
// @Failure		403	{object}	contactResponse	"Нет прав для выполнения операции"
// @Router			/contacts/{id} [delete]
func DeleteContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		err = a.DeleteContact(c, ownerId, id)
		switch {
		case err == nil:
			c.JSON(http.StatusOK, contactResponse{
				Data: nil,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}
