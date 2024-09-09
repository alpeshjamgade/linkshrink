package urls

import (
	"context"
	"urlshortner/logger"
	"urlshortner/models"
)

func (repo *UrlsRepo) ListUrls(ctx context.Context) ([]models.Url, error) {
	log := logger.CreateLoggerWithCtx(ctx)
	urls := []models.Url{}
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
