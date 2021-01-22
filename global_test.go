package plugin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrCreateScope(t *testing.T) {
	ast := assert.New(t)

	name := "test-get-or-create"
	s1 := GetOrCreateScope(name)
	ast.NotNil(s1)

	s2 := GetOrCreateScope(name)
	ast.NotNil(s2)

	ast.Equal(s1, s2)
}

func TestExecute(t *testing.T) {
	ast := assert.New(t)

	Execute("not-exists-scope", context.Background(), nil)

	name := "test-execute"
	s := GetOrCreateScope(name)

	s.Plugin().Register("plugin1", createTestPlugin("plugin1"))
	s.Plugin().Register("plugin2", createTestPlugin("plugin2"))
	s.Plugin().Before("plugin2").Register("plugin3", createTestPlugin("plugin3"))
	s.Plugin().After("plugin1").Register("plugin4", createTestPlugin("plugin4"))

	ctx := context.Background()
	msg := make([]string, 0)

	Execute(name, ctx, &msg)
	ast.Equal([]string{"plugin1", "plugin4", "plugin3", "plugin2"}, msg)

	s.Plugin().Replace("plugin4", createTestPlugin("plugin5"))
	s.Plugin().Remove("plugin2")

	msg = make([]string, 0)
	Execute(name, ctx, &msg)
	ast.Equal([]string{"plugin1", "plugin5", "plugin3"}, msg)
}
