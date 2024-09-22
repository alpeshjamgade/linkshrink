package urls

import (
	"context"
	"fmt"
	"shrinklink/config"
	"shrinklink/internal/models"
	"shrinklink/internal/utils"
)

func (srv *UrlService) GetAllUrls(ctx context.Context) ([]map[string]string, error) {

	urls, err := srv.repo.GetAllUrls(ctx)
	result := make([]map[string]string, len(urls))
	for i, url := range urls {
		result[i] = map[string]string{
			"url":       url.Url,
			"short_url": fmt.Sprintf("%s/%s", config.DOMAIN, url.Hash),
		}
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *UrlService) AddUrl(ctx context.Context, url string) (string, error) {
	urlPayload := models.Url{}

	//hash := utils.GenerateHash(url)
	hash := utils.GenerateCRC32Hash(url)

	urlPayload.Url = url
	urlPayload.Hash = hash

	if err := srv.repo.AddUrl(ctx, urlPayload); err != nil {
		return "", err
	}

	shortUrl := fmt.Sprintf("%s/%s", config.DOMAIN, hash)
	return shortUrl, nil
}

func (srv *UrlService) GetUrlWithHash(ctx context.Context, hash string) (string, error) {
	url, err := srv.repo.GetUrlWithHash(ctx, hash)
	if err != nil {
		return "", err
	}
	return url, nil
}
