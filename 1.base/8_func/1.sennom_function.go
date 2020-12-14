package main

import "fmt"

// 匿名函数使用
/*
匿名函数最大的作用就是模块块级使用域，避免数据污染
 */
func main() {
	// 一般使用
	func() {
		fmt.Println("hello") //hello
	}()

	// 匿名函数传参
	func(msg string) {
		fmt.Println(msg) //hello
	}("hello")
}
