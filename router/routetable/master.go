package routetable

import (
	"context"

	"github.com/go-pantheon/fabrica-util/errors"
)

var _ RouteTable = (*masterRouteTable)(nil)

// masterRouteTable is a basic implementation of the RouteTable interface.
// It provides routing functionality with configurable TTL and key generation.
type masterRouteTable struct {
	ReNewalRouteTable

	data Data
}

// NewMasterRouteTable creates a new BaseRouteTable with the given data store,
// name, key generation function, and optional configuration options.
func NewMasterRouteTable(rtd Data, name string, opts ...Option) *masterRouteTable {
	rt := &masterRouteTable{
		ReNewalRouteTable: NewRenewalRouteTable(rtd, name, opts...),
		data:              rtd,
	}

	return rt
}

// GetSet atomically gets the old value and sets a new value for a routing entry.
func (r *masterRouteTable) GetSet(ctx context.Context, color string, uid int64, addr string) (old string, err error) {
	old, err = r.data.GetSet(ctx, r.BuildKey(color, uid), addr, r.TTL())
	if err != nil {
		return "", errors.Wrapf(err, "getset route table failed. color=%s uid=%d addr=%s", color, uid, addr)
	}

	return old, nil
}

// Set stores a routing entry in the route table with the default TTL.
func (r *masterRouteTable) Set(ctx context.Context, color string, uid int64, addr string) error {
	if err := r.data.Set(ctx, r.BuildKey(color, uid), addr, r.TTL()); err != nil {
		return errors.Wrapf(err, "set route table failed. color=%s uid=%d addr=%s", color, uid, addr)
	}

	return nil
}

// SetNxOrGet sets a routing entry only if it doesn't already exist.
// Returns true if the entry was set, along with the result and any error.
func (r *masterRouteTable) SetNxOrGet(ctx context.Context, color string, uid int64, addr string) (ok bool, result string, err error) {
	ok, result, err = r.data.SetNxOrGet(ctx, r.BuildKey(color, uid), addr, r.TTL())
	if err != nil {
		return false, "", errors.Wrapf(err, "setnx route table failed. color=%s uid=%d addr=%s", color, uid, addr)
	}

	return ok, result, nil
}

// GetEx loads a routing entry and extends its expiration time.
func (r *masterRouteTable) GetEx(ctx context.Context, color string, uid int64) (addr string, err error) {
	addr, err = r.data.GetEx(ctx, r.BuildKey(color, uid), r.TTL())
	if err != nil {
		return "", errors.Wrapf(err, "getex route table failed. color=%s uid=%d", color, uid)
	}

	return addr, nil
}

// // Del deletes a routing entry from the route table.
// func (r *BaseRouteTable) Del(ctx context.Context, color string, uid int64) error {
// 	if err := r.Del(ctx, r.buildKey(r.name, color, uid)); err != nil {
// 		return errors.Wrapf(err, "del route table failed. color=%s uid=%d", color, uid)
// 	}

// 	return nil
// }

// // DelDelay marks a routing entry for delayed deletion after the specified expiration time.
// func (r *BaseRouteTable) DelDelay(ctx context.Context, color string, uid int64, exp time.Duration) error {
// 	if err := r.Expire(ctx, r.buildKey(r.name, color, uid), exp); err != nil {
// 		return errors.Wrapf(err, "expire route table failed. color=%s uid=%d exp=%02fs", color, uid, exp.Seconds())
// 	}

// 	return nil
// }

// // DelIfSame deletes a routing entry only if its current value matches the specified value.
// func (r *BaseRouteTable) DelIfSame(ctx context.Context, color string, uid int64, value string) error {
// 	if err := r.DelIfSame(ctx, r.buildKey(r.name, color, uid), value); err != nil {
// 		return errors.Wrapf(err, "DelIfSame failed. color=%s uid=%d value=%s", color, uid, value)
// 	}

// 	return nil
// }

// // DelIfSame deletes a routing entry only if its current value matches the specified value.
// func (r *BaseRouteTable) DelDelayIfSame(ctx context.Context, color string, uid int64, value string, exp time.Duration) error {
// 	if err := r.ExpireIfSame(ctx, r.buildKey(r.name, color, uid), value, exp); err != nil {
// 		return errors.Wrapf(err, "DelDelayIfSame failed. color=%s uid=%d value=%s exp=%02fs", color, uid, value, exp.Seconds())
// 	}

// 	return nil
// }
