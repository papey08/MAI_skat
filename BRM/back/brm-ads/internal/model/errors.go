package model

import "errors"

var (
	ErrAdNotExists       = errors.New("ad with required id does not exist")
	ErrCompanyNotExists  = errors.New("company with required id does not exist")
	ErrEmployeeNotExists = errors.New("employee with required id does not exist")
	ErrIndustryNotExists = errors.New("industry with required id does not exist")

	ErrAuthorization   = errors.New("you don't have rights for this operation")
	ErrValidationError = errors.New("validation error")
	ErrSameCompany     = errors.New("you can't response to ad from your company")

	ErrDatabaseError     = errors.New("something wrong with ads database")
	ErrCoreError         = errors.New("something wrong with brm-core service")
	ErrAdsServiceError   = errors.New("something wrong with ads service")
	ErrLeadsServiceError = errors.New("something wrong with leads service")
)
