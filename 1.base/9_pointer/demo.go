package main

import "fmt"

func main() {
	// 指针地址和指针类型
	// &取地址  *根据地址取值
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)              // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T value:%d\n", b, b, *b) // b:0xc00001a078 type:*int value:10
	fmt.Println(&b)                                 // 0xc00000e018

	// 指针传值示例
	num := 10
	modify1(num)
	fmt.Println(num) //10
	modify2(&num)
	fmt.Println(num) //100

	// new 和 make
	demoNew()

	// make
	demoMake()

}

// 指针传值示例
func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100
}

// new 示例
func demoNew() {
	// new 是一个内置函数，不太常用，new函数只接受一个参数，这个参数是一个类型
	// new函数返回一个指向该类型内存地址的指针,并且该指针对应的值为该类型的零值
	a := new(int)
	fmt.Println(a)  //0xc0000120d0
	fmt.Println(*a) //0
}

// make 示例
func demoMake() {
	// make用于内存分配，它只用于 slice ,map 和 channel 的内存创建，返回的类型就是这三个类型本身，而不是他们的指针类型
	// make函数是无可替代的，我们在使用slice、map channel 的时候，都需要使用make进行初始化
	var b map[string]int
	b = make(map[string]int, 10)
	b["yy"] = 20
	fmt.Println(b)

}
