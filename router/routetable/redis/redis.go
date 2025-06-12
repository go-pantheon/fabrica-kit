// Package redis provides a Redis-based implementation of the route table.
package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/redis/go-redis/v9"
)

var _ routetable.Data = (*RouteTable)(nil)

// RouteTable implements the routetable.Data interface using Redis.
type RouteTable struct {
	client redis.UniversalClient
}

// New creates a new Redis-based route table data store.
func New(client redis.UniversalClient) *RouteTable {
	return &RouteTable{
		client: client,
	}
}

// Set stores a key-value pair in Redis with an expiration time.
func (r *RouteTable) Set(ctx context.Context, key, val string, expire time.Duration) error {
	return r.client.Set(ctx, key, val, expire).Err()
}

// Load retrieves a value from Redis by key.
func (r *RouteTable) Load(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Del deletes a key from Redis.
func (r *RouteTable) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// GetSet atomically sets a new value and returns the old value.
func (r *RouteTable) GetSet(ctx context.Context, key, val string, expire time.Duration) (string, error) {
	cmd := r.client.GetSet(ctx, key, val)
	if err := cmd.Err(); err != nil {
		return "", err
	}

	if err := r.client.Expire(ctx, key, expire).Err(); err != nil {
		return "", err
	}

	return cmd.Val(), nil
}

// SetNxOrGet sets a key-value pair only if the key does not exist.
func (r *RouteTable) SetNxOrGet(ctx context.Context, key, val string, expire time.Duration) (bool, string, error) {
	ok, err := r.client.SetNX(ctx, key, val, expire).Result()
	if err != nil {
		return false, "", err
	}

	if !ok {
		v, err := r.client.Get(ctx, key).Result()
		return false, v, err
	}

	return true, val, nil
}

// LoadAndExpire loads a value and resets its expiration time.
func (r *RouteTable) LoadAndExpire(ctx context.Context, key string, exp time.Duration) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	if err := r.client.Expire(ctx, key, exp).Err(); err != nil {
		return "", err
	}

	return val, nil
}

// Expire sets an expiration time for a key.
func (r *RouteTable) Expire(ctx context.Context, key string, expire time.Duration) error {
	return r.client.Expire(ctx, key, expire).Err()
}

func (r *RouteTable) ExpireIfSame(ctx context.Context, key, expect string, expire time.Duration) error {
	txf := func(tx *redis.Tx) error {
		v, err := tx.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil
		}

		if err != nil {
			return err
		}

		if v != expect {
			return nil
		}

		_, err = tx.Expire(ctx, key, expire).Result()

		return err
	}

	// NOTE: Some Redis client implementations might not support TxPipeline directly.
	// In those cases, an error will be returned and you may need to
	// implement a custom transaction mechanism or use a different approach.
	client, ok := r.client.(*redis.Client)
	if !ok {
		// Fallback for non-redis.Client types
		val, err := r.client.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil
		}

		if err != nil {
			return err
		}

		if val != expect {
			return nil
		}

		return r.client.Del(ctx, key).Err()
	}

	return client.Watch(ctx, txf, key)
}

// DelIfSame deletes a key only if its current value matches the specified value.
func (r *RouteTable) DelIfSame(ctx context.Context, key, expect string) error {
	txf := func(tx *redis.Tx) error {
		v, err := tx.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil
		}

		if err != nil {
			return err
		}

		if v != expect {
			return nil
		}

		_, err = tx.Del(ctx, key).Result()

		return err
	}

	// NOTE: Some Redis client implementations might not support TxPipeline directly.
	// In those cases, an error will be returned and you may need to
	// implement a custom transaction mechanism or use a different approach.
	client, ok := r.client.(*redis.Client)
	if !ok {
		// Fallback for non-redis.Client types
		val, err := r.client.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return nil
		}

		if err != nil {
			return err
		}

		if val != expect {
			return nil
		}

		return r.client.Del(ctx, key).Err()
	}

	return client.Watch(ctx, txf, key)
}
