package urls

import (
	"context"
	"encoding/base64"
	"shrink-link/logger"
	"shrink-link/models"
)

func (srv *UrlService) GetAllUrls(ctx context.Context) ([]models.Url, error) {

	urls, err := srv.repo.GetAllUrls(ctx)
	if err != nil {
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
