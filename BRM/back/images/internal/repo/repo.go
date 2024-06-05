package repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"images/internal/model"
)

type repoImpl struct {
	*pgxpool.Pool
}

const (
	addImageQuery = `
		INSERT INTO "images" (data)
		VALUES ($1)
		RETURNING "id";`

	getImageQuery = `
		SELECT "data" FROM "images"
		WHERE "id" = $1;`
)

func (r *repoImpl) AddImage(ctx context.Context, img model.Image) (uint64, error) {
	var id uint64
	if err := r.QueryRow(ctx, addImageQuery, img).Scan(&id); err != nil {
		return 0, errors.Join(model.ErrDatabaseError, err)
	}
	return id, nil
}

func (r *repoImpl) GetImage(ctx context.Context, id uint64) (model.Image, error) {
	var img model.Image
	if err := r.QueryRow(ctx, getImageQuery, id).Scan(&img); errors.Is(err, pgx.ErrNoRows) {
		return nil, model.ErrImageNotExists
	} else if err != nil {
		return nil, errors.Join(model.ErrDatabaseError, err)
	}
	return img, nil
}
