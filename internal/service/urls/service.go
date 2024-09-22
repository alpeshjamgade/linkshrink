package urls

import (
	"context"
	"shrinklink/internal/repo/urls"
)

type IUrlService interface {
	GetAllUrls(ctx context.Context) ([]map[string]string, error)
	AddUrl(ctx context.Context, url string) (string, error)
	GetUrlWithHash(ctx context.Context, hash string) (string, error)
}

type UrlService struct {
	repo urls.IUrlsRepo
}

func NewUrlsService(repo urls.IUrlsRepo) *UrlService {
	service := &UrlService{repo: repo}
	return service
}
