package plugin_test

import (
	"context"
	"fmt"

	"github.com/techxmind/config"
	"github.com/techxmind/plugin"
)

func Example() {
	var (
		ctx = context.Background()
		msg = config.NewMapConfig(nil, false)
	)

	// Get or Create a plugin scope
	scope := plugin.GetOrCreateScope("example")
	// Register plugins in the scope "example"
	scope.Plugin().Register("plugin1", myPlugin1)
	scope.Plugin().Before("plugin1").Register("plugin2", myPlugin2)
	scope.Plugin().After("plugin2").Register("plugin3", myPlugin3)

	// Execute plugins in the scope "example"
	plugin.Execute("example", ctx, msg)

	fmt.Println(msg.String("protocal.something")) //plugin1-result
	fmt.Println(msg.Int("protocal.other1.a"))     //1
	fmt.Println(msg.Int("protocal.other2.c"))     //2
	fmt.Println(msg.Int("protocal.other3.e"))     //4

	// Output:
	// my plugin2!
	// my plugin3!
	// my plugin1!
	// plugin1-result
	// 1
	// 2
	// 4
}

func myPlugin1(ctx context.Context, msg *plugin.Message) {
	msg.Lock()
	defer msg.Unlock()
	if m, ok := msg.Data.(*config.MapConfig); ok {
		m.Set("protocal.something", "plugin1-result")
		m.Set("protocal.other1", map[string]interface{}{
			"a": 1,
			"b": 2,
		})
	}
	fmt.Println("my plugin1!")
}

func myPlugin2(ctx context.Context, msg *plugin.Message) {
	msg.Lock()
	defer msg.Unlock()
	if m, ok := msg.Data.(*config.MapConfig); ok {
		m.Set("protocal.something", "plugin2-result")
		m.Set("protocal.other2", map[string]interface{}{
			"c": 2,
			"d": 4,
		})
	}
	fmt.Println("my plugin2!")
}

func myPlugin3(ctx context.Context, msg *plugin.Message) {
	msg.Lock()
	defer msg.Unlock()
	if m, ok := msg.Data.(*config.MapConfig); ok {
		m.Set("protocal.something", "plugin3-result")
		m.Set("protocal.other3", map[string]interface{}{
			"e": 4,
			"f": 8,
		})
	}
	fmt.Println("my plugin3!")
}
