package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 安装
/*
go get github.com/spf13/viper
go get github.com/fsnotify/fsnotify

*/
// 使用viper读取yaml配置文件
func main() {
	//  方式一
	viper.SetConfigName("config_yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		panic(err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更:", e.Name)
	})
	//url := viper.Get("mysql.uri")
	//fmt.Printf("mysql url:%s username:%s password:%s", viper.Get("mysql.url"), viper.Get("mysql.username"), viper.Get("mysql.password"))
	// 方式二
	configSource := viper.New()
	configSource.AddConfigPath(".")
	configSource.SetConfigName("config_yaml")
	configSource.SetConfigType("yml")
	if err := configSource.ReadInConfig(); err != nil {
		panic(err)
	}
	var v map[string]Mysql
	if err := configSource.Unmarshal(&v); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", v["mysql1"])
}

type Mysql struct {
	Url       string `yaml:"url"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}
