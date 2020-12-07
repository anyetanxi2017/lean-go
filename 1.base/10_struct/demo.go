package main

import (
	"encoding/json"
	"fmt"
)

// Go语言中没有类的概念，也不支持类的继承等面向对象的概念。
// Go语言中通过结构体的内嵌再配合接口纟面向对象具有更高的扩展性和灵活性

// Go语言中可以使用 type 关键字来定义自定义类型

// type示例
// 将MyInt定义为int类型
// 通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性
type MyInt int

// 类型别名和类型定义的区别
//类型定义
type NewInt int

// 类型别名
type MeInt = int

// ------------------------------------------------------------

// 结构体
// Go语言中通过 struct 来实现面向对象
type Person struct {
	name string
	city string
	age  int8
}

//同样类型的字段也可以写在一行
type person1 struct {
	name, city string
	age        int8
}

// 结构体实例化
// 只有当结构体实例化时，都会真正地分配内存
func demoStruct() {
	var p1 Person
	p1.name = "yy"
	p1.city = "上海"
	p1.age = 22
	fmt.Printf("p1=%v\n", p1)  //p1={yy 上海 22}
	fmt.Printf("p1=%#v\n", p1) //p1=main.Person{name:"yy", city:"上海", age:22}

	//使用键值对初始化
	p2 := Person{name: "yy", city: " 上海", age: 22}
	fmt.Printf("p2=%#v\n", p2)
	// 直接赋值
	p3 := Person{"yy", "上海", 22}
	fmt.Printf("p3=%#v\n", p3)

}

// 匿名结构体
// 在定义一些临时数据结构等场景下还可以使用匿名结构体
func demoStruct2() {
	var user struct {
		name string
		age  int
	}
	user.name = "yy"
	user.age = 20
	fmt.Printf("%#v", user)
}

// 构造函数
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的结构体的指针类型
func newPerson(name, city string, age int8) *Person {
	return &Person{name, city, age}
}

// ------------------------------------------------------------------------------------------------------

// 什么时候应该珍指针类型接收者?
// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者

// 方法和接收者
// Dream Person 梦想的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言", p.name)
}

// 指针类型的接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// 值类型接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

// ------------------------------------------------------------------------------------------------------

// 任意类型添加方法
// 在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
// 为MyInt 添加一个 SayHello 方法
func (m MyInt) SayHello() {
	fmt.Println("Hello,我是一个Int。")
}

// ------------------------------------------------------------------------------------------------------
// 结构体的匿名字段
// 这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，
// 因此一个结构体中同种类型的匿名字段只能有一个。
type Person2 struct {
	string
	int
}

// ------------------------------------------------------------------------------------------------------
// 嵌套结构体

type Address struct {
	Province, City string
}
type User struct {
	Name, Gender string
	Address      Address
}

// 嵌套匿名字段

type User2 struct {
	Name, Gender string
	Address
}

// ------------------------------------------------------------------------------------------------------
// 结构体的“继承”
// Go语言中使用结构体也可以实现继承

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d Dog) wang() {
	fmt.Printf("%s 会 汪汪汪~\n", d.name)
}

// ------------------------------------------------------------------------------------------------------
// 结构体与JSON
type Student struct {
	Id           int
	Gender, Name string
}

// 班级
type Class struct {
	Title    string
	Students []*Student
}

func demoJSON() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			Id:     i,
		}
		c.Students = append(c.Students, stu)
	}
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("%s\n", data)

	// JSON 反序列化
	str := `{"Title":"101","Students":[{"Id":0,"Gender":"男","Name":"stu00"},{"Id":1,"Gender":"男","Name":"stu01"},{"Id":2,"Gender":"男","Name":"stu02"},{"Id":3,"Gender":"男","Name":"stu03"},{"Id":4,"Gender":"男","Name":"stu04"},{"Id":5,"Gender":"男","Name":"stu05"},{"Id":6,"Gender":"男","Name":"stu06"},{"Id":7,"Gender":"男","Name":"stu07"},{"Id":8,"Gender":"男","Name":"stu08"},{"Id":9,"Gender":"男","Name":"stu09"}]}   `
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

// ------------------------------------------------------------------------------------------------------
// 结构体标签 Tag
// Tag是结构体的元信息，可以在运行的时候通过反射机制读取出来 格式如下
// `key1:"value1" key2:"value2"`
type Wallet struct {
	ID         int `json:"id"` //通过指定tag现实json序列化该字段时的key
	Username   string
	createTime string // 私有不能被json包访问
}

func demoWallet() {
	w := Wallet{ID: 1, Username: "yy", createTime: "2020-12-8"}
	data, err := json.Marshal(w)
	if err != nil {
		fmt.Println("json marshal failed!")
	}
	fmt.Printf("%s\n",data)

}

func main() {
	demoWallet()
	// 类型别名和类型定义的区别 示例
	//var a NewInt
	//var b MeInt
	//fmt.Printf("type of a %T\n", a) //main.NewInt
	//fmt.Printf("type of b %T\n", b) //int

	//结构体
	//demoStruct()
	//demoStruct2()

	// 构造函数 示例
	//yy := newPerson("yy", "上海", 22)
	//yy.age = 26
	//fmt.Println(*yy)

	// 方法 示例
	//p1 := Person{"hx", "成都", 30}
	//p1.Dream()

	// 指针类型接收者 示例
	//fmt.Println(p1) //30
	//p1.SetAge(35)
	//fmt.Println(p1) //35

	// 值类型接收者 示例
	//fmt.Println(p1) //30
	//p1.SetAge2(35)  //不会改变
	//fmt.Println(p1) //30

	// 任意类型添加方法 示例
	//var m MyInt
	//m.SayHello()
	//m = 100
	//fmt.Printf("%#v, %T\n", m, m)

	// 匿名字段 示例
	//p := Person2{"yy", 22}
	//fmt.Printf("%#v\n", p) //main.Person2{string:"yy", int:22}

	// 嵌套结构体 示例
	//user := User{Name: "yy", Gender: "女", Address: Address{Province: "上海", City: "上海"}}
	//fmt.Printf("user=%#v", user)

	// 继承 示例
	//dog := Dog{Feet: 4, Animal: &Animal{" 苹果"}}
	//dog.wang()
	//dog.move()
	//fmt.Println(dog)
	//demoJSON()
}
