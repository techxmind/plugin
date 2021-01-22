// 插件框架 应用于核心流程与扩展的分支流程代码解耦
//
// 核心流程逻辑，可以在任意场景定义一个全局惟一的插件作用域(Scope)
// 通用context及自定义的interface{} message和注册的插件交互 (定义数据协议)
// 分支流程在自己的代码域内，将自己注册进相关的Scope中即可
//
// core:
//    plugin.GetOrCreateScope("my-scope-name").Execute(ctx, msg)
//    //对ctx或msg的副作用数据进行检测，做相应的逻辑分支或读取配置
//
// branch:
//    plugin.GetOrCreateScope("my-scope-name").Plugin().Register("my-plugin", func(ctx context.Context, msg interface{}) {
//        //..do with ctx or msg
//    })
package plugin
