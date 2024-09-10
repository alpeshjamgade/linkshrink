package urls

import (
	"context"
	"shrink-link/clients/cache"
	"shrink-link/clients/db"
	"shrink-link/models"

	_ "github.com/lib/pq"
)

type IUrlsRepo interface {
	GetAllUrls(ctx context.Context) ([]models.Url, error)
	AddUrl(ctx context.Context, url models.Url) error
	GetUrlWithHash(ctx context.Context, hash string) (string, error)
}

type UrlsRepo struct {
	db    db.DB
	cache cache.ICache
}

func NewUrlsRepo(db db.DB, cache cache.ICache) *UrlsRepo {
	repo := &UrlsRepo{db: db, cache: cache}
	return repo
}

func (repo *UrlsRepo) GetCache() cache.ICache { return repo.cache }
