package routetable

import (
	"context"
	"time"
)

const (
	defaultTTL = time.Hour * 24 * 7
)

var _ RouteTable = (*BaseRouteTable)(nil)

type getKeyFunc func(name, color string, oid int64) string

// Option is a function type that configures a BaseRouteTable instance.
type Option func(*BaseRouteTable)

// WithTTL returns an Option that sets the time-to-live for route table entries.
func WithTTL(dur time.Duration) Option {
	return func(r *BaseRouteTable) {
		r.ttl = dur
	}
}

// BaseRouteTable is a basic implementation of the RouteTable interface.
// It provides routing functionality with configurable TTL and key generation.
type BaseRouteTable struct {
	Data

	name   string
	getKey getKeyFunc
	ttl    time.Duration
}

// NewBaseRouteTable creates a new BaseRouteTable with the given data store,
// name, key generation function, and optional configuration options.
func NewBaseRouteTable(rtd Data, name string, getKey getKeyFunc, opts ...Option) *BaseRouteTable {
	rt := &BaseRouteTable{
		Data:   rtd,
		name:   name,
		getKey: getKey,
		ttl:    defaultTTL,
	}
	for _, opt := range opts {
		opt(rt)
	}

	return rt
}

// Store stores a routing entry in the route table with the default TTL.
func (r *BaseRouteTable) Store(ctx context.Context, color string, uid int64, addr string) error {
	return r.Set(ctx, r.getKey(r.name, color, uid), addr, r.ttl)
}

// GetSet atomically gets the old value and sets a new value for a routing entry.
func (r *BaseRouteTable) GetSet(ctx context.Context, color string, uid int64, addr string) (old string, err error) {
	return r.Data.GetSet(ctx, r.getKey(r.name, color, uid), addr, r.ttl)
}

// SetNx sets a routing entry only if it doesn't already exist.
// Returns true if the entry was set, along with the result and any error.
func (r *BaseRouteTable) SetNx(ctx context.Context, color string, uid int64, addr string) (ok bool, result string, err error) {
	return r.Data.SetNx(ctx, r.getKey(r.name, color, uid), addr, r.ttl)
}

// Load retrieves a routing entry from the route table.
func (r *BaseRouteTable) Load(ctx context.Context, color string, uid int64) (addr string, err error) {
	return r.Data.Load(ctx, r.getKey(r.name, color, uid))
}

// LoadAndExpire loads a routing entry and extends its expiration time.
func (r *BaseRouteTable) LoadAndExpire(ctx context.Context, color string, uid int64) (addr string, err error) {
	return r.Data.LoadAndExpire(ctx, r.getKey(r.name, color, uid), r.ttl)
}

// Del deletes a routing entry from the route table.
func (r *BaseRouteTable) Del(ctx context.Context, color string, uid int64) error {
	return r.Data.Del(ctx, r.getKey(r.name, color, uid))
}

// DelDelay marks a routing entry for delayed deletion after the specified expiration time.
func (r *BaseRouteTable) DelDelay(ctx context.Context, color string, uid int64, expiration time.Duration) error {
	return r.Expire(ctx, r.getKey(r.name, color, uid), expiration)
}

// DelIfSame deletes a routing entry only if its current value matches the specified value.
func (r *BaseRouteTable) DelIfSame(ctx context.Context, color string, uid int64, value string) error {
	return r.Data.DelIfSame(ctx, r.getKey(r.name, color, uid), value)
}
