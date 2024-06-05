package app

import (
	"context"
	"errors"
	"images/internal/model"
	"images/internal/repo"
	"images/pkg/logger"
	"net/http"
)

type appImpl struct {
	r    repo.ImageRepo
	logs logger.Logger
}

func (a *appImpl) AddImage(ctx context.Context, img model.Image) (id uint64, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"Method": "AddImage",
		}, err)
	}()

	if len(img) > model.ImageMaxSize {
		err = model.ErrImageTooBig
		return 0, err
	}

	if _, ok := model.PermittedImageTypes[http.DetectContentType(img)]; !ok {
		err = model.ErrWrongFormat
		return 0, err
	}

	return a.r.AddImage(ctx, img)
}

func (a *appImpl) GetImage(ctx context.Context, id uint64) (img model.Image, err error) {
	defer func() {
		a.writeLog(logger.Fields{
			"Method": "GetImage",
		}, err)
	}()

	return a.r.GetImage(ctx, id)
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrDatabaseError) || errors.Is(err, model.ErrServiceError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
