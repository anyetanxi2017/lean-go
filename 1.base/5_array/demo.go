package main

import (
	"fmt"
)

func main() {
	//数组是同一种数据类型元素的集合。
	//在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。

	// 数组定义
	// 定义一个长度为3元素类型为int的数组 默认值为0 0 0
	// 数组可以通过下标进行访问，下标是从0开始，最后一个元素的下标是 len-1
	var a [3]int
	fmt.Println(a)

	// 数组初始化
	// 方法一
	var array [3]int                            //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //会用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定初始值完成初始化
	fmt.Println(array)
	fmt.Println(numArray)
	fmt.Println(cityArray)
	// 方法二
	// 一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
	var a2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a2)
	// 方法三
	// 我们还可以使用指定索引值的方式来初始化数组
	a3 := [...]int{1: 1, 3: 5}
	fmt.Println(a3)

	// 遍历
	a4 := [...]string{"北京", "上海", "深圳"}
	// 方法一
	for i := 0; i < len(a4); i++ {
		fmt.Println(a4[i])
	}
	//方法二
	for index, value := range a4 {
		fmt.Println(index, value)
	}

	// 多维数组  这里以二维为例
	// 定义
	a5 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a5)
	fmt.Println(a5[2][1])
	// 遍历
	for _, v1 := range a5 {
		for _, v2 := range v1 {
			fmt.Printf("%s \t", v2)
		}
		fmt.Println()
	}
	// 多维数组只有第一层可以使用 ... 来让编译器推导数组长度
	//支持的写法
	a6 := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//不支持多维数组的内层使用...
	//a7 := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	fmt.Println(a6)

	// 数组是值类型
	// 赋值和传参会复制整个数组，因此改变副本的值，不会改变本身的值
	update(a6)   //这种方式是不会修改到a6 本身的值
	update2(&a6) //这样才可以修改到

	fmt.Println(a6)
}

func update2(a6 *[3][2]string) {
	a6[0][1] = "2"
}

func update(a6 [3][2]string) {
	a6[0][1] = "2"
}
