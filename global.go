package plugin

import (
	"context"
	"sync"
)

var (
	scopes sync.Map
)

func GetScope(name string) *Scope {
	if v, ok := scopes.Load(name); ok {
		return v.(*Scope)
	}

	return nil
}

func GetOrCreateScope(name string) *Scope {
	scope := GetScope(name)
	if scope != nil {
		return scope
	}

	scope = NewScope(name)
	scopes.Store(name, scope)

	return scope
}

// Execute plugins in specified scope name
func Execute(scopeName string, ctx context.Context, message interface{}) {
	scope := GetScope(scopeName)
	if scope != nil {
		scope.Execute(ctx, message)
	}
}
