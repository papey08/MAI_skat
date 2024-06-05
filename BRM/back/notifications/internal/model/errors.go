package model

import "errors"

var (
	ErrPermissionDenied          = errors.New("no rights to make this operation")
	ErrNotificationNotFound      = errors.New("notification with required id does not exist")
	ErrWrongNotificationType     = errors.New("found notification with wrong type")
	ErrClosedLeadAlreadyAnswered = errors.New("notification of closed lead with required id already answered")
	ErrCompanyNotFound           = errors.New("company with required id does not exist")

	ErrDatabaseError = errors.New("something wrong with notifications database")
	ErrStatsError    = errors.New("something wrong with stats service")
	ErrServiceError  = errors.New("something wrong with notifications service")
)
