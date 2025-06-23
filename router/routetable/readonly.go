package routetable

import (
	"context"

	"github.com/go-pantheon/fabrica-util/errors"
)

type buildKeyFunc func(name, color string, oid int64) string

var _ ReadOnlyRouteTable = (*readOnlyRouteTable)(nil)

// readOnlyRouteTable is a basic implementation of the ReadOnlyRouteTable interface.
// It provides routing functionality with key generation.
type readOnlyRouteTable struct {
	data     Data
	name     string
	buildKey buildKeyFunc
}

// NewReadOnlyRouteTable creates a new readOnlyRouteTable with the given data store,
// name, key generation function.
func NewReadOnlyRouteTable(rtd Data, name string) *readOnlyRouteTable {
	rt := &readOnlyRouteTable{
		data:     rtd,
		name:     name,
		buildKey: key,
	}

	return rt
}

func (r *readOnlyRouteTable) BuildKey(color string, oid int64) string {
	return r.buildKey(r.name, color, oid)
}

// Get retrieves a routing entry from the route table.
func (r *readOnlyRouteTable) Get(ctx context.Context, color string, uid int64) (addr string, err error) {
	addr, err = r.data.Get(ctx, r.buildKey(r.name, color, uid))
	if err != nil {
		return "", errors.Wrapf(err, "get route table failed. color=%s uid=%d", color, uid)
	}

	return addr, nil
}
