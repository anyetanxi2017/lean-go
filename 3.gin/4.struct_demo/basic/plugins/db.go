package plugins

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/spf13/viper"
	"lean-go/3.gin/4.struct_demo/basic/config"
	"lean-go/3.gin/4.struct_demo/basic/db"
)

func PluginDb() (err error) {
	// 加载数据库配置
	LoadMysqlConfig()
	for _, config := range config.GConfig.ToolAssembly.Mysql {
		if config.NameSpace == "" && config.Addr != "" {
			log.Warnf("Database config is error [%s]", config.Addr)
			continue
		}
		// 初始化数据库连接
		db.Init(config)
	}
	return
}

func LoadMysqlConfig() {
	configSource := viper.New()
	configSource.AddConfigPath("./conf/" + config.Env)
	configSource.SetConfigName("database")
	configSource.SetConfigType("yml")
	if err := configSource.ReadInConfig(); err != nil {
		panic(err)
	}
	var v map[string]config.Mysql
	if err := configSource.Unmarshal(&v); err != nil {
		panic(err)
	}
	for k, d := range v {
		if k == "" {
			continue
		}
		d.NameSpace = k
		config.GConfig.ToolAssembly.Mysql = append(config.GConfig.ToolAssembly.Mysql, d)
	}
	fmt.Println("数据库配置信息：", config.GConfig.ToolAssembly.Mysql)
}
