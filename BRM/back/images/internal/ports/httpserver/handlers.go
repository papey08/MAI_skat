package httpserver

import (
	"errors"
	"github.com/gin-gonic/gin"
	"images/internal/app"
	"images/internal/model"
	"io"
	"net/http"
	"strconv"
)

// @Summary		Добавление изображения
// @Description	Добавляет изображение
// @Accept			mpfd
// @Produce		json
// @Param			file	formData	file		true	"Изображение"
// @Success		200		{object}	idResponse	"Успешное добавление изображения"
// @Failure		500		{object}	idResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	idResponse	"Неверный формат входных данных"
// @Router			/images [post]
func handleAddImage(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		defer func() {
			_ = file.Close()
		}()

		data, err := io.ReadAll(file)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
			return
		}

		id, err := a.AddImage(c, data)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, idResponse{
				Id:  &id,
				Err: nil,
			})
		case errors.Is(err, model.ErrImageTooBig):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrImageTooBig))
		case errors.Is(err, model.ErrWrongFormat):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrWrongFormat))
		case errors.Is(err, model.ErrDatabaseError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrDatabaseError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}

// @Summary		Получение изображения
// @Description	Получает изображение
// @Produce		image/png
// @Param			id	path		int			true	"id изображения"
// @Success		200	{file}		file		"Успешное получение изображения"
// @Failure		500	{object}	idResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	idResponse	"Неверный формат входных данных"
// @Failure		404	{object}	idResponse	"Изображение не найдено"
// @Router			/images/{id} [get]
func handleGetImage(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		img, err := a.GetImage(c, id)

		switch {
		case err == nil:
			c.Data(http.StatusOK, http.DetectContentType(img), img)
		case errors.Is(err, model.ErrImageNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrImageNotExists))
		case errors.Is(err, model.ErrDatabaseError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrDatabaseError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}
