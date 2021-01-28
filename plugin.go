// mechanism like gokit.endpoint
package plugin

import (
	"context"
	"sync"
)

type Plugin func(ctx context.Context, message *Message)

func Nop(context.Context, interface{}) {}

// Message wrap plugin message
type Message struct {
	sync.RWMutex
	Data interface{}
}

var messagePool = sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}
