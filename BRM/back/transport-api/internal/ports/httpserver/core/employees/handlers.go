package employees

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

// @Summary		Добавление нового сотрудника
// @Description	Добавляет нового сотрудника в компанию
// @Tags			core/employees
// @Security		ApiKeyAuth
// @Accept			json
// @Produce		json
// @Param			input	body		addEmployeeRequest	true	"id сотрудника, которого добавляют в контакты"
// @Success		200		{object}	employeeResponse	"Успешное добавление сотрудника"
// @Failure		500		{object}	employeeResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	employeeResponse	"Неверный формат входных данных"
// @Failure		401		{object}	employeeResponse	"Ошибка авторизации"
// @Failure		403		{object}	employeeResponse	"Нет прав для выполнения операции"
// @Router			/employees [post]
func AddEmployee(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		var req addEmployeeRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		employee, err := a.CreateEmployee(c, companyId, ownerId, core.Employee{
			Id:           0,
			CompanyId:    req.CompanyId,
			FirstName:    req.FirstName,
			SecondName:   req.SecondName,
			Email:        req.Email,
			Password:     req.Password,
			JobTitle:     req.JobTitle,
			Department:   req.Department,
			ImageURL:     req.ImageURL,
			CreationDate: 0,
			IsDeleted:    false,
		})

		switch {
		case err == nil:
			data := employeeToEmployeeData(employee)
			c.JSON(http.StatusOK, employeeResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrEmailRegistered):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(model.ErrEmailRegistered))
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		case errors.Is(err, model.ErrAuthError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrAuthError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}

// @Summary		Получение списка сотрудников
// @Description	Получает список сотрудников компании с использованием фильтрации и пагинации
// @Tags			core/employees
// @Security		ApiKeyAuth
// @Produce		json
// @Param			limit		query		int						true	"Limit"
// @Param			offset		query		int						true	"Offset"
// @Param			name		query		string					false	"Поиск по имени/фамилии"
// @Param			jobtitle	query		string					false	"Поиск по должности"
// @Param			department	query		string					false	"Поиск по названию отдела"
// @Success		200			{object}	employeeListResponse	"Успешное получение сотрудников"
// @Failure		500			{object}	employeeListResponse	"Проблемы на стороне сервера"
// @Failure		400			{object}	employeeListResponse	"Неверный формат входных данных"
// @Failure		401			{object}	employeeResponse		"Ошибка авторизации"
// @Failure		403			{object}	employeeResponse		"Нет прав для выполнения операции"
// @Router			/employees [get]
func GetEmployeesList(a app.App) gin.HandlerFunc {
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

		var employees []core.Employee
		var amount uint

		if pattern, byName := c.GetQuery("name"); byName {
			employees, amount, err = a.GetEmployeeByName(c,
				companyId,
				employeeId,
				core.EmployeeByName{
					Pattern: pattern,
					Limit:   uint(limit),
					Offset:  uint(offset),
				},
			)
		} else {
			var filter core.FilterEmployee
			filter.Limit = uint(limit)
			filter.Offset = uint(offset)
			filter.JobTitle, filter.ByJobTitle = c.GetQuery("jobtitle")
			filter.Department, filter.ByDepartment = c.GetQuery("department")
			employees, amount, err = a.GetCompanyEmployees(c,
				companyId,
				employeeId,
				filter)
		}

		switch {
		case err == nil:
			data := &employeeListData{
				Employees: employeesToEmployeeDataList(employees),
				Amount:    amount,
			}
			c.JSON(http.StatusOK, employeeListResponse{
				Data: data,
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

// @Summary		Получение сотрудника
// @Description	Получает сотрудника по id
// @Tags			core/employees
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int					true	"id сотрудника"
// @Success		200	{object}	employeeResponse	"Успешное получение сотрудника"
// @Failure		500	{object}	employeeResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	employeeResponse	"Неверный формат входных данных"
// @Failure		404	{object}	employeeResponse	"Сотрудник не найден"
// @Failure		401	{object}	employeeResponse	"Ошибка авторизации"
// @Failure		403	{object}	employeeResponse	"Нет прав для выполнения операции"
// @Router			/employees/{id} [get]
func GetEmployee(a app.App) gin.HandlerFunc {
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

		employee, err := a.GetEmployeeById(c,
			companyId,
			employeeId,
			id)

		switch {
		case err == nil:
			data := employeeToEmployeeData(employee)
			c.JSON(http.StatusOK, employeeResponse{
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

// @Summary		Редактирование сотрудника
// @Description	Изменяет одно или несколько полей сотрудника
// @Tags			core/employees
// @Security		ApiKeyAuth
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"id сотрудника"
// @Param			input	body		updateEmployeeRequest	true	"Новые поля"
// @Success		200		{object}	employeeResponse		"Успешное обновление данных"
// @Failure		500		{object}	employeeResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	employeeResponse		"Неверный формат входных данных"
// @Failure		404		{object}	employeeResponse		"Сотрудник не найден"
// @Failure		401		{object}	employeeResponse		"Ошибка авторизации"
// @Failure		403		{object}	employeeResponse		"Нет прав для выполнения операции"
// @Router			/employees/{id} [put]
func UpdateEmployee(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		var req updateEmployeeRequest
		if err = c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		employee, err := a.UpdateEmployee(c,
			companyId,
			ownerId,
			id,
			core.UpdateEmployee{
				FirstName:  req.FirstName,
				SecondName: req.SecondName,
				JobTitle:   req.JobTitle,
				Department: req.Department,
				ImageURL:   req.ImageURL,
			},
		)
		switch {
		case err == nil:
			data := employeeToEmployeeData(employee)
			c.JSON(http.StatusOK, employeeResponse{
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

// @Summary		Удаление сотрудника
// @Description	Безвозвратно удаляет сотрудника и все его поля
// @Tags			core/employees
// @Security		ApiKeyAuth
// @Produce		json
// @Param			id	path		int					true	"id сотрудника"
// @Success		200	{object}	employeeResponse	"Успешное удаление контакта"
// @Failure		500	{object}	employeeResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	employeeResponse	"Неверный формат входных данных"
// @Failure		404	{object}	employeeResponse	"Сотрудник не найден"
// @Failure		401	{object}	employeeResponse	"Ошибка авторизации"
// @Failure		403	{object}	employeeResponse	"Нет прав для выполнения операции"
// @Router			/employees/{id} [delete]
func DeleteEmployee(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		err = a.DeleteEmployee(c,
			companyId,
			ownerId,
			id)
		switch {
		case err == nil:
			c.JSON(http.StatusOK, employeeResponse{
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
		case errors.Is(err, model.ErrOwnerDeletion):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(model.ErrOwnerDeletion))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}

	}
}
