package main

import (
	"fmt"
	"strconv"
)

//Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换。
func demo60() {

	// Atoi()  string to int
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%d \n", i1, i1)
	}

	// Itoa() int to string
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%s\n", s2, s2)
}

// Parse系列函数
func demo61() {
	// ParseInt
	// base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	// bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	n1, _ := strconv.ParseInt("20", 10, 64)
	fmt.Printf("type:%T value:%d\n", n1, n1)

	// ParseUnit()
	// ParseUint类似ParseInt但不接受正负号，用于无符号整型。
	n2, _ := strconv.ParseUint("1", 10, 64)
	fmt.Printf("type:%T value:%d\n", n2, n2)

	// ParseFloat()
	// bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
	f1, _ := strconv.ParseFloat("3", 64)
	fmt.Printf("type:%T value:%f\n", f1, f1)
}

// Format系统函数
func demo62() {
	// FormatBool()
	s1 := strconv.FormatBool(true)
	fmt.Printf("type:%T value:%s\n", s1, s1) //type:string value:true

	// FormatInt()
	// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。
	s2 := strconv.FormatInt(5, 23)
	fmt.Printf("type:%T value:%s\n", s2, s2) //type:string value:true

	// FormatUint()
	s3 := strconv.FormatUint(10, 10)
	fmt.Printf("type:%T value:%s\n", s3, s3)

	// FormatFloat()
	// fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
	// prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	s4 := strconv.FormatFloat(3.1415, 'f', 2, 64)
	fmt.Printf("type:%T value:%s\n", s4, s4)
}
func main() {
	demo62()
}
