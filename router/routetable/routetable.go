// Package routetable provides functionality for distributed routing tables
// used for tracking and managing service instances and their connection states.
package routetable

import (
	"context"
	"fmt"
	"time"
)

// RouteTable is an interface for managing routing table entries.
// It provides methods for storing, retrieving, and manipulating route information.
type RouteTable interface {
	ReadOnlyRouteTable

	LoadAndExpire(ctx context.Context, color string, oid int64) (addr string, err error)
	Store(ctx context.Context, color string, key int64, addr string) error
	GetSet(ctx context.Context, color string, key int64, addr string) (old string, err error)
	SetNxOrGet(ctx context.Context, color string, key int64, addr string) (ok bool, result string, err error)
	DelDelay(ctx context.Context, color string, key int64, delay time.Duration) error
	DelIfSame(ctx context.Context, color string, key int64, value string) error
	DelDelayIfSame(ctx context.Context, color string, key int64, value string, delay time.Duration) error
	Del(ctx context.Context, color string, key int64) error
}

// ReadOnlyRouteTable is an interface for read-only access to the routing table.
type ReadOnlyRouteTable interface {
	Load(ctx context.Context, color string, key int64) (addr string, err error)
}

// Data is an interface for the underlying data storage of route tables.
type Data interface {
	Load(ctx context.Context, key string) (addr string, err error)
	LoadAndExpire(ctx context.Context, key string, ttl time.Duration) (addr string, err error)
	SetNxOrGet(ctx context.Context, key, addr string, ttl time.Duration) (set bool, ret string, err error)
	GetSet(ctx context.Context, key, addr string, ttl time.Duration) (old string, err error)
	Set(ctx context.Context, key, addr string, ttl time.Duration) error
	Del(ctx context.Context, key string) error
	DelIfSame(ctx context.Context, key, value string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	ExpireIfSame(ctx context.Context, key, value string, expiration time.Duration) error
}

// NewRouteTable creates a new RouteTable instance with the specified name and data store.
func NewRouteTable(name string, rt Data, opts ...Option) RouteTable {
	return NewBaseRouteTable(rt, name, key, opts...)
}

func key(name, color string, oid int64) string {
	return fmt.Sprintf("r_%s_{%s}_{%d}", name, color, oid)
}
