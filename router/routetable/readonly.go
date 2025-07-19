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
		buildKey: Key,
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
		return "", errors.WithMessage(err, "get route table failed")
	}

	return addr, nil
}

func (r *readOnlyRouteTable) BatchGet(ctx context.Context, color string, keys []int64) (addrs []string, err error) {
	keysStr := make([]string, 0, len(keys))
	for _, key := range keys {
		keysStr = append(keysStr, r.buildKey(r.name, color, key))
	}

	addrs, err = r.data.BatchGet(ctx, keysStr)
	if err != nil {
		return nil, errors.WithMessage(err, "batch get route table failed")
	}

	return addrs, nil
}
