package grpcserver

import (
	"brm-leads/internal/model"
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
	case errors.Is(err, model.ErrLeadNotExists):
		c = codes.NotFound
		resErr = model.ErrLeadNotExists
	case errors.Is(err, model.ErrCompanyNotExists):
		c = codes.NotFound
		resErr = model.ErrCompanyNotExists
	case errors.Is(err, model.ErrEmployeeNotExists):
		c = codes.NotFound
		resErr = model.ErrEmployeeNotExists
	case errors.Is(err, model.ErrStatusNotExists):
		c = codes.NotFound
		resErr = model.ErrStatusNotExists
	case errors.Is(err, model.ErrAdNotExists):
		c = codes.NotFound
		resErr = model.ErrAdNotExists
	case errors.Is(err, model.ErrAuthorization):
		c = codes.PermissionDenied
		resErr = model.ErrAuthorization
	case errors.Is(err, model.ErrValidationError):
		c = codes.FailedPrecondition
		resErr = model.ErrValidationError
	case errors.Is(err, model.ErrDatabaseError):
		c = codes.ResourceExhausted
		resErr = model.ErrDatabaseError
	case errors.Is(err, model.ErrServiceError):
		c = codes.ResourceExhausted
		resErr = model.ErrServiceError
	case errors.Is(err, model.ErrCoreError):
		c = codes.ResourceExhausted
		resErr = model.ErrCoreError
	case errors.Is(err, model.ErrAdsError):
		c = codes.ResourceExhausted
		resErr = model.ErrAdsError
	default:
		c = codes.Unknown
		resErr = model.ErrLeadsServiceError
	}

	return status.Errorf(c, resErr.Error())
}
