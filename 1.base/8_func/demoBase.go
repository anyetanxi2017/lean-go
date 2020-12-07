package main

import "fmt"

// 函数是组织好的，可重复使用的、用于执行指定任务的代码块。
// Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于”一等公民“
//func 函数名(参数)(返回值){
//    函数体
//}
// 字母不能是数字。在同一个包内，函数名也称不能重名
//  返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
func main() {
	//intSum2(1, 2, 3)
	//intSum3("做一个加法", 2, 3)
	//sum, sub := calc(1, 2)
	//fmt.Println(sum, sub)
	strings := someFunc("")
	fmt.Println(strings == nil) //true

}

func sayHello() {
	fmt.Println("hello")
}
func intSum(x int, y int) int {
	return x + y
}

// 可变参数1
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
	return sum
}

// 可变参数2
func intSum3(name string, y ...int) int {
	fmt.Println(name, y)
	sum := 0
	for _, v := range y {
		sum += v
	}
	fmt.Println(sum)
	return sum
}

// 多返回值
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 返回值补充
func someFunc(x string) []string {
	if x == "" {
		return nil
	}
	strings := make([]string, 0, 10)
	return strings
}
