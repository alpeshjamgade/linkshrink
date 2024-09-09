package urls

import (
	"context"
	"github.com/jmoiron/sqlx"
	"urlshortner/logger"
	"urlshortner/models"
)

func (repo *UrlsRepo) ListUrls(ctx context.Context) ([]models.Url, error) {
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

func (repo *UrlsRepo) AddUrl(_ context.Context, url models.Url) error {
	_, err := repo.db.DB().Exec("INSERT INTO urls(url, short_url) VALUES ($1, $2)", url.Url, url.ShortUrl)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UrlsRepo) GetUrlWithShortUrl(_ context.Context, shortUrl string) (string, error) {
	var url models.Url
	sqlQuery := `SELECT url FROM urls where short_url=$1`
	var sqlRow *sqlx.Row

	sqlRow = repo.db.DB().QueryRowx(sqlQuery, shortUrl)
	err := sqlRow.StructScan(&url)

	if err != nil {
		return "", nil
	}

	return url.Url, nil
}
