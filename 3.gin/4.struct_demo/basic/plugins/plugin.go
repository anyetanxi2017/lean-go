package plugins

// 插件类型
type PluginRunFunction func() (err error)

// 插件切片
var PluginRunFunctions = &[]PluginRunFunction{}

// 添加插件
func UsePlugin(p ...PluginRunFunction) {
	*PluginRunFunctions = append(*PluginRunFunctions, p...)
}

// 执行每个插件，看是否正常
func ExecPlugin() {
	var err error
	for _, plugin := range *PluginRunFunctions {
		if err = plugin(); err != nil {
			panic(err)
		}
	}
}
