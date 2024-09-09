package urls

import (
	"context"
	"urlshortner/clients/cache"
	"urlshortner/clients/db"
	"urlshortner/models"

	_ "github.com/lib/pq"
)

type IUrlsRepo interface {
	ListUrls(ctx context.Context) ([]models.Url, error)
}

type UrlsRepo struct {
	db    db.DB
	cache cache.ICache
}

func NewUrlsRepo(db db.DB, cache cache.ICache) *UrlsRepo {
	repo := &UrlsRepo{db: db, cache: cache}
	return repo
}

func (repo UrlsRepo) GetCache() cache.ICache { return repo.cache }
