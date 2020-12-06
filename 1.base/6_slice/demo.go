package main

import (
	"fmt"
)

func main() {
	// 切片(Slice) 相当于Java的集合，是一个拥有相同类型元素的可变长度的序列。
	// 它是基于数组类型做的一层封装。它非常灵活，支持自动扩容
	// 切片是一个引用类型，它的内部结构包含 地址、长度和容量。切片一般用于快速地操作一块数据集合。
	//demoDefinition() // 定义
	//demoLenAndCap() // 长度和容量
	//demoExpression() // 切片表达式
	//demoMake() //函数构造切片
	//demoCopy() //切片赋值拷贝
	//demoAppend() //添加元素
	demoRemove() // 删除元素

}

func demoDefinition() {
	var a []string           //声明一个字符串切片
	b := []int{}             //声明一个整型切片并初始化
	c := []bool{false, true} // 声明一个布尔切片并初始化
	//var d = []bool{false, true}
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) //false
	//fmt.Println(c == d)   // 切片是引用类型，不支持直接比较，只能和nil比较

}

func demoLenAndCap() {
	var a []int
	for i := 0; i < 17; i++ {
		a = append(a, i)
	}
	fmt.Println(len(a))
	fmt.Println(cap(a))
}

func demoExpression() {
	// 简单切片表达式
	// 切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。
	// 切片表达式中的low和high表示一个索引范围（左包含，右不包含），
	// 也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，
	// 得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s :=a[low:high]
	fmt.Println(s, len(s), cap(s))
	// 为了方便起见，可以省略切片表达式中任何索引。
	fmt.Println(a[2:])  //[3 4 5 ]
	fmt.Println(a[:3])  // [1 2 3]
	fmt.Println(a[:])   // [1 2 3 4 5]
	fmt.Println(a[1:3]) // [2 3]

}

func demoMake() {
	// 我们上面都是基于数组来创建切片，如果需要动态的创建一个切片，就需要使用内置的make()函数
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) // 2
	fmt.Println(cap(a)) // 10
}

func demoCopy() {
	a1 := make([]int, 3)
	a2 := a1 //将 a1 直接赋值给a2 ，a1和a2 共用一个底层数组
	a2[0] = 100
	fmt.Println(a1)
	fmt.Println(a2)

}

func demoAppend() {
	var s []int
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	s2 := []int{5, 6, 7}
	s = append(s, s2...)
	fmt.Println(s)
	s3 := []int{10, 11, 12}
	s = append(s, s3...)
	fmt.Println(s)

}

func demoRemove() {
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	a = append(a[:2], a[3:]...)
	for index, i := range a {
		if i == 37 {
			a = append(a[:index], a[index+1:]...)
		}
	}
	fmt.Println(a)

}
