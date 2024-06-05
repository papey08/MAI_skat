package grpcserver

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"notifications/internal/model"
)

func mapErrors(err error) error {
	var c codes.Code
	var resErr error

	switch {
	case err == nil:
		return nil
	case errors.Is(err, model.ErrPermissionDenied):
		c = codes.PermissionDenied
		resErr = model.ErrPermissionDenied
	case errors.Is(err, model.ErrNotificationNotFound):
		c = codes.NotFound
		resErr = model.ErrNotificationNotFound
	case errors.Is(err, model.ErrWrongNotificationType):
		c = codes.FailedPrecondition
		resErr = model.ErrWrongNotificationType
	case errors.Is(err, model.ErrClosedLeadAlreadyAnswered):
		c = codes.Canceled
		resErr = model.ErrClosedLeadAlreadyAnswered
	case errors.Is(err, model.ErrCompanyNotFound):
		c = codes.NotFound
		resErr = model.ErrCompanyNotFound
	case errors.Is(err, model.ErrDatabaseError):
		c = codes.ResourceExhausted
		resErr = model.ErrDatabaseError
	case errors.Is(err, model.ErrStatsError):
		c = codes.ResourceExhausted
		resErr = model.ErrStatsError
	case errors.Is(err, model.ErrServiceError):
		c = codes.ResourceExhausted
		resErr = model.ErrServiceError
	}
	return status.Errorf(c, resErr.Error())
}
