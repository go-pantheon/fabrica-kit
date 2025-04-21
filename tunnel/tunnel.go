// Package tunnel provides interfaces and functionality for managing communication
// tunnels between services, supporting messaging, routing and forwarding capabilities.
package tunnel

import (
	"context"

	"github.com/go-pantheon/fabrica-util/xsync"
)

// Holder is an interface that combines Pusher functionality with the ability
// to retrieve a specific Tunnel instance.
type Holder interface {
	Pusher
	Tunnel(ctx context.Context, key int32, oid int64) (Tunnel, error)
}

// Pusher is an interface that provides the ability to push data through a tunnel.
type Pusher interface {
	Push(ctx context.Context, pack []byte) error
}

// Worker is an interface that combines tunnel holding, pushing,
// and lifecycle management capabilities.
type Worker interface {
	xsync.Stoppable
	xsync.CountdownStopper
	Holder
	Pusher
}

// Tunnel is an interface for a communication channel that can
// push messages and forward specialized messages.
type Tunnel interface {
	xsync.Stoppable
	Pusher

	Type() int32
	Forward(ctx context.Context, msg ForwardMessage) error
}

// ForwardMessage is an interface for messages that can be forwarded
// through a tunnel with module, sequence, object ID, and payload data.
type ForwardMessage interface {
	GetMod() int32
	GetSeq() int32
	GetObj() int64
	GetData() []byte
}
