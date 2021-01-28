package plugin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestPlugin(name string) Plugin {
	return func(ctx context.Context, msg *Message) {
		s := msg.Data.(*[]string)
		*s = append(*s, name)
	}
}

func TestRegister(t *testing.T) {
	ast := assert.New(t)

	s := NewScope("test-register")

	s.Plugin().Register("plugin1", createTestPlugin("plugin1"))
	s.Plugin().Register("plugin2", createTestPlugin("plugin2"))
	s.Plugin().Before("plugin2").Register("plugin3", createTestPlugin("plugin3"))
	s.Plugin().After("plugin1").Register("plugin4", createTestPlugin("plugin4"))

	ctx := context.Background()
	msg := make([]string, 0)

	s.Execute(ctx, &msg)
	ast.Equal([]string{"plugin1", "plugin4", "plugin3", "plugin2"}, msg)

	s.Plugin().Replace("plugin4", createTestPlugin("plugin5"))
	s.Plugin().Remove("plugin2")

	msg = make([]string, 0)
	s.Execute(ctx, &msg)
	ast.Equal([]string{"plugin1", "plugin5", "plugin3"}, msg)
}
