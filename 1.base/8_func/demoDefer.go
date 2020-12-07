package main

import "fmt"

// defer 语句
// Go语言中被 defer的语句将会被延迟执行,如果有多个 defer最先被 defer的语句最后执行
// 由于 defer 语句延迟调用的特性，所以 defer语句能非常方便的处理资源释放问题。
// 比如：资源清理、文件关闭、解锁及记录时间等
// https://www.liwenzhou.com/posts/Go/09_function/ 可具体查看 Go语言 defer 执行时机
// Go语言 return 函数语句的底层实现是 返回值=x -> RET指令
// defer 语句执行的时机为  返回值x -? 运行defer -> RET指令
func main() {
	//fmt.Println("start")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("end")
	//start
	//end
	//3
	//2
	//1
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //5
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //6
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //5
}
