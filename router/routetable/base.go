package routetable

import (
	"context"
	"time"

	"github.com/go-pantheon/fabrica-util/errors"
)

const (
	defaultTTL = time.Hour * 24 * 2
)

// Option is a function type that configures a BaseRouteTable instance.
type Option func(*BaseRouteTable)

// WithTTL returns an Option that sets the time-to-live for route table entries.
func WithTTL(dur time.Duration) Option {
	return func(r *BaseRouteTable) {
		r.ttl = dur
	}
}

type buildKeyFunc func(name, color string, oid int64) string

var _ RouteTable = (*BaseRouteTable)(nil)

// BaseRouteTable is a basic implementation of the RouteTable interface.
// It provides routing functionality with configurable TTL and key generation.
type BaseRouteTable struct {
	Data

	name     string
	buildKey buildKeyFunc
	ttl      time.Duration
}

// NewBaseRouteTable creates a new BaseRouteTable with the given data store,
// name, key generation function, and optional configuration options.
func NewBaseRouteTable(rtd Data, name string, buildKey buildKeyFunc, opts ...Option) *BaseRouteTable {
	rt := &BaseRouteTable{
		Data:     rtd,
		name:     name,
		buildKey: buildKey,
		ttl:      defaultTTL,
	}
	for _, opt := range opts {
		opt(rt)
	}

	return rt
}

// Store stores a routing entry in the route table with the default TTL.
func (r *BaseRouteTable) Store(ctx context.Context, color string, uid int64, addr string) error {
	if err := r.Set(ctx, r.buildKey(r.name, color, uid), addr, r.ttl); err != nil {
		return errors.Wrapf(err, "set route table failed. color=%s uid=%d addr=%s", color, uid, addr)
	}

	return nil
}

// GetSet atomically gets the old value and sets a new value for a routing entry.
func (r *BaseRouteTable) GetSet(ctx context.Context, color string, uid int64, addr string) (old string, err error) {
	old, err = r.Data.GetSet(ctx, r.buildKey(r.name, color, uid), addr, r.ttl)
	if err != nil {
		return "", errors.Wrapf(err, "getset route table failed. color=%s uid=%d addr=%s", color, uid, addr)
	}

	return old, nil
}

// SetNxOrGet sets a routing entry only if it doesn't already exist.
// Returns true if the entry was set, along with the result and any error.
func (r *BaseRouteTable) SetNxOrGet(ctx context.Context, color string, uid int64, addr string) (ok bool, result string, err error) {
	ok, result, err = r.Data.SetNxOrGet(ctx, r.buildKey(r.name, color, uid), addr, r.ttl)
	if err != nil {
		return false, "", errors.Wrapf(err, "setnx route table failed. color=%s uid=%d addr=%s", color, uid, addr)
	}

	return ok, result, nil
}

// Load retrieves a routing entry from the route table.
func (r *BaseRouteTable) Load(ctx context.Context, color string, uid int64) (addr string, err error) {
	addr, err = r.Data.Load(ctx, r.buildKey(r.name, color, uid))
	if err != nil {
		return "", errors.Wrapf(err, "load route table failed. color=%s uid=%d", color, uid)
	}

	return addr, nil
}

// LoadAndExpire loads a routing entry and extends its expiration time.
func (r *BaseRouteTable) LoadAndExpire(ctx context.Context, color string, uid int64) (addr string, err error) {
	addr, err = r.Data.LoadAndExpire(ctx, r.buildKey(r.name, color, uid), r.ttl)
	if err != nil {
		return "", errors.Wrapf(err, "loadandexpire route table failed. color=%s uid=%d", color, uid)
	}

	return addr, nil
}

// Del deletes a routing entry from the route table.
func (r *BaseRouteTable) Del(ctx context.Context, color string, uid int64) error {
	if err := r.Data.Del(ctx, r.buildKey(r.name, color, uid)); err != nil {
		return errors.Wrapf(err, "del route table failed. color=%s uid=%d", color, uid)
	}

	return nil
}

// DelDelay marks a routing entry for delayed deletion after the specified expiration time.
func (r *BaseRouteTable) DelDelay(ctx context.Context, color string, uid int64, exp time.Duration) error {
	if err := r.Expire(ctx, r.buildKey(r.name, color, uid), exp); err != nil {
		return errors.Wrapf(err, "expire route table failed. color=%s uid=%d exp=%02fs", color, uid, exp.Seconds())
	}

	return nil
}

// DelIfSame deletes a routing entry only if its current value matches the specified value.
func (r *BaseRouteTable) DelIfSame(ctx context.Context, color string, uid int64, value string) error {
	if err := r.Data.DelIfSame(ctx, r.buildKey(r.name, color, uid), value); err != nil {
		return errors.Wrapf(err, "DelIfSame failed. color=%s uid=%d value=%s", color, uid, value)
	}

	return nil
}

// DelIfSame deletes a routing entry only if its current value matches the specified value.
func (r *BaseRouteTable) DelDelayIfSame(ctx context.Context, color string, uid int64, value string, exp time.Duration) error {
	if err := r.ExpireIfSame(ctx, r.buildKey(r.name, color, uid), value, exp); err != nil {
		return errors.Wrapf(err, "DelDelayIfSame failed. color=%s uid=%d value=%s exp=%02fs", color, uid, value, exp.Seconds())
	}

	return nil
}
