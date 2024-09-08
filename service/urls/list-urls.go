package urls

import (
	"context"
	"urlshortner/logger"
	"urlshortner/models"
)

func (srv *UrlService) ListUrls(ctx context.Context) ([]models.Url, error) {

	log := logger.CreateLoggerWithCtx(ctx)
	urls, err := srv.repo.ListUrls(ctx)
	if err != nil {
		log.Errorw("error listing urls from repo", "error", err)
		return nil, err
	}
	return urls, nil
}
