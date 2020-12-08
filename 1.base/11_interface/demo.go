package main

import (
	"fmt"
)

// 在Go语言中接口是一种类型，一种抽象类型
type Sayer interface {
	say()
}
type dog struct {
}
type cat struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}
func (c cat) say() {
	fmt.Println("喵喵喵")
}
func demoBases() {
	var x Sayer
	x = cat{}
	x.say()
	x = dog{}
	x.say()
}

// --------------------------------------------------------------------------------
// 值接收者和指针接收者实现接口的区别

type Mover interface {
	move()
}

func (d dog) move() {
	fmt.Println("狗会动")
}
func demoMove() {
	var x Mover
	x = dog{}  //x 可以接收值类型
	x.move()   // 狗会动
	x = &dog{} // x 也可以接收指针类型
	x.move()   // 狗会动
	// 从上面的代码中可以发现，使用值 接收者实现接口之后，不管是dog结构体还是指针*dog类型的变量都可以赋值给该接口变量。
	// 因为Go语言中有对指针类型变量求值的语法糖.
}

// --------------------------------------------------------------------------------
// 类型与接口一个特别关系
// 一个接口方法，不一定需要一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现

// 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct {
}

// 甩干器 实现dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

//海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

func (h haier) wash() {
	fmt.Println("洗洗洗")
}
func demoWashingMachine() {
	var x WashingMachine
	x = haier{}
	x.dry()
}

// --------------------------------------------------------------------------------
// 接口嵌套
type animal interface {
	Sayer
	Mover
}
type bird struct {
}

func (b bird) move() {
	fmt.Println("鸟儿飞走了")
}

func (b bird) say() {
	fmt.Println("鸟儿在唱歌")
}
func demo() {
	var animal animal
	animal = bird{}
	animal.move()
	animal.say()
}

// --------------------------------------------------------------------------------
// 空接口是没有定义任何方法的接口。因此任何类型都实现了空接口
// 空接口类型的变量可以任意类型的变量
func demo2() {
	var x interface{}
	x = "Hello "
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100
	fmt.Printf("type:%T value:%v\n", x, x)
	x = false
	fmt.Printf("type:%T value:%v\n", x, x)
}

// 空接口作为函数的参数
// 使用空接口实现可以接收任意的函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 空接口作为map的值
func demo3() {
	var userInfo = make(map[string]interface{})
	userInfo["name"] = "yy"
	userInfo["age"] = 18
	userInfo["married"] = false
	fmt.Println(userInfo)
}

// -----------------------------------------------------------------------------------------------
// 类型断言
// 空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？
// 想要判断接口的值是哪个类型可以使用如下语法
// x.(T)  x：表示类型为interface{}的变量 T：表示x可能是的类型
func demoUseIf() {
	var x interface{}
	x = dog{}
	_, ok := x.(dog)
	if ok {
		fmt.Println("是我是我")
	} else {
		fmt.Println("类型判断失败")
	}
}
func demoUseSwitch() {
	var x interface{}
	x = cat{}
	switch x.(type) {
	case dog:
		fmt.Println("是小狗")
	case string:
		fmt.Println("是字符串")
	case cat:
		fmt.Println("是圆圆家的猫")
	case bird:
		fmt.Println("是小鸟")
	}

}

func main() {
	demoUseSwitch()
	//demoWashingMachine()
	//demoMove()
	//demoBases()
}
