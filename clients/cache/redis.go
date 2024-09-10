package cache

import (
	"context"
	"encoding/json"
	"linkshrink/config"
	"linkshrink/logger"
	"time"
)

func (r *RedisCache) Get(ctx context.Context, key string) string {
	if config.REDIS_CLUSTER {
		return r.ClusterClient.Get(ctx, key).Val()
	} else {
		return r.Client.Get(ctx, key).Val()
	}
}

func (r *RedisCache) HGet(ctx context.Context, hash, key string) string {
	if config.REDIS_CLUSTER {
		return r.ClusterClient.HGet(ctx, hash, key).Val()
	} else {
		return r.Client.HGet(ctx, hash, key).Val()
	}
}

func (r *RedisCache) HGetAll(ctx context.Context, hash string) map[string]string {
	if config.REDIS_CLUSTER {
		return r.ClusterClient.HGetAll(ctx, hash).Val()
	} else {
		return r.Client.HGetAll(ctx, hash).Val()
	}
}

func (r *RedisCache) Del(ctx context.Context, key string) (int64, error) {
	if config.REDIS_CLUSTER {
		cmd := r.ClusterClient.Del(ctx, key)
		return cmd.Result()
	} else {
		cmd := r.Client.Del(ctx, key)
		return cmd.Result()
	}
}

func (r *RedisCache) Set(ctx context.Context, key, value string, duration time.Duration) (string, error) {
	log := logger.CreateLoggerWithCtx(ctx)
	marshValue, err := json.Marshal(value)
	if err != nil {
		log.Errorf("marshal err: %v", err)
		return "", err
	}
	if config.REDIS_CLUSTER {
		cmd := r.ClusterClient.Set(ctx, key, marshValue, duration)
		return cmd.Result()
	} else {
		cmd := r.Client.Set(ctx, key, marshValue, duration)
		return cmd.Result()
	}
}

func (r *RedisCache) HSet(ctx context.Context, hash, key, value string, duration time.Duration) (int64, error) {
	log := logger.CreateLoggerWithCtx(ctx)
	marshValue, err := json.Marshal(value)
	if err != nil {
		log.Errorf("HSet Error: %s", err.Error())
		return -1, err
	}

	if config.REDIS_CLUSTER {
		cmd := r.ClusterClient.HSet(ctx, hash, key, marshValue, duration)
		return cmd.Result()
	} else {
		cmd := r.Client.HSet(ctx, hash, key, marshValue, duration)
		return cmd.Result()
	}
}
