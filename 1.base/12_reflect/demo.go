package main

import (
	"fmt"
	"reflect"
)

func p(x interface{}) {
	fmt.Println(x)
}

// 反射
// 反射介绍
/*
反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，
变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，
并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有修改它们。
Go程序在运行期使用reflect包访问程序的反射信息。
*/

// -------------------------------------------------------------------------------------------
// reflect包
/*
在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，
并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取.
*/

// TypeOf
/*
在Go语言中，使用 reflect.TypeOf()函数可以获得任意值的类型对象(reflect.Type),
程序通过类型对象可以访问任意值的类型信息
*/
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v value:%v\n", v, x)

}
func demo1() {
	a := 3.14
	reflectType(a) //type:float64 value:3.14
	var b int64 = 100
	reflectType(b) //type:int64 value:100
}

// -------------------------------------------------------------------------------------------
// type name 和 type kind
/*
在反射中关于类型还分为两种  类型（Type）和种类（Kind）。因为在Go语言中我们可以使用
type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分
指针、结构体等大品种的类型时，就会用到种类（Kind）。
*/
type myInt int64

func printType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t, t.Kind())
}
func demo2() {
	var a *float32 // 指针
	var b myInt    //自定义类型
	var c rune     //类型别名
	printType(a)   //type:*float32 kind:ptr
	printType(b)   //type:main.myInt kind:int64
	printType(c)   // type:int32 kind:int32
}

// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的 .Name()都是返回空
/*
在 reflect包中定义的Kind类型如下：
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
*/

// -----------------------------------------------------------------------------------------------
// ValueOf
/*
reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
reflect.Value类型提供的获取原始值的方法如下：
方法                         说明
Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
Int() int64					将值以 int 类型返回，所有有符号整型均可以此方式返回
Uint() uint64				将值以 uint 类型返回，所有无符号整型均可以此方式返回
Float() float64				将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
Bool() bool					将值以 bool 类型返回
Bytes() []bytes				将值以字节数组 []bytes 类型返回
String() string				将值以字符串类型返回
*/

// 通过反射获取值
func printValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", v.Float())
	}
}
func demo3() {
	var a float32 = 3.14
	var b int64 = 100
	printValue(a)
	printValue(b)
	c := reflect.ValueOf(10)
	fmt.Printf("type c:%T\n", c)
}

// ------------------------------------------------------------------------------------------------
// 通过反射设置变量的值
/*
想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。
而反射中使用专有的Elem()方法来获取指针对应的值。
*/
func setValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本， reflect 包会引发 panic
	}
}
func setValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func demo4() {
	var a int64 = 100
	//setValue(&a) //注意这里是传指针
	setValue2(&a)
	fmt.Println(a)
}

// -----------------------------------------------------------------------------------------------------------------
// isNil() 和 isValid()
/*
isNil() 常被用于指针是否为空
isValid() 常被用于判定返回值是否有效
*/
func demo5() {
	var a *int
	fmt.Println("var a 8int IsNil:", reflect.ValueOf(a).IsNil()) //true
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())  //false
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找abc字段
	fmt.Println("不存在的体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid()) //false
	fmt.Println("不存在的方法:", reflect.ValueOf(b).MethodByName("abc").IsValid()) //false
	// map
	c := map[string]int{}
	fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("yy")).IsValid())
}

// --------------------------------------------------------------------------------------------------
// 结构体反射
// 当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字段名去获取指定的字段信息。
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func demo6() {
	yy := student{"yy", 90}
	t := reflect.TypeOf(yy)
	fmt.Println(t.Name(), t.Kind()) //student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", filed.Name, filed.Index, filed.Type, filed.Tag.Get("json"))
	}
}

// 接下来编写一个函数 来遍历打印s包含的方法

func (s student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}
func (s student) Sleep() string {
	msg := " 好好睡觉，快快长大"
	fmt.Println(msg)
	return msg
}
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		var args []reflect.Value
		v.Method(i).Call(args)
	}
}

/*
// 反射是把双刃剑
反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。
1. 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
2. 大量使用反射的代码通常难以理解。
3. 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级
*/
func main() {
	var s student
	printMethod(s)
}
