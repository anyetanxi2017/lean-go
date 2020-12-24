package config

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var (
	GConfig Config
	Env     string
)

// 初始化环境
func InitEnv() {
	if Env = os.Getenv(ENVKey); Env == "" {
		Env = LocalEnv
	}
	log.Println("current env:", Env)
}

const (
	ENVKey = "ENV"
	PodId  = "PODID"

	LocalEnv     = "local"     // 本地
	BxAliTestEnv = "bxalitest" // 蓝缤阿里云测试环境
	BxPreEnv     = "bxpre"     // 预发布环境
	BxProdEnv    = "bxprod"    // 生产环境
)

type Config struct {
	WxDomain   string `json:"wx_domain"`
	WxMiniId   string `json:"wx_mini_id"`
	WxMiniType string `json:"wx_mini_type"`
	//WxAppSecret  string       `json:"wx_app_secret"`
	ToolAssembly   ToolAssembly `json:"tool_assembly" yaml:"tool_assembly"`
	DuibaAppKey    string       `json:"duiba_app_key"`
	DuibaAppSecret string       `json:"duiba_app_secret"`
	DuibaUrl       string       `json:"duiba_url"`
	ScUrl          string       `json:"sc_url"` // 神策
}

//系统工具组件
type ToolAssembly struct {
	Mysql []Mysql `json:"mysql" yaml:"mysql"`
	Redis Redis   `json:"redis" yaml:"redis"`
	Es    Es      `json:"es" yaml:"es"`
}

type Mysql struct {
	NameSpace    string `json:"name_space"`
	Addr         string `json:"addr" yaml:"addr"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"`
}

type Redis struct {
	NameSpace    string   `json:"name_space"`
	Addr         []string `json:"addr" yaml:"addr"`
	DB           int      `json:"db" yaml:"db"`
	Password     string   `json:"password" yaml:"password"`
	PoolSize     int      `json:"poolSize" yaml:"poolSize"`
	MinIdleConns int      `json:"minIdleConns" yaml:"minIdleConns"`
}

type Es struct {
	Addrs    []string `json:"addr1" yaml:"addr1"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
}

// 根据不同环境进行配置
func GetLfsHook(stats string) *lfshook.LfsHook {
	podId := os.Getenv(PodId)
	logFilePath := "/logs/"
	logFileName := podId + "." + stats + "."
	if Env == LocalEnv {
		logFilePath = "./logs/"
	}
	// 日誌檔案
	fileName := path.Join(logFilePath, logFileName)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+"%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(-1),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println("日志输出错误了：", err)
	}
	//lfshook 是一个日志框架
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return lfHook
}
