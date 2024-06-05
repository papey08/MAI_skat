package grpcserver

import (
	"brm-core/internal/model"
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
	case errors.Is(err, model.ErrCompanyNotExists):
		c = codes.NotFound
		resErr = model.ErrCompanyNotExists
	case errors.Is(err, model.ErrEmployeeNotExists):
		c = codes.NotFound
		resErr = model.ErrEmployeeNotExists
	case errors.Is(err, model.ErrContactNotExists):
		c = codes.NotFound
		resErr = model.ErrContactNotExists
	case errors.Is(err, model.ErrIndustryNotExists):
		c = codes.NotFound
		resErr = model.ErrIndustryNotExists
	case errors.Is(err, model.ErrAuthorization):
		c = codes.PermissionDenied
		resErr = model.ErrAuthorization
	case errors.Is(err, model.ErrEmailRegistered):
		c = codes.AlreadyExists
		resErr = model.ErrEmailRegistered
	case errors.Is(err, model.ErrContactExist):
		c = codes.AlreadyExists
		resErr = model.ErrContactExist
	case errors.Is(err, model.ErrDatabaseError):
		c = codes.ResourceExhausted
		resErr = model.ErrDatabaseError
	case errors.Is(err, model.ErrAuthServiceError):
		c = codes.ResourceExhausted
		resErr = model.ErrAuthServiceError
	case errors.Is(err, model.ErrValidationError):
		c = codes.FailedPrecondition
		resErr = model.ErrValidationError
	case errors.Is(err, model.ErrOwnerDeletion):
		c = codes.Aborted
		resErr = model.ErrOwnerDeletion
	case errors.Is(err, model.ErrSelfContact):
		c = codes.FailedPrecondition
		resErr = model.ErrSelfContact
	default:
		c = codes.Unknown
		resErr = model.ErrServiceError
	}

	return status.Errorf(c, resErr.Error())
}
