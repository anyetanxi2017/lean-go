package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	//testGlobalVar()
	//fmt.Println(num)
	var c calculation
	c = add
	fmt.Println(c(1, 2)) // 像调用add 一样调用c

	f := c
	fmt.Println(f(1, 2)) //像调用add一样调用f

	// 将函数作为参数
	res := calc2(1, 2, add)
	fmt.Println(res)

	//函数作为返回值
	f2, _ := do("+")
	fmt.Println(f2(1, 2))

	f2, _ = do("-")
	fmt.Println(f2(1, 2))
	// 匿名函数
	demoAnonymousFunc()
	// 闭包
	demoClosure()
}

// 全局变量
// 全局变量是定义在函数外部的量变，它在程序整个运行周期内都有效。

// 定义全局变量num
var num int64 = 10

func testGlobalVar() {
	num = 12 // 如果局部变量和全局变量重名，优先访问局部变更
	fmt.Println(num)
}

// 定义函数类型
// 我们可以使用 type关键字来定义一个函数类型，具体格式如下
type calculation func(int, int) int

// 上面语句定义了一个 calculation 类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值
// 简单来说，凡是满足这个条件的函数都是 calculation 类型的函数，例如下面的 add 和 sub 都是 calculation 类型

func add(x, y int) int { //add和sub 都能赋值给 calculation 类型的变量
	return x + y
}
func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func calc2(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 匿名函数
// 匿名函数多用于实现回调函数和闭包
func demoAnonymousFunc() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) //通过变量调用匿名函数
}

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说， 闭包=函数+引用环境。
func demoClosure() {
	//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。
	// 示例1
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(10)) //20
	fmt.Println(f(10)) //30

	f1 := adder()
	fmt.Println(f1(20)) //20
	fmt.Println(f1(20)) //40
	fmt.Println("-------")

	// 示例2
	f2 := adder2(1)    //初始化x
	fmt.Println(f2(1)) //2
	fmt.Println(f2(1)) //3
	fmt.Println(f2(1)) //4

	// 示例3
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test.jpg"))
	fmt.Println(txtFunc("test"))

	// 示例4
	f3, f4 := calc3(10)
	fmt.Println(f3(1), f4(1))
	fmt.Println(f3(1), f4(1))
	fmt.Println(f3(1), f4(1))

}

// 这个函数返回的是一个闭包函数，该闭包函数保留了对变量x的引用，在调用期间内x变量值会被叠加

// 闭包进阶 示例1
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包进阶 示例2
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包进阶 示例3
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包进阶 示例4
func calc3(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
