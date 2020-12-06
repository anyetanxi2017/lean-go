package main

import (
	"fmt"
	"math"
	"strings"
)

//Go语言中有丰富的数据类型，
//除了基本的整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等。
//Go 语言的基本类型和其他语言大同小异。

//本章学习地址 https://www.liwenzhou.com/posts/Go/02_datatype/
func main() {
	// 整型
	// 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
	v := 123_456 //123456  还允许我们用 _ 来分隔数字，比如说： v := 123_456 表示 v 的值等于 123456。
	fmt.Println(v)
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)

	// 八进制 以0开头
	var b int = 07
	fmt.Printf("%o \n", b)

	// 十六进制 以0x开头
	var c = 0xff
	fmt.Printf("%x \n", c)

	// 浮点
	// Go语言支持两种浮点型数：float32和float64。
	fmt.Println(math.MaxFloat64)

	// 字符串
	// Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
	// Go 语言里的字符串的内部实现使用UTF-8编码。
	// 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符
	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1, s2)

	// 字符串转义符
	// \r	回车符（返回行首）
	// \n	换行符（直接跳到下一行的同列位置）
	// \t	制表符
	// \'	单引号
	// \"	双引号
	// \\	反斜杠
	fmt.Println("\\")
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"") //str := "c:\Code\lesson1\go.exe"

	// 多行字符串
	// Go语言中要定义一个多行字符串时，就必须使用反引号字符：
	s3 := `
第一行
第二行
第三行
`
	fmt.Println(s3)

	// 字符串常用操作
	str := "hello"
	fmt.Println(len(str))                    // 求长度
	sprintf := fmt.Sprintf("hello %s", "yy") //拼接字符串
	fmt.Println(sprintf)
	fmt.Println(strings.Contains(str, "h"))   //判断是否包含
	fmt.Println(strings.HasPrefix(str, "h"))  //判断前缀
	fmt.Println(strings.HasSuffix(str, "lo")) //判断后缀
	fmt.Println(strings.Index(str, "e"))      //子串出现的位置
	split := strings.Split(str, "")           //分割
	fmt.Println(strings.Join(split, ","))     //join操作  h,e,l,l,o

	// 类型转换
	// Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
	a, b = 2, 3
	var num int
	num = int(float64(a*a + b + b))
	fmt.Println(num)

}
