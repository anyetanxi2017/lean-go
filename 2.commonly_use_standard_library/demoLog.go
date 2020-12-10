package main

import (
	"fmt"
	"log"
	"os"
)

// log
/*
Go语言内置的log包实现了简单的日志服务。本文介绍了标准库log的基本使用。
*/

// 使用Logger
/*
log包定义了Logger类型，该类型提供了一些格式化输出的方法。
本包也提供了一个预定义的“标准”logger，可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、
和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用。

logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。
Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
*/

func demo30() {
	log.Println("这是一条很普通的日志")
	log.Printf("这是一条%s日志", "很普通的日志")
}

//配置logger
/*
标准logger的配置

默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。
log标准库中为我们提供了定制这些设置的方法。

log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。

*/

// flag选项
/*
log标准库提供了如下的flag选项，它们是一系列定义好的常量。
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
*/
func demo31() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[yy]") //设置前缀
	log.Println("这是一条很普通的日志。")
}

// 配置日志输出位置
/*
SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。

说明：
如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中。
*/

func demo32() {
	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed,err:", err)
		return
	}
	log.SetOutput(file)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志")

}
func demo33() {
	log.Println("加油hx")
}
func main() {
	demo33()
}
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
}
