package urls

import (
	"context"
	"shrinklink/logger"
	"shrinklink/models"

	"github.com/jmoiron/sqlx"
)

func (repo *UrlsRepo) GetAllUrls(ctx context.Context) ([]models.Url, error) {
	log := logger.CreateLoggerWithCtx(ctx)
	var urls []models.Url
	sqlRow, err := repo.db.DB().Queryx("SELECT * FROM urls ORDER BY id DESC")
	if err != nil {
		log.Errorw("error while fetching urls", "err", err)
		return nil, err
	}
	for sqlRow.Next() {
		var url models.Url
		err = sqlRow.StructScan(&url)
		if err != nil {
			log.Errorw("Error while scanning the row", "err", err)
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (repo *UrlsRepo) AddUrl(ctx context.Context, url models.Url) error {
	_, err := repo.db.DB().Exec(`
			INSERT INTO urls(url, hash) 
			VALUES ($1, $2) 
			ON CONFLICT (url) DO NOTHING
			RETURNING url`, url.Url, url.Hash)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UrlsRepo) GetUrlWithHash(_ context.Context, hash string) (string, error) {
	var url models.Url
	sqlQuery := `SELECT url FROM urls where hash=$1`
	var sqlRow *sqlx.Row

	sqlRow = repo.db.DB().QueryRowx(sqlQuery, hash)
	err := sqlRow.StructScan(&url)

	if err != nil {
		return "", nil
	}

	return url.Url, nil
}
