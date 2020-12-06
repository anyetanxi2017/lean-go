package main

import (
	"fmt"
	"strings"
)

func main() {
	//demoIf()
	//demoFor() // 无限循环 直接 for {} 就好
	//demoRange()  // range  还可以遍历map的 k v值
	//demoSwitchCase() // case 可以是表达式
	demoGoto() //goto语句通过标签进行代码间的无条件跳转

}
func demoIf() {
	// if条件判断特殊写法

	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}
}

func demoFor() {
	// for循环结构
	// Go语言中的所有循环类型均可使用for关键字来完成
	// for 循环可以通过 break goto return panic 语句强制退出循环
	// for 循环可以通过 continue 跳过本次循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	for i < 10 {
		fmt.Println(i)
		i++
	}
	//无限循环
	//for {
	//	fmt.Println(time.Second)
	//}
	fmt.Println("----------")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func demoRange() {

	// 遍历数组1
	split := strings.Split("hello", "")
	for i, str := range split {
		fmt.Println(i, str)
	}
	// 遍历数组2 忽略index
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// 遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}

}

func demoSwitchCase() {
	switch n := 3; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}

	//  分支可以有多个值
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	// 分支还可以使用表达式
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	// fallthrough 语法可以执行满足条件的case的下一个case，是为了兼容c语言中的case设计的
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

}

func demoGoto() {
	//
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 7 {
				goto breakTag
			}
			fmt.Println("继续学习")
		}
		if breakFlag {
			break
		}
	}
	fmt.Println("别走啊")
breakTag:
	fmt.Printf("来了")
}
