package urls

import (
	"context"
	"urlshortner/models"
)

type IUrlsRepo interface {
	ListUrls(ctx context.Context) ([]models.Url, error)
}

type UrlsRepo struct {
}

func NewUrlsRepo() *UrlsRepo {
	repo := &UrlsRepo{}
	return repo
}
