package grpcserver

import (
	"brm-ads/internal/model"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapErrors(err error) error {
	var c codes.Code
	var resErr error

	switch {
	case err == nil:
		return nil
	case errors.Is(err, model.ErrAdNotExists):
		c = codes.NotFound
		resErr = model.ErrAdNotExists
	case errors.Is(err, model.ErrCompanyNotExists):
		c = codes.NotFound
		resErr = model.ErrCompanyNotExists
	case errors.Is(err, model.ErrIndustryNotExists):
		c = codes.NotFound
		resErr = model.ErrIndustryNotExists
	case errors.Is(err, model.ErrAuthorization):
		c = codes.PermissionDenied
		resErr = model.ErrAdNotExists
	case errors.Is(err, model.ErrSameCompany):
		c = codes.FailedPrecondition
		resErr = model.ErrSameCompany
	case errors.Is(err, model.ErrValidationError):
		c = codes.FailedPrecondition
		resErr = model.ErrValidationError
	case errors.Is(err, model.ErrDatabaseError):
		c = codes.ResourceExhausted
		resErr = model.ErrDatabaseError
	case errors.Is(err, model.ErrCoreError):
		c = codes.ResourceExhausted
		resErr = model.ErrCoreError
	case errors.Is(err, model.ErrLeadsServiceError):
		c = codes.ResourceExhausted
		resErr = model.ErrLeadsServiceError
	default:
		c = codes.Unknown
		resErr = model.ErrAdsServiceError
	}

	return status.Errorf(c, resErr.Error())
}
