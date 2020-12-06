package main

import "fmt"

// 相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值
// 常量的声明和变量非常类似，只是把 var 换成了 const，常量在定义时必须赋值

const pi = 3.1415

// 多个常量也可以一起声明
const (
	n1 = 100 //const 同时声明多个变量时，如果省略了值则表示和上面一行的值相同 下面三个值都为100
	n2
	n3
)

// iota 在const关键字出现时将被重置为0。
//const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
//使用iota能简化定义，在定义枚举时很有用。
const (
	m0 = iota //0
	m1        //1
	m2        //2
	m3        //3
)

//iota声明中间插队

const (
	x1 = iota //0
	x2 = 100  //100
	x3 = iota //2
	x4        //3
)

// 多个 iota 定义在一行
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)
const n5 = iota //0
func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(m0, m1, m2, m3)
	fmt.Println(a, b, c, d, e, f)

}
