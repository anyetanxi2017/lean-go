package main

import "fmt"

//全局变量 m
var count = 100

func main() {
	// 标准声明  var 变量名 变量类型
	var name string //""
	var age int     //0
	var isOk bool   //false
	fmt.Println(name, age, isOk)

	// 批量声明
	var (
		a string  //""
		b int     // 0
		c bool    //false
		d float32 // 0.000000
	)
	fmt.Println(a, b, c, d)

	// 变量初始化 var 变量名 类型 = 表达式
	var name1 string = "yy"
	var age1 int = 18
	var name2, age2 = "yy", 20
	fmt.Println(name1, age1, name2, age2)

	// 类型推导
	// 有时候我们会将变量的类型省略，这时候编译器会根据等号右边的值来推导变量的类型
	var name3 = "yy"
	var age3 = 18
	fmt.Println(name3, age3)

	// 短变量声明
	// 在函数内部，可以使用更简略的 :=方式声明初始化变量
	n := 10
	m := 200 //此处声明局部变量 m
	fmt.Println(n, m, count)

	// 匿名变量
	// 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量。匿名变量用一个下划线_表示
	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("yy=", y)

	// 注意
	// 1. 函数外的每个语句都必须以关键字开始 var,const func等
	// 2. :=不能使用在函数外
	// 3 _多用于占位，表示忽略值
}
func foo() (int, string) {
	return 10, "yy"

}
