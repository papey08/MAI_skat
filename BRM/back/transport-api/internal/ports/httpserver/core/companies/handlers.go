package companies

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
	"transport-api/internal/ports/httpserver/middleware"
	//"transport-api/internal/ports/httpserver"
)

// @Summary		Получение информации о компании
// @Description	Возвращает название и статистику компании для главной страницы
// @Tags			core/companies
// @Security		ApiKeyAuth
// @Produce		json
// @Success		200	{object}	mainPageResponse	"Успешное получение данных"
// @Failure		500	{object}	mainPageResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	mainPageResponse	"Неверный формат входных данных"
// @Failure		404	{object}	mainPageResponse	"Компания не найдена"
// @Failure		401	{object}	mainPageResponse	"Ошибка авторизации"
// @Failure		403	{object}	mainPageResponse	"Нет прав для выполнения операции"
// @Router			/companies/mainpage [get]
func GetCompanyMainPage(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		page, err := a.GetCompanyMainPage(c, companyId)
		switch {
		case err == nil:
			data := mainPageToData(page)
			c.JSON(http.StatusOK, mainPageResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrStatsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrStatsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrStatsError))
		}
	}
}

// @Summary		Получение информации о компании
// @Description	Возвращает поля компании для страницы редактирования
// @Tags			core/companies
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int				true	"id компании"
// @Success		200	{object}	companyResponse	"Успешное получение данных"
// @Failure		500	{object}	companyResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	companyResponse	"Неверный формат входных данных"
// @Failure		404	{object}	companyResponse	"Компания не найдена"
// @Failure		401	{object}	companyResponse	"Ошибка авторизации"
// @Failure		403	{object}	companyResponse	"Нет прав для выполнения операции"
// @Router			/companies/{id} [get]
func GetCompany(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		_, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		company, err := a.GetCompany(c, id)
		switch {
		case err == nil:
			data := companyToCompanyData(company)
			c.JSON(http.StatusOK, companyResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}

// @Summary		Редактирование полей компании
// @Description	Изменяет одно или несколько полей компании
// @Tags			core/companies
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id		path		int						true	"id компании"
// @Param			input	body		updateCompanyRequest	true	"Новые поля"
// @Success		200		{object}	companyResponse			"Успешное обновление данных"
// @Failure		500		{object}	companyResponse			"Проблемы на стороне сервера"
// @Failure		400		{object}	companyResponse			"Неверный формат входных данных"
// @Failure		404		{object}	companyResponse			"Компания не найдена"
// @Failure		401		{object}	companyResponse			"Ошибка авторизации"
// @Failure		403		{object}	companyResponse			"Нет прав для выполнения операции"
// @Router			/companies/{id} [put]
func UpdateCompany(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		ownerId, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		var req updateCompanyRequest
		if err = c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		company, err := a.UpdateCompany(c, companyId, ownerId, core.UpdateCompany{
			Name:        req.Name,
			Description: req.Description,
			Industry:    req.Industry,
			OwnerId:     req.OwnerId,
		})

		switch {
		case err == nil:
			data := companyToCompanyData(company)
			c.JSON(http.StatusOK, companyResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrIndustryNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrIndustryNotExists))
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

// @Summary		Получение отраслей
// @Description	Возвращает словарь из возможных отраслей компаний и их id
// @Tags			core/companies
// @Security		ApiKeyAuth
// @Produce		json
// @Success		200	{object}	industriesResponse	"Успешное получение данных"
// @Failure		500	{object}	industriesResponse	"Проблемы на стороне сервера"
// @Router			/companies/industries [get]
func GetIndustriesMap(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		industries, err := a.GetCompanyIndustries(c)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, industriesResponse{
				Data: industries,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}
