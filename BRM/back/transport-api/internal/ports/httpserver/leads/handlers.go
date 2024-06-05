package leads

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/model/leads"
	"transport-api/internal/ports/httpserver/middleware"
)

// @Summary		Получение сделки
// @Description	Получает сделку по id
// @Tags			leads
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int				true	"id сделки"
// @Success		200	{object}	leadResponse	"Успешное получение сделки"
// @Failure		500	{object}	leadResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	leadResponse	"Неверный формат входных данных"
// @Failure		404	{object}	leadResponse	"Сделка не найдена"
// @Failure		401	{object}	leadResponse	"Ошибка авторизации"
// @Router			/leads/{id} [get]
func GetLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		lead, err := a.GetLeadById(c, companyId, employeeId, id)

		switch {
		case err == nil:
			data := leadToLeadData(lead)
			c.JSON(http.StatusOK, leadResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrLeadNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrLeadNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrLeadsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		}
	}
}

// @Summary		Получение списка сделок
// @Description	Получает список сделок с использованием фильтрации и пагинации
// @Tags			leads
// @Security		ApiKeyAuth
// @Produce		json
// @Param			limit		query		int					true	"Limit"
// @Param			offset		query		int					true	"Offset"
// @Param			responsible	query		int					false	"Фильтрация по id ответственного"
// @Param			status		query		string				false	"Фильтрация по статусу"
// @Success		200			{object}	leadsListResponse	"Успешное получение сделок"
// @Failure		500			{object}	leadsListResponse	"Проблемы на стороне сервера"
// @Failure		400			{object}	leadsListResponse	"Неверный формат входных данных"
// @Failure		401			{object}	leadResponse		"Ошибка авторизации"
// @Router			/leads [get]
func GetLeadsList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, companyId, ok := middleware.GetAuthData(c)
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

		var filter leads.Filter
		filter.Limit = uint(limit)
		filter.Offset = uint(offset)

		if responsibleStr, ok := c.GetQuery("responsible"); ok {
			filter.ByResponsible = true
			filter.Responsible, err = strconv.ParseUint(responsibleStr, 10, 64)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
				return
			}
		}

		filter.Status, filter.ByStatus = c.GetQuery("status")

		leadsList, amount, err := a.GetLeads(c, companyId, employeeId, filter)
		switch {
		case err == nil:
			data := &leadsListData{
				Leads:  leadsToLeadsDataList(leadsList),
				Amount: amount,
			}
			c.JSON(http.StatusOK, leadsListResponse{
				Data: data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrLeadsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		}
	}
}

// @Summary		Редактирование сделки
// @Description	Изменяет одно или несколько полей сделки
// @Tags			leads
// @Security		ApiKeyAuth
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"id сделки"
// @Param			input	body		updateLeadRequest	true	"Новые поля"
// @Success		200		{object}	leadResponse		"Успешное обновление данных"
// @Failure		500		{object}	leadResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	leadResponse		"Неверный формат входных данных"
// @Failure		404		{object}	leadResponse		"Сделка не найдена"
// @Failure		403		{object}	leadResponse		"Нет прав на редактрование сделки"
// @Failure		401		{object}	leadResponse		"Ошибка авторизации"
// @Router			/leads/{id} [put]
func UpdateLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		var req updateLeadRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		lead, err := a.UpdateLead(c, companyId, employeeId, id, leads.UpdateLead{
			Title:       req.Title,
			Description: req.Description,
			Price:       req.Price,
			Status:      req.Status,
			Responsible: req.Responsible,
		})

		switch {
		case err == nil:
			data := leadToLeadData(lead)
			c.JSON(http.StatusOK, leadResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrStatusNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrStatusNotExists))
		case errors.Is(err, model.ErrLeadNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrLeadNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrLeadsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		}
	}
}

// @Summary		Получение статусов и их id
// @Description	Возвращает мапу со статусами и их id
// @Tags			leads
// @Security		ApiKeyAuth
// @Produce		json
// @Success		200	{object}	statusesResponse	"Успешное получение"
// @Failure		500	{object}	statusesResponse	"Проблемы на стороне сервера"
// @Failure		401	{object}	statusesResponse	"Ошибка авторизации"
// @Router			/leads/statuses [get]
func GetStatuses(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		statuses, err := a.GetStatuses(c)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, statusesResponse{
				Data: statuses,
				Err:  nil,
			})
		case errors.Is(err, model.ErrLeadsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrLeadsError))
		}
	}
}
