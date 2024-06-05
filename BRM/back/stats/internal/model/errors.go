package model

import "errors"

var (
	ErrCompanyNotExists = errors.New("company with required id does not exist")
	ErrCoreDatabase     = errors.New("something wrong with core database")
	ErrAdsDatabase      = errors.New("something wrong with ads database")
	ErrLeadsDatabase    = errors.New("something wrong with leads database")
	ErrServiceErr       = errors.New("something wrong with stats service")
)
