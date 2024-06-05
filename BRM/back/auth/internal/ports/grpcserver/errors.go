package grpcserver

import (
	"auth/internal/model"
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
	case errors.Is(err, model.ErrEmployeeNotExists):
		c = codes.NotFound
		resErr = model.ErrEmployeeNotExists
	case errors.Is(err, model.ErrEmailRegistered):
		c = codes.AlreadyExists
		resErr = model.ErrEmailRegistered
	case errors.Is(err, model.ErrPassRepoError):
		c = codes.ResourceExhausted
		resErr = model.ErrPassRepoError
	case errors.Is(err, model.ErrInvalidInput):
		c = codes.FailedPrecondition
		resErr = model.ErrInvalidInput
	default:
		c = codes.Unknown
		resErr = model.ErrServiceError
	}
	return status.Errorf(c, resErr.Error())
}
