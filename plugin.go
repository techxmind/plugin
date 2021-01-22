// mechanism like gokit.endpoint
package plugin

import (
	"context"
)

type Plugin func(ctx context.Context, message interface{})

func Nop(context.Context, interface{}) {}
