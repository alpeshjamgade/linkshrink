package urls

import (
	"context"
	"linkshrink/models"
	"linkshrink/repo/urls"
)

type IUrlService interface {
	GetAllUrls(ctx context.Context) ([]models.Url, error)
	AddUrl(ctx context.Context, url string) (string, error)
	GetUrlWithShortUrl(ctx context.Context, shortUrl string) (string, error)
}

type UrlService struct {
	repo urls.IUrlsRepo
}

func NewUrlsService(repo urls.IUrlsRepo) *UrlService {
	service := &UrlService{repo: repo}
	return service
}
