package repo

import (
	"context"
	"fmt"
	"go-srv/internal/entities"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	insertQuery = `
	insert into response (id, original_text, resp_category) values ($1, $2, $3);
	`

	selectResponsesQuery = `
	select id, original_text, resp_category from responses limit $1 offset $2;
	`

	updateResponseById = `
	update responses set resp_category = $1 where id = $2;
	`

	statTotal      = "select count(*) from responses;"
	statGratitude  = "select count(*) from responses where resp_category = 'gratitude';"
	statSuggestion = "select count(*) from responses where resp_category = 'suggestion';"
	statClaim      = "select count(*) from responses where resp_category = 'claim';"
)

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(
	ctx context.Context,
	host string,
	port int,
	username string,
	password string,
	dbName string,
) (*Repo, error) {
	repoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		dbName,
	)

	var conn *pgxpool.Pool
	var err error
	for i := 0; i < 30; i++ {
		conn, err = pgxpool.Connect(ctx, repoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return &Repo{
				pool: conn,
			}, nil
		}
	}
	return nil, err
}

func (r *Repo) SaveResponses(ctx context.Context, responses []entities.Response) error {
	batch := &pgx.Batch{}
	for _, response := range responses {
		batch.Queue(
			insertQuery,
			response.Id,
			response.OriginalText,
			response.Category,
		)
	}

	br, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer br.Release()

	if err := br.Conn().SendBatch(ctx, batch).Close(); err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetResponses(ctx context.Context, limit int, offset int) ([]entities.Response, error) {
	rows, err := r.pool.Query(ctx, selectResponsesQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	responses := make([]entities.Response, 0)
	for rows.Next() {
		var response entities.Response
		_ = rows.Scan(
			&response.Id,
			&response.OriginalText,
			&response.Category,
		)

		responses = append(responses, response)
	}

	return responses, nil
}

func (r *Repo) UpdateResponse(ctx context.Context, id int, category string) error {
	_, err := r.pool.Exec(ctx, updateResponseById, category, id)
	return err
}

func (r *Repo) GetStatistics(ctx context.Context) (entities.Statistics, error) {
	var stat entities.Statistics
	if err := r.pool.QueryRow(ctx, statTotal).Scan(&stat.Total); err != nil {
		return entities.Statistics{}, err
	}
	if err := r.pool.QueryRow(ctx, statGratitude).Scan(&stat.Gratitudes); err != nil {
		return entities.Statistics{}, err
	}
	if err := r.pool.QueryRow(ctx, statSuggestion).Scan(&stat.Suggestions); err != nil {
		return entities.Statistics{}, err
	}
	if err := r.pool.QueryRow(ctx, statClaim).Scan(&stat.Claims); err != nil {
		return entities.Statistics{}, err
	}
	return stat, nil
}
