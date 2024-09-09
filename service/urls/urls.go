package urls

import (
	"context"
	"encoding/base64"
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

func (srv *UrlService) AddUrl(ctx context.Context, url string) (string, error) {
	urlPayload := models.Url{Url: url}
	shortUrl, err := createShortUrl(ctx, url)
	if err != nil {
		return "", err
	}
	urlPayload.ShortUrl = shortUrl
	err = srv.repo.AddUrl(ctx, urlPayload)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func createShortUrl(ctx context.Context, url string) (string, error) {
	dataBytes := []byte(url)
	encoded := base64.StdEncoding.EncodeToString(dataBytes)
	log := logger.CreateLoggerWithCtx(ctx)
	log.Infof(encoded)

	return encoded, nil
}

func (srv *UrlService) GetUrlWithShortUrl(ctx context.Context, shortUrl string) (string, error) {
	url, err := srv.repo.GetUrlWithShortUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}
