// Package redis provides a Redis-based implementation of the route table.
package redis

import (
	"context"
	"time"

	"github.com/go-pantheon/fabrica-kit/router/routetable"
	"github.com/go-pantheon/fabrica-kit/xerrors"
	"github.com/go-pantheon/fabrica-util/errors"
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

// Get retrieves a value from Redis by key.
func (r *RouteTable) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", xerrors.ErrRouteTableNotFoundFunc(key)
	}

	return val, errors.Wrapf(err, "key=%s", key)
}

// GetEx loads a value and resets its expiration time.
func (r *RouteTable) GetEx(ctx context.Context, key string, exp time.Duration) (string, error) {
	val, err := r.client.GetEx(ctx, key, exp).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", xerrors.ErrRouteTableNotFoundFunc(key)
		}

		return "", errors.Wrapf(err, "key=%s", key)
	}

	return val, nil
}

// GetSet atomically sets a new value and returns the old value.
func (r *RouteTable) GetSet(ctx context.Context, key, val string, expire time.Duration) (string, error) {
	cmd := r.client.GetSet(ctx, key, val)
	if err := cmd.Err(); err != nil {
		return "", errors.Wrapf(err, "key=%s", key)
	}

	if err := r.client.Expire(ctx, key, expire).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return "", xerrors.ErrRouteTableNotFoundFunc(key)
		}

		return "", err
	}

	return cmd.Val(), nil
}

// Set stores a key-value pair in Redis with an expiration time.
func (r *RouteTable) Set(ctx context.Context, key, val string, expire time.Duration) error {
	if err := r.client.Set(ctx, key, val, expire).Err(); err != nil {
		return errors.Wrapf(err, "key=%s val=%s expire=%s", key, val, expire)
	}

	return nil
}

// SetNxOrGet sets a key-value pair only if the key does not exist.
func (r *RouteTable) SetNxOrGet(ctx context.Context, key, val string, expire time.Duration) (bool, string, error) {
	ok, err := r.client.SetNX(ctx, key, val, expire).Result()
	if err != nil {
		return false, "", errors.Wrapf(err, "setnx route table failed. key=%s val=%s expire=%s", key, val, expire)
	}

	if !ok {
		v, err := r.client.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return false, "", xerrors.ErrRouteTableNotFoundFunc(key)
		}

		return false, v, errors.Wrapf(err, "key=%s", key)
	}

	return true, val, nil
}

// Expire sets an expiration time for a key.
func (r *RouteTable) Expire(ctx context.Context, key string, expire time.Duration) error {
	if err := r.client.Expire(ctx, key, expire).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return xerrors.ErrRouteTableNotFoundFunc(key)
		}

		return errors.Wrapf(err, "key=%s expire=%s", key, expire)
	}

	return nil
}

func (r *RouteTable) ExpireIfSame(ctx context.Context, key, expect string, expire time.Duration) error {
	txf := func(tx *redis.Tx) error {
		v, err := tx.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return xerrors.ErrRouteTableNotFoundFunc(key)
		}

		if err != nil {
			return errors.Wrapf(err, "key=%s", key)
		}

		if v != expect {
			return xerrors.ErrRouteTableValueNotSameFunc(key, expect)
		}

		_, err = tx.Expire(ctx, key, expire).Result()

		return errors.Wrapf(err, "key=%s", key)
	}

	if err := r.client.Watch(ctx, txf, key); err != nil {
		return errors.Wrapf(err, "key=%s", key)
	}

	return nil
}

// Del deletes a key from Redis.
func (r *RouteTable) Del(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}

		return errors.Wrapf(err, "key=%s", key)
	}

	return nil
}

// DelIfSame deletes a key only if its current value matches the specified value.
func (r *RouteTable) DelIfSame(ctx context.Context, key, expect string) error {
	txf := func(tx *redis.Tx) error {
		v, err := tx.Get(ctx, key).Result()
		if errors.Is(err, redis.Nil) {
			return xerrors.ErrRouteTableNotFoundFunc(key)
		}

		if err != nil {
			return errors.Wrapf(err, "key=%s", key)
		}

		if v != expect {
			return xerrors.ErrRouteTableValueNotSameFunc(key, expect)
		}

		_, err = tx.Del(ctx, key).Result()

		return errors.Wrapf(err, "key=%s", key)
	}

	if err := r.client.Watch(ctx, txf, key); err != nil {
		return errors.Wrapf(err, "key=%s", key)
	}

	return nil
}
