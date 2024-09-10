package cache

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"linkshrink/config"
	"linkshrink/logger"
	"sync"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string) string
	HGet(ctx context.Context, hash, key string) string
	HGetAll(ctx context.Context, hash string) map[string]string
	Del(ctx context.Context, key string) (int64, error)
	Set(ctx context.Context, key string, value string, duration time.Duration) (string, error)
	HSet(ctx context.Context, hash string, key string, value string, duration time.Duration) (int64, error)
}

type RedisCache struct {
	Host          string
	Port          string
	ClusterClient *redis.ClusterClient
	Client        *redis.Client
	ClientLocker  *redislock.Client
	CacheMapLock  sync.Mutex
	Context       context.Context
}

func NewRedisCache(host string, port string) *RedisCache {
	return &RedisCache{Host: host, Port: port}
}

func (r *RedisCache) Connect(ctx context.Context) error {
	log := logger.CreateLoggerWithCtx(ctx)
	log.Infof("connecting to redis, host:%s port:%s", config.REDIS_HOST, config.REDIS_PORT)

	if config.REDIS_CLUSTER {
		r.ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:              []string{fmt.Sprintf("%s:%s", r.Host, r.Port)},
			RouteByLatency:     true,
			MinIdleConns:       10,
			PoolSize:           config.REDIS_POOL_SIZE,
			IdleTimeout:        time.Second * 10,
			IdleCheckFrequency: time.Second * 5,
		})

		r.Context = r.ClusterClient.Context()
		_, err := r.ClusterClient.Ping(r.Context).Result()
		if err != nil {
			log.Errorw("Unable to connect to cluster redis", "error", err)
			return err
		}
	} else {
		r.Client = redis.NewClient(&redis.Options{
			Addr:               fmt.Sprintf("%s:%s", r.Host, r.Port),
			Password:           "",
			DB:                 0,
			MinIdleConns:       10,
			PoolSize:           config.REDIS_POOL_SIZE,
			IdleTimeout:        time.Second * 10,
			IdleCheckFrequency: time.Second * 5,
		})
		r.Context = r.Client.Context()
		_, err := r.Client.Ping(r.Context).Result()
		if err != nil {
			log.Errorw("Unable to connect to single redis", "error", err)
			return err
		}
		r.ClientLocker = redislock.New(r.Client)
	}
	return nil
}
