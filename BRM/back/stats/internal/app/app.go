package app

import (
	"context"
	"errors"
	"stats/internal/model"
	"stats/internal/repo"
	"stats/pkg/logger"
)

type appImpl struct {
	repo repo.Repo

	logs logger.Logger
}

func (a *appImpl) GetCompanyMainPageStats(ctx context.Context, id uint64) (stat model.MainPageStats, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": id,
			"Method":     "GetCompanyMainPageStats",
		}, err)
	}()

	return a.repo.GetCompanyMainPageStats(ctx, id)
}

func (a *appImpl) UpdateRatingByClosedLead(ctx context.Context, companyId uint64, submit bool) (err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"Method":     "UpdateRatingByClosedLead",
		}, err)
	}()

	currRating, err := a.repo.GetCompanyRating(ctx, companyId)
	if err != nil {
		return err
	}

	var newRating float64
	if submit {
		newRating = increaseRating(currRating)
	} else {
		newRating = decreaseRating(currRating)
	}
	return a.repo.SetCompanyRating(ctx, companyId, newRating)
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrCoreDatabase) || errors.Is(err, model.ErrLeadsDatabase) || errors.Is(err, model.ErrAdsDatabase) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}

func increaseRating(currRating float64) float64 {
	switch {
	case currRating == 5:
		return 5
	case currRating >= 4:
		currRating += 0.05
	case currRating >= 3:
		currRating += 0.1
	case currRating >= 2:
		currRating += 0.25
	default:
		currRating += 0.5
	}
	if currRating >= 5 {
		currRating = 5
	}
	return currRating
}

func decreaseRating(currRating float64) float64 {
	currRating -= 0.5
	if currRating < 0 {
		return 0
	}
	return currRating
}
