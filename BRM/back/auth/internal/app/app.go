package app

import (
	"auth/internal/app/tokenizer"
	"auth/internal/app/valid"
	"auth/internal/model"
	"auth/internal/repo/authrepo"
	"auth/internal/repo/passrepo"
	"auth/pkg/logger"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
)

type appImpl struct {
	passwordSalt       string
	refreshTokenLength int

	authRepo  authrepo.AuthRepo
	passRepo  passrepo.PassRepo
	tokenizer tokenizer.Tokenizer

	logs logger.Logger
}

func (a *appImpl) RegisterEmployee(ctx context.Context, employee model.Employee) (err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"employee_email": employee.Email,
			"employee_id":    employee.EmployeeId,
			"company_id":     employee.CompanyId,
			"Method":         "RegisterEmployee",
		}, err)
	}()

	if !valid.Password(employee.Password) {
		return model.ErrInvalidInput
	}

	employee.Password = a.getHashedPassword(employee.Password)
	return a.passRepo.CreateEmployee(ctx, employee)
}

func (a *appImpl) DeleteEmployee(ctx context.Context, email string) (err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"employee_email": email,
			"Method":         "DeleteEmployee",
		}, err)
	}()

	return a.passRepo.DeleteEmployee(ctx, email)
}

func (a *appImpl) LoginEmployee(ctx context.Context, email string, password string) (pair model.TokensPair, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"employee_email": email,
			"Method":         "LoginEmployee",
		}, err)
	}()

	employee, err := a.passRepo.GetEmployee(ctx, email)
	if err != nil {
		return model.TokensPair{}, err
	}

	if a.getHashedPassword(password) != employee.Password {
		return model.TokensPair{}, model.ErrWrongPassword
	}

	var tokens model.TokensPair
	tokens.Access, err = a.tokenizer.CreateToken(employee.EmployeeId, employee.CompanyId)
	if err != nil {
		return model.TokensPair{}, model.ErrCreateAccessToken
	}

	tokens.Refresh = a.createRandomString()

	if err = a.authRepo.SetTokens(ctx, tokens); err != nil {
		return model.TokensPair{}, model.ErrAuthRepoError
	}

	return tokens, nil
}

func (a *appImpl) RefreshTokens(ctx context.Context, tokens model.TokensPair) (pair model.TokensPair, err error) {
	var employeeId, companyId uint64
	defer func() {
		a.writeLog(logger.Fields{
			"employee_id": employeeId,
			"company_id":  companyId,
			"Method":      "RefreshTokens",
		}, err)
	}()

	existingTokens, err := a.authRepo.GetTokens(ctx, tokens.Access)
	if err != nil {
		return model.TokensPair{}, err
	}

	var isValid bool
	if isValid, err = a.tokenizer.CheckExpiration(existingTokens.Access); err != nil {
		return model.TokensPair{}, model.ErrParsingAccessToken
	} else if isValid {
		return model.TokensPair{}, model.ErrAccessTokenNotExpired
	}

	err = a.authRepo.DeleteTokens(ctx, existingTokens.Access)
	if err != nil {
		return model.TokensPair{}, err
	}

	employeeId, companyId, err = a.tokenizer.DecryptToken(existingTokens.Access)
	if err != nil {
		return model.TokensPair{}, err
	}

	var newTokens model.TokensPair
	newTokens.Access, err = a.tokenizer.CreateToken(employeeId, companyId)
	if err != nil {
		return model.TokensPair{}, model.ErrCreateAccessToken
	}

	newTokens.Refresh = a.createRandomString()

	if err = a.authRepo.SetTokens(ctx, tokens); err != nil {
		return model.TokensPair{}, model.ErrAuthRepoError
	}
	return newTokens, nil
}

func (a *appImpl) LogoutEmployee(ctx context.Context, tokens model.TokensPair) (err error) {
	employeeId, companyId, err := a.tokenizer.DecryptToken(tokens.Access)

	err = a.authRepo.DeleteTokens(ctx, tokens.Access)

	a.writeLog(logger.Fields{
		"employee_id": employeeId,
		"company_id":  companyId,
		"Method":      "LogoutEmployee",
	}, err)
	return err
}

func (a *appImpl) getHashedPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + a.passwordSalt))
	hashSum := hash.Sum(nil)
	return hex.EncodeToString(hashSum)
}

func (a *appImpl) createRandomString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	buffer := make([]byte, a.refreshTokenLength)
	for i := range buffer {
		buffer[i] = charset[rand.Intn(len(charset))]
	}
	return string(buffer)
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrCreateAccessToken) ||
		errors.Is(err, model.ErrEmailRegistered) || // по идее этот кейс обрабатывается в ядре, поэтому тут его быть не должно
		errors.Is(err, model.ErrAuthRepoError) ||
		errors.Is(err, model.ErrPassRepoError) ||
		errors.Is(err, model.ErrServiceError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
