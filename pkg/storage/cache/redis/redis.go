package redis

import (
	"context"
	"time"

	"github.com/OzkrOssa/isp-hub/pkg/config"
	"github.com/OzkrOssa/isp-hub/pkg/storage/cache"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

/**
 * Redis implements port.CacheRepository interface
 * and provides an access to the redis library
 */
type Redis struct {
	Client *redis.Client
}

// New creates a new instance of Redis
func New(ctx context.Context, config *config.Redis) (cache.CacheRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       0,
	})

	if err := redisotel.InstrumentTracing(client); err != nil {
		return nil, err
	}

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client}, nil
}

// Set stores the value in the redis database
func (r *Redis) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

// Get retrieves the value from the redis database
func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := r.Client.Get(ctx, key).Result()
	bytes := []byte(res)
	return bytes, err
}

// Delete removes the value from the redis database
func (r *Redis) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// DeleteByPrefix removes the value from the redis database with the given prefix
func (r *Redis) DeleteByPrefix(ctx context.Context, prefix string) error {
	var cursor uint64
	var keys []string

	for {
		var err error
		keys, cursor, err = r.Client.Scan(ctx, cursor, prefix, 100).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			err := r.Client.Del(ctx, key).Err()
			if err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

// Close closes the connection to the redis database
func (r *Redis) Close() error {
	return r.Client.Close()
}
