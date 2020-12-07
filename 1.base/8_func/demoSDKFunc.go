package main

import (
	"fmt"
)

func main() {
	funcA()
	funcB()
	funcC()
}

type testErr struct {
	code int
	msg  string
}

// panic/recover
// Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
// panic可以在任何地方引发，但recover只有在defer调用的函数中有效。
// recover()必须搭配defer使用。
// defer一定要在可能引发panic的语句之前定义。

func funcA() {
	fmt.Println("func A")
}
func funcB() {
	// 程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。
	// 这个时候我们就可以通过recover将程序恢复回来，继续往后执行。
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}
func funcC() {
	fmt.Println("func C")
}
