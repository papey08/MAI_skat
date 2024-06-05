package grpcserver

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stats/internal/model"
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
	case errors.Is(err, model.ErrCoreDatabase):
		c = codes.ResourceExhausted
		resErr = model.ErrCoreDatabase
	case errors.Is(err, model.ErrLeadsDatabase):
		c = codes.ResourceExhausted
		resErr = model.ErrLeadsDatabase
	case errors.Is(err, model.ErrAdsDatabase):
		c = codes.ResourceExhausted
		resErr = model.ErrAdsDatabase
	default:
		c = codes.Unknown
		resErr = model.ErrServiceErr
	}

	return status.Errorf(c, resErr.Error())
}
