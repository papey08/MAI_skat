package notifications

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/ports/httpserver/middleware"
)

// @Summary		Получение списка уведомлений
// @Description	Получает список уведомлений
// @Tags			notifications
// @Security		ApiKeyAuth
// @Produce		json
// @Param			limit			query		int							true	"Limit"
// @Param			offset			query		int							true	"Offset"
// @Param			only_not_viewed	query		bool						false	"Вернуть только непрочитанные уведомления"
// @Success		200				{object}	notificationListResponse	"Успешное получение объявлений"
// @Failure		500				{object}	notificationListResponse	"Проблемы на стороне сервера"
// @Failure		400				{object}	notificationListResponse	"Неверный формат входных данных"
// @Failure		401				{object}	notificationListResponse	"Ошибка авторизации"
// @Router			/notifications [get]
func GetNotifications(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, companyId, ok := middleware.GetAuthData(c)
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
		onlyNotViewedStr, _ := c.GetQuery("only_not_viewed")
		onlyNotViewed := onlyNotViewedStr == "true"

		notifications, amount, err := a.GetNotifications(c, companyId, uint(limit), uint(offset), onlyNotViewed)
		switch {
		case err == nil:
			data := &notificationListData{
				Notifications: notificationsToNotificationsDataList(notifications),
				Amount:        amount,
			}
			c.JSON(http.StatusOK, notificationListResponse{
				Data: data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrNotificationsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsUnknown))
		}
	}
}

// @Summary		Получение уведомления
// @Description	Получает уведомление по id
// @Tags			notifications
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int						true	"id уведомления"
// @Success		200	{object}	notificationResponse	"Успешное получение уведомления"
// @Failure		500	{object}	notificationResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	notificationResponse	"Неверный формат входных данных"
// @Failure		401	{object}	notificationResponse	"Ошибка авторизации"
// @Failure		404	{object}	notificationResponse	"Уведомление не найдено"
// @Failure		403	{object}	notificationResponse	"Нет прав для выполнения действия"
// @Router			/notifications/{id} [get]
func GetNotification(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		notification, err := a.GetNotification(c, companyId, id)

		switch {
		case err == nil:
			data := notificationToNotificationData(notification)
			c.JSON(http.StatusOK, notificationResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrNotificationNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrNotificationNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrNotificationsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsUnknown))
		}
	}
}

// @Summary		Подтверждение закрытой сделки
// @Description	Когда компания-поставщик услуги отмечает сделку закрытой, в компанию клиента приходит уведомление об этом. Клиент может подтвердить закрытие сделки, чтобы у поставщика вырос рейтинг
// @Tags			notifications
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id		path		int						true	"id уведомления"
// @Param			submit	query		bool					true	"Подтверждено/не подтверждено"
// @Success		200		{object}	notificationResponse	"Успешное подтверждение"
// @Failure		500		{object}	notificationResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	notificationResponse	"Неверный формат входных данных"
// @Failure		401		{object}	notificationResponse	"Ошибка авторизации"
// @Failure		404		{object}	notificationResponse	"Уведомление не найдено"
// @Failure		403		{object}	notificationResponse	"Нет прав для выполнения действия"
// @Failure		409		{object}	notificationResponse	"Попытка повторного подтверждения"
// @Router			/notifications/{id} [post]
func SubmitClosedLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		submitStr, _ := c.GetQuery("submit")
		submit := submitStr == "true"

		err = a.SubmitClosedLead(c, companyId, id, submit)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, errorResponse(nil))
		case errors.Is(err, model.ErrNotificationNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrNotificationNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrNotificationAnswered):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(model.ErrNotificationAnswered))
		case errors.Is(err, model.ErrNotificationsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsUnknown))
		}
	}
}
