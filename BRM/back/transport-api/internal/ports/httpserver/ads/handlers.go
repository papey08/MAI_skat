package ads

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/model/ads"
	"transport-api/internal/ports/httpserver/middleware"
)

// @Summary		Добавление нового объявления
// @Description	Добавляет новое объявление
// @Tags			ads
// @Security		ApiKeyAuth
// @Accept			json
// @Produce		json
// @Param			input	body		addAdRequest	true	"Новое объявление в JSON"
// @Success		200		{object}	adResponse		"Успешное добавление объявления"
// @Failure		500		{object}	adResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	adResponse		"Неверный формат входных данных"
// @Failure		401		{object}	adResponse		"Ошибка авторизации"
// @Router			/market [post]
func AddAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		var req addAdRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		ad, err := a.CreateAd(c, companyId, employeeId, ads.Ad{
			CompanyId:   companyId,
			Title:       req.Title,
			Text:        req.Text,
			Industry:    req.Industry,
			Price:       req.Price,
			ImageURL:    req.ImageURL,
			CreatedBy:   employeeId,
			Responsible: employeeId,
		})

		switch {
		case err == nil:
			data := adToAdData(ad)
			c.JSON(http.StatusOK, adResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrAdsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsError))
		case errors.Is(err, model.ErrIndustryNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrIndustryNotExists))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsUnknown))
		}
	}
}

// @Summary		Получение объявления
// @Description	Получает объявление по id
// @Tags			ads
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int			true	"id объявления"
// @Success		200	{object}	adResponse	"Успешное получение объявления"
// @Failure		500	{object}	adResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	adResponse	"Неверный формат входных данных"
// @Failure		401	{object}	adResponse	"Ошибка авторизации"
// @Failure		404	{object}	adResponse	"Объявление не найдено"
// @Router			/market/{id} [get]
func GetAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		ad, err := a.GetAdById(c, id)
		switch {
		case err == nil:
			data := adToAdData(ad)
			c.JSON(http.StatusOK, adResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrAdNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrAdNotExists))
		case errors.Is(err, model.ErrAdsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsUnknown))
		}

	}
}

// @Summary		Получение списка объявлений
// @Description	Получает список объявлений с использованием фильтрации и пагинации
// @Tags			ads
// @Security		ApiKeyAuth
// @Produce		json
// @Param			limit			query		int				true	"Limit"
// @Param			offset			query		int				true	"Offset"
// @Param			pattern			query		string			false	"Поиск по названию/тексту"
// @Param			company_id		query		int				false	"Поиск по компании"
// @Param			industry		query		string			false	"Поиск по отрасли"
// @Param			by_price		query		bool			false	"Сортировка по возрастанию цены"
// @Param			by_price_desc	query		bool			false	"Сортировка по убыванию цены"
// @Param			by_date			query		bool			false	"Сортировка по возрастанию даты создания"
// @Param			by_date_desc	query		bool			false	"Сортировка по убыванию даты создания"
// @Success		200				{object}	adListResponse	"Успешное получение объявлений"
// @Failure		500				{object}	adListResponse	"Проблемы на стороне сервера"
// @Failure		400				{object}	adListResponse	"Неверный формат входных данных"
// @Failure		401				{object}	adResponse		"Ошибка авторизации"
// @Router			/market [get]
func GetAdsList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, ok := middleware.GetAuthData(c)
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

		var params ads.ListParams
		params.Limit = uint(limit)
		params.Offset = uint(offset)

		var adList []ads.Ad
		var amount uint

		if pattern, byName := c.GetQuery("pattern"); byName {
			params.Search = &ads.AdSearcher{Pattern: pattern}
			adList, amount, err = a.GetAdsList(c, params)
		} else {
			filter := &ads.AdFilter{}
			_, filter.ByIndustry = c.GetQuery("industry")
			_, filter.ByCompany = c.GetQuery("company_id")
			if !filter.ByIndustry && !filter.ByCompany {
				filter = nil
			} else {
				if filter.ByIndustry {
					filter.Industry = c.Query("industry")
				}
				if filter.ByCompany {
					filter.CompanyId, err = strconv.ParseUint(c.Query("company_id"), 10, 64)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
						return
					}
				}
			}

			sorter := &ads.AdSorter{}

			byDate, _ := c.GetQuery("by_date")
			byPrice, _ := c.GetQuery("by_price")
			byPriceDesc, _ := c.GetQuery("by_price_desc")
			byDateDesc, _ := c.GetQuery("by_date_desc")

			switch {
			case byDate == "true":
				sorter.ByDateAsc = true
			case byPrice == "true":
				sorter.ByPriceAsc = true
			case byPriceDesc == "true":
				sorter.ByPriceDesc = true
			case byDateDesc == "true":
				sorter.ByDateDesc = true
			}

			if !sorter.ByDateAsc && !sorter.ByDateDesc && !sorter.ByPriceAsc && !sorter.ByPriceDesc {
				sorter = nil
			}

			params.Filter = filter
			params.Sort = sorter
			adList, amount, err = a.GetAdsList(c, params)
		}

		switch {
		case err == nil:
			data := &adListData{
				Ads:    adsToAdDataList(adList),
				Amount: amount,
			}
			c.JSON(http.StatusOK, adListResponse{
				Data: data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrAdsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsUnknown))
		}
	}
}

// @Summary		Редактирование объявления
// @Description	Изменяет одно или несколько полей объявления
// @Tags			ads
// @Security		ApiKeyAuth
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"id объявления"
// @Param			input	body		updateAdRequest	true	"Новые поля"
// @Success		200		{object}	adResponse		"Успешное обновление данных"
// @Failure		500		{object}	adResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	adResponse		"Неверный формат входных данных"
// @Failure		404		{object}	adResponse		"Объявление не найдено"
// @Router			/market/{id} [put]
func UpdateAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		var req updateAdRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		ad, err := a.UpdateAd(c, companyId, employeeId, id, ads.UpdateAd{
			Title:       req.Title,
			Text:        req.Text,
			Industry:    req.Industry,
			Price:       req.Price,
			Responsible: req.Responsible,
		})

		switch {
		case err == nil:
			data := adToAdData(ad)
			c.JSON(http.StatusOK, adResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrAdNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrAdNotExists))
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrIndustryNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrIndustryNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrAdsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsUnknown))
		}
	}
}

// @Summary		Удаление объявления
// @Description	Безвозвратно удаляет объявление
// @Tags			ads
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int			true	"id объявления"
// @Success		200	{object}	adResponse	"Успешное удаление объявления"
// @Failure		500	{object}	adResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	adResponse	"Неверный формат входных данных"
// @Failure		404	{object}	adResponse	"Объявление не найдено"
// @Router			/market/{id} [delete]
func DeleteAd(a app.App) gin.HandlerFunc {
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

		err = a.DeleteAd(c, companyId, employeeId, id)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, adResponse{
				Data: nil,
				Err:  nil,
			})
		case errors.Is(err, model.ErrAdNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrAdNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrAdsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsUnknown))
		}
	}
}

// @Summary		Откликнуться на объявление
// @Description	Создаёт отклик у откликнувшейся компании и сделку у владельца объявления
// @Tags			ads
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int					true	"id объявления"
// @Success		200	{object}	responseResponse	"Успешное создание отклика"
// @Failure		500	{object}	responseResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	responseResponse	"Неверный формат входных данных"
// @Failure		404	{object}	responseResponse	"Объявление не найдено"
// @Failure		409	{object}	responseResponse	"Попытка откликнуться на объявление своей же компании"
// @Router			/market/{id}/response [post]
func AddResponse(a app.App) gin.HandlerFunc {
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

		resp, err := a.CreateResponse(c, companyId, employeeId, id)

		switch {
		case err == nil:
			data := responseToResponseData(resp)
			c.JSON(http.StatusOK, responseResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrAdNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrAdNotExists))
		case errors.Is(err, model.ErrSameCompany):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(model.ErrSameCompany))
		case errors.Is(err, model.ErrAdsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsUnknown))
		}
	}
}

// @Summary		Получение списка откликов на объявления
// @Description	Возвращает список откликов компании на объявления
// @Tags			ads
// @Security		ApiKeyAuth
// @Produce		json
// @Param			limit	query		int						true	"Limit"
// @Param			offset	query		int						true	"Offset"
// @Success		200		{object}	responseListResponse	"Успешное получение списка"
// @Failure		500		{object}	responseListResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	responseListResponse	"Неверный формат входных данных"
// @Router			/responses [get]
func GetResponsesList(a app.App) gin.HandlerFunc {
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

		responses, amount, err := a.GetResponses(c, companyId, employeeId, uint(limit), uint(offset))

		switch {
		case err == nil:
			data := &responseListData{
				Responses: responsesToResponseDataList(responses),
				Amount:    amount,
			}
			c.JSON(http.StatusOK, responseListResponse{
				Data: data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrAdsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAdsUnknown))
		}
	}
}

// @Summary		Получение отраслей
// @Description	Возвращает словарь из возможных отраслей объявлений и их id
// @Tags			ads
// @Security		ApiKeyAuth
// @Produce		json
// @Success		200	{object}	industriesResponse	"Успешное получение данных"
// @Failure		500	{object}	industriesResponse	"Проблемы на стороне сервера"
// @Router			/market/industries [get]
func GetIndustriesMap(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		industries, err := a.GetAdsIndustries(c)

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
