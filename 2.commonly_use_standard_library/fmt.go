package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// fmt标准库是我们在学习Go语言过程中接触最早最频繁的一个了，本文介绍了fmtb包的一些常用函数。
// 向外输出
// Print
/*
Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，
Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。
*/
func demoPrint() {
	fmt.Print("在终端打印该信息")
	fmt.Printf("我是:%s\n", "yy")
	fmt.Println("在终端打印单独一行显示")
}

// Fprint
/*
Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
*/
func demoFprint() {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	msg := "ok"
	_, _ = fmt.Fprintf(file, "信文件中写信息:%s", msg)
}

// Sprint
/*
	Sprint系列函数会把传入的数据生成并返回一个字符串。
*/
func demoSprint() {
	s1 := fmt.Sprintf("yy")
	name := "yy"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("yyyy")
	fmt.Println(s1, s2, s3)
}

// Errorf
/*
Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
通常使用这种方式来自定义错误类型
*/
func demoErrorf() {
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)
}

// ---------------------------------------------------------------------------------------------------------
// 格式化占位符
// 通用占位符
/*
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	打印值的类型
%%	百分号
*/
func demo1() {
	fmt.Printf("%v\n", 100)   // 100
	fmt.Printf("%v\n", false) // false
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)  // {小王子}
	fmt.Printf("%#v\n", o) // struct { name string }{name:"小王子"}
	fmt.Printf("%T\n", o)  // struct { name string }
	fmt.Printf("100%%\n")  // 100%
}

// 布尔型
// %t	true或false

// 整型
/*
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于”U+%04X”
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
*/
func demo2() {
	n := 65
	fmt.Printf("%b\n", n) // 1000001
	fmt.Printf("%c\n", n) // A
	fmt.Printf("%d\n", n) // 65
	fmt.Printf("%o\n", n) // 101
	fmt.Printf("%x\n", n) // 41
	fmt.Printf("%X\n", n) // 41
}

// 浮点数与复数
/*
占位符	说明
%b	无小数部分、二进制指数的科学计数法，如-123456p-78
%e	科学计数法，如-1234.456e+78
%E	科学计数法，如-1234.456E+78
%f	有小数部分但无指数部分，如123.456
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
*/

func demo3() {
	f := 12.34
	fmt.Printf("%b\n", f) // 6946802425218990p-49
	fmt.Printf("%e\n", f) // 1.234000e+01
	fmt.Printf("%E\n", f) // 1.234000E+01
	fmt.Printf("%f\n", f) // 12.340000
	fmt.Printf("%g\n", f) // 12.34
	fmt.Printf("%G\n", f) // 12.34
}

// 字符串和[]byte
/*

占位符	说明
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f
%X	每个字节用两字符十六进制数表示（使用A-F）
*/
func demo4() {
	s := "小王子"
	fmt.Printf("%s\n", s) // 小王子
	fmt.Printf("%q\n", s) // "小王子"
	fmt.Printf("%x\n", s) // e5b08fe78e8be5ad90
	fmt.Printf("%X\n", s) // E5B08FE78E8BE5AD90
}

// 指针
/*
占位符	说明
%p	表示为十六进制，并加上前导的0x
*/
func demo5() {
	a := 10
	fmt.Printf("%p\n", &a)  // 0xc000094000
	fmt.Printf("%#p\n", &a) // c000094000
}

// 宽度标识符
/*
宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。
精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；
如果点号后没有跟数字，表示精度为0。
占位符	说明
%f	默认宽度，默认精度
%9f	宽度9，默认精度
%.2f	默认宽度，精度2
%9.2f	宽度9，精度2
%9.f	宽度9，精度0
*/
func demo6() {
	n := 12.34
	fmt.Printf("%f\n", n)    // 12.340000
	fmt.Printf("%9f\n", n)   // 12.340000
	fmt.Printf("%.2f\n", n)  // 12.34
	fmt.Printf("%9.2f\n", n) //     12.34
	fmt.Printf("%9.f\n", n)  //        12
}

// 获取输入
/*
Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。

Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。

本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
*/

func demo7() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
func demo8() {
	fmt.Println("请输入你的姓名:")
	var name string
	fmt.Scan(&name)
	fmt.Printf("你好:%s", name)
}

// fmt.Scanf
/*
简单一句话 在一行里面进行输入
Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
本函数返回成功扫描的数据个数和遇到的任何错误。
*/
func demo9() {
	var (
		name string
		age  int
	)
	fmt.Scan(&name, &age)
	fmt.Printf("扫描结果 name:%s age:%d\n", name, age)
}

// fmt.Scanln
/*
Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
*/
func demo10() {
	var (
		name string
		age  int
	)
	fmt.Scanln(&name, &age)
	fmt.Println(name, age)
}

// bufio.NewReader
/*
有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。
*/
func demo11() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	text, _ := reader.ReadString('\n') //读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

// Fscan系列
/*
这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。
*/

// Sscan系列
/*
这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
*/
func main() {
	demo11()
	//demo1()
	//demoErrorf()
	//demoSprint()
	//demoFprint()
	//demoPrint()
}
