// Package routetable provides functionality for distributed routing tables
// used for tracking and managing service instances and their connection states.
package routetable

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrRouteTableNotFound     = errors.New("route table not found")
	ErrRouteTableValueNotSame = errors.New("route table not same")
)

type RouteTable interface {
	MasterRouteTable
}

// MasterRouteTable is an interface for managing routing table entries.
// It provides methods for storing, retrieving, and manipulating route information.
type MasterRouteTable interface {
	ReNewalRouteTable

	GetEx(ctx context.Context, color string, oid int64) (addr string, err error)
	Set(ctx context.Context, color string, key int64, addr string) error
	GetSet(ctx context.Context, color string, key int64, addr string) (old string, err error)
	SetNxOrGet(ctx context.Context, color string, key int64, addr string) (ok bool, result string, err error)

	// DelDelay(ctx context.Context, color string, key int64, delay time.Duration) error
	// DelIfSame(ctx context.Context, color string, key int64, value string) error
	// DelDelayIfSame(ctx context.Context, color string, key int64, value string, delay time.Duration) error
	// Del(ctx context.Context, color string, key int64) error
}

// ReNewalRouteTable is an interface for read-only access to the routing table.
type ReNewalRouteTable interface {
	ReadOnlyRouteTable

	RenewSelf(ctx context.Context, color string, key int64, value string) error
	TTL() time.Duration
}

type ReadOnlyRouteTable interface {
	BuildKey(color string, oid int64) string
	Get(ctx context.Context, color string, key int64) (addr string, err error)
}

// Data is an interface for the underlying data storage of route tables.
type Data interface {
	Get(ctx context.Context, key string) (addr string, err error)
	GetEx(ctx context.Context, key string, ttl time.Duration) (addr string, err error)
	Set(ctx context.Context, key, addr string, ttl time.Duration) error
	SetNxOrGet(ctx context.Context, key, addr string, ttl time.Duration) (set bool, ret string, err error)
	GetSet(ctx context.Context, key, addr string, ttl time.Duration) (old string, err error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	ExpireIfSame(ctx context.Context, key, value string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	DelIfSame(ctx context.Context, key, value string) error
}

func key(name, color string, oid int64) string {
	return fmt.Sprintf("r_%s_{%s}_{%d}", name, color, oid)
}
