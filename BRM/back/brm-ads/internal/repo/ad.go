package repo

import (
	"brm-ads/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	createAdQuery = `
		INSERT INTO "ads" ("company_id", "title", "text", "industry", "price", "image_url", "creation_date", "created_by", "responsible", "is_deleted")
		VALUES ($1, $2, $3, (SELECT "industries"."id" FROM "industries" WHERE "name" = $4), $5, $6, $7, $8, $9, $10)
		RETURNING "id";`

	updateAdQuery = `
		UPDATE "ads"
		SET "title" = $2,
		    "text" = $3,
		    "industry" = (SELECT "industries"."id" FROM "industries" WHERE "name" = $4),
		    "price" = $5,
		    "image_url" = $6,
		    "responsible" = $7
		WHERE "id" = $1 AND (NOT "is_deleted");`

	deleteAdQuery = `
		UPDATE "ads"
		SET "is_deleted" = TRUE
		WHERE "id" = $1 AND (NOT "is_deleted");`

	getAdByIdQuery = `
		SELECT  "ads"."id",
		        "ads"."company_id",
		        "ads"."title",
		        "ads"."text",
		        "industries"."name",
		        "ads"."price",
		        "ads"."image_url",
		        "ads"."creation_date",
		        "ads"."created_by",
		        "ads"."responsible",
		        "ads"."is_deleted"
		       FROM "ads"
		INNER JOIN "industries" ON "industries"."id" = "ads"."industry"        
		WHERE "ads"."id" = $1 AND (NOT "is_deleted");`

	getAdsByPatternQuery = `
		SELECT  "ads"."id",
		        "ads"."company_id",
		        "ads"."title",
		        "ads"."text",
		        "industries"."name",
		        "ads"."price",
		        "ads"."image_url",
		        "ads"."creation_date",
		        "ads"."created_by",
		        "ads"."responsible",
		        "ads"."is_deleted"
		       FROM "ads"
		INNER JOIN "industries" ON "industries"."id" = "ads"."industry" 
		WHERE ("title" LIKE $1 OR "text" LIKE $1) AND (NOT "is_deleted")
		LIMIT $2 OFFSET $3;`

	getAdsByPatternAmountQuery = `
		SELECT  COUNT(*) FROM "ads"
		WHERE ("title" LIKE $1 OR "text" LIKE $1) AND (NOT "is_deleted");`

	getIndustriesQuery = `
		SELECT * FROM "industries";`
)

func (a *adRepoImpl) CreateAd(ctx context.Context, ad model.Ad) (model.Ad, error) {
	var adId uint64
	var pgErr *pgconn.PgError
	if err := a.QueryRow(ctx, createAdQuery,
		ad.CompanyId,
		ad.Title,
		ad.Text,
		ad.Industry,
		ad.Price,
		ad.ImageURL,
		ad.CreationDate,
		ad.CreatedBy,
		ad.Responsible,
		ad.IsDeleted,
	).Scan(&adId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23502":
			return model.Ad{}, model.ErrIndustryNotExists
		default:
			return model.Ad{}, errors.Join(model.ErrDatabaseError, err)
		}
	} else {
		ad.Id = adId
		return ad, nil
	}
}

func (a *adRepoImpl) UpdateAd(ctx context.Context, adId uint64, upd model.UpdateAd) (model.Ad, error) {
	var pgErr *pgconn.PgError
	if e, err := a.Exec(ctx, updateAdQuery,
		adId,
		upd.Title,
		upd.Text,
		upd.Industry,
		upd.Price,
		upd.ImageURL,
		upd.Responsible,
	); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23502":
			return model.Ad{}, model.ErrIndustryNotExists
		default:
			return model.Ad{}, errors.Join(model.ErrDatabaseError, err)
		}
	} else if e.RowsAffected() == 0 {
		return model.Ad{}, model.ErrAdNotExists
	} else {
		return a.GetAdById(ctx, adId)
	}
}

func (a *adRepoImpl) DeleteAd(ctx context.Context, adId uint64) error {
	if e, err := a.Exec(ctx, deleteAdQuery,
		adId,
	); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrAdNotExists
	} else {
		return nil
	}
}

func (a *adRepoImpl) GetAdById(ctx context.Context, id uint64) (model.Ad, error) {
	row := a.QueryRow(ctx, getAdByIdQuery, id)
	var ad model.Ad
	if err := row.Scan(
		&ad.Id,
		&ad.CompanyId,
		&ad.Title,
		&ad.Text,
		&ad.Industry,
		&ad.Price,
		&ad.ImageURL,
		&ad.CreationDate,
		&ad.CreatedBy,
		&ad.Responsible,
		&ad.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Ad{}, model.ErrAdNotExists
	} else if err != nil {
		return model.Ad{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		return ad, nil
	}
}

func (a *adRepoImpl) GetAdsList(ctx context.Context, params model.AdsListParams) ([]model.Ad, uint, error) {
	if params.Search != nil {
		var amount uint
		if err := a.QueryRow(ctx, getAdsByPatternAmountQuery,
			params.Search.Pattern+"%",
		).Scan(&amount); err != nil {
			return []model.Ad{}, 0, errors.Join(model.ErrDatabaseError, err)
		} else if amount == 0 {
			return []model.Ad{}, 0, nil
		}
		rows, err := a.Query(ctx, getAdsByPatternQuery,
			params.Search.Pattern+"%",
			params.Limit,
			params.Offset)
		if err != nil {
			return []model.Ad{}, 0, errors.Join(model.ErrDatabaseError, err)
		}
		defer rows.Close()

		ads := make([]model.Ad, 0)
		for rows.Next() {
			var ad model.Ad
			_ = rows.Scan(
				&ad.Id,
				&ad.CompanyId,
				&ad.Title,
				&ad.Text,
				&ad.Industry,
				&ad.Price,
				&ad.ImageURL,
				&ad.CreationDate,
				&ad.CreatedBy,
				&ad.Responsible,
				&ad.IsDeleted)
			ads = append(ads, ad)
		}
		return ads, amount, nil
	} else {
		if params.Filter == nil {
			params.Filter = &model.AdFilter{
				ByCompany:  false,
				CompanyId:  0,
				ByIndustry: false,
				Industry:   "",
			}
		}
		getAdsQuery := fmt.Sprintf(`
			SELECT  "ads"."id",
					"ads"."company_id",
					"ads"."title",
					"ads"."text",
					"industries"."name",
					"ads"."price",
					"ads"."image_url",
					"ads"."creation_date",
					"ads"."created_by",
					"ads"."responsible",
					"ads"."is_deleted"
				    FROM "ads"
			INNER JOIN "industries" ON "industries"."id" = "ads"."industry" 
			WHERE (NOT "is_deleted")
				AND ((NOT $1) OR "company_id" = $2)
				AND ((NOT $3) OR "industry" = (SELECT "industries"."id" FROM "industries" WHERE "name" = $4))
			%s
			LIMIT $5 OFFSET $6;`, getOrderParam(params.Sort))
		getAdsAmountQuery := `
			SELECT  COUNT(*) FROM "ads"
			WHERE (NOT "is_deleted")
				AND ((NOT $1) OR "company_id" = $2)
				AND ((NOT $3) OR "industry" = (SELECT "industries"."id" FROM "industries" WHERE "name" = $4));`

		var amount uint
		if err := a.QueryRow(ctx, getAdsAmountQuery,
			params.Filter.ByCompany,
			params.Filter.CompanyId,
			params.Filter.ByIndustry,
			params.Filter.Industry,
		).Scan(&amount); err != nil {
			return []model.Ad{}, 0, errors.Join(model.ErrDatabaseError, err)
		} else if amount == 0 {
			return []model.Ad{}, 0, nil
		}

		rows, err := a.Query(ctx, getAdsQuery,
			params.Filter.ByCompany,
			params.Filter.CompanyId,
			params.Filter.ByIndustry,
			params.Filter.Industry,
			params.Limit,
			params.Offset)
		if err != nil {
			return []model.Ad{}, 0, errors.Join(model.ErrDatabaseError, err)
		}
		defer rows.Close()

		ads := make([]model.Ad, 0)
		for rows.Next() {
			var ad model.Ad
			_ = rows.Scan(
				&ad.Id,
				&ad.CompanyId,
				&ad.Title,
				&ad.Text,
				&ad.Industry,
				&ad.Price,
				&ad.ImageURL,
				&ad.CreationDate,
				&ad.CreatedBy,
				&ad.Responsible,
				&ad.IsDeleted)
			ads = append(ads, ad)
		}
		return ads, amount, nil
	}
}

func getOrderParam(s *model.AdSorter) string {
	if s == nil {
		return "ORDER BY \"creation_date\" DESC"
	}
	switch {
	case s.ByPriceAsc:
		return "ORDER BY \"price\""
	case s.ByPriceDesc:
		return "ORDER BY \"price\" DESC"
	case s.ByDateAsc:
		return "ORDER BY \"creation_date\""
	case s.ByDateDesc:
		return "ORDER BY \"creation_date\" DESC"
	default:
		return "ORDER BY \"creation_date\" DESC"
	}
}

func (a *adRepoImpl) GetIndustries(ctx context.Context) (map[string]uint64, error) {
	rows, err := a.Query(ctx, getIndustriesQuery)
	if err != nil {
		return map[string]uint64{}, model.ErrDatabaseError
	}
	defer rows.Close()

	industries := make(map[string]uint64)
	for rows.Next() {
		var id uint64
		var industry string
		_ = rows.Scan(&id, &industry)
		industries[industry] = id
	}
	return industries, nil
}
