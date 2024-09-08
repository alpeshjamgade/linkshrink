package urls

import (
	"context"
	"urlshortner/models"
	"urlshortner/repo/urls"
)

type IUrlService interface {
	ListUrls(ctx context.Context) ([]models.Url, error)
}

type UrlService struct {
	repo urls.IUrlsRepo
}

func NewUrlsService(repo urls.IUrlsRepo) *UrlService {
	service := &UrlService{repo: repo}
	return service
}
