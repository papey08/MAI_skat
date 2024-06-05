package httpserver

import (
	"auth/internal/app"
	"auth/internal/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Обновление токенов
// @Description	Обновляет access и refresh-токены, старая пара становится непригодной
// @Accept			json
// @Produce		json
// @Param			input	body		refreshRequest	true	"Пара токенов"
// @Success		200		{object}	tokensResponse	"Успешное обновление токенов"
// @Failure		500		{object}	tokensResponse	"Проблемы на стороне сервера"
// @Failure		404		{object}	tokensResponse	"Пара токенов не найдена (refresh-токен истёк)"
// @Failure		400		{object}	tokensResponse	"Неверный формат входных данных"
// @Router			/refresh [post]
func refresh(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req refreshRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		tokens, err := a.RefreshTokens(c, model.TokensPair{
			Access:  req.Access,
			Refresh: req.Refresh,
		})

		switch {
		case err == nil:
			data := tokensData{
				Access:  tokens.Access,
				Refresh: tokens.Refresh,
			}
			c.JSON(http.StatusOK, tokensResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCreateAccessToken):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCreateAccessToken))
		case errors.Is(err, model.ErrParsingAccessToken):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrParsingAccessToken))
		case errors.Is(err, model.ErrAccessTokenNotExpired):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrAccessTokenNotExpired))
		case errors.Is(err, model.ErrTokensNotExist):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrTokensNotExist))
		case errors.Is(err, model.ErrAuthRepoError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAuthRepoError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}

// @Summary		Получение токенов
// @Description	Получает access и refresh-токены, используя аутентификацию по логину и паролю
// @Accept			json
// @Produce		json
// @Param			input	body		loginRequest	true	"Логин и пароль"
// @Success		200		{object}	tokensResponse	"Успешное получение токенов"
// @Failure		500		{object}	tokensResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	tokensResponse	"Неверный формат входных данных"
// @Failure		403		{object}	tokensResponse	"Неверный пароль"
// @Failure		404		{object}	tokensResponse	"Пользователь с запрашиваемым email не найден"
// @Router			/login [post]
func login(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		tokens, err := a.LoginEmployee(c, req.Email, req.Password)
		switch {
		case err == nil:
			data := tokensData{
				Access:  tokens.Access,
				Refresh: tokens.Refresh,
			}
			c.JSON(http.StatusOK, tokensResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCreateAccessToken):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCreateAccessToken))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrWrongPassword):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrWrongPassword))
		case errors.Is(err, model.ErrAuthRepoError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAuthRepoError))
		case errors.Is(err, model.ErrPassRepoError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrPassRepoError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}

	}
}

// @Summary		Выход из аккаунта
// @Description	Удаляет пару токенов
// @Accept			json
// @Produce		json
// @Param			input	body		logoutRequest	true	"Пара токенов"
// @Success		200		{object}	tokensResponse	"Успешный выход из аккаунта"
// @Failure		500		{object}	tokensResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	tokensResponse	"Неверный формат входных данных"
// @Router			/logout [post]
func logout(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req logoutRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		err := a.LogoutEmployee(c, model.TokensPair{
			Access:  req.Access,
			Refresh: req.Refresh,
		})

		switch {
		case err == nil:
			c.JSON(http.StatusOK, tokensResponse{
				Data: nil,
				Err:  nil,
			})
		case errors.Is(err, model.ErrAuthRepoError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAuthRepoError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}
