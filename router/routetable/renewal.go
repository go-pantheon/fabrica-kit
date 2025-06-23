package routetable

import (
	"context"
	"time"

	"github.com/go-pantheon/fabrica-util/errors"
)

const (
	defaultTTL = time.Hour * 24
)

// Option is a function type that configures a BaseRouteTable instance.
type Option func(*renewalRouteTable)

// WithTTL returns an Option that sets the time-to-live for route table entries.
// dur must be greater than 0, otherwise the default TTL will be used.
func WithTTL(dur time.Duration) Option {
	return func(r *renewalRouteTable) {
		if dur <= 0 {
			dur = defaultTTL
		}

		r.ttl = dur
	}
}

var _ ReNewalRouteTable = (*renewalRouteTable)(nil)

// renewalRouteTable is a basic implementation of the ReNewalRouteTable interface.
// It provides routing functionality with configurable TTL and key generation.
type renewalRouteTable struct {
	ReadOnlyRouteTable

	data Data
	ttl  time.Duration
}

// NewRenewalRouteTable creates a new renewalRouteTable with the given data store,
// name, key generation function, and optional configuration options.
func NewRenewalRouteTable(rtd Data, name string, opts ...Option) *renewalRouteTable {
	rt := &renewalRouteTable{
		ReadOnlyRouteTable: NewReadOnlyRouteTable(rtd, name),
		data:               rtd,
		ttl:                defaultTTL,
	}
	for _, opt := range opts {
		opt(rt)
	}

	return rt
}

func (r *renewalRouteTable) TTL() time.Duration {
	return r.ttl
}

// RenewSelf expires a routing entry only if its current value matches the specified value.
func (r *renewalRouteTable) RenewSelf(ctx context.Context, color string, uid int64, value string) error {
	if err := r.data.ExpireIfSame(ctx, r.BuildKey(color, uid), value, r.ttl); err != nil {
		return errors.Wrapf(err, "renewIfSame route table failed. color=%s uid=%d value=%s", color, uid, value)
	}

	return nil
}
