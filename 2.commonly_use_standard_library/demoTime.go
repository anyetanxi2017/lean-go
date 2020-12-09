package main

import (
	"fmt"
	"time"
)

// 时间类型
/*
time.Time类型表示时间。我们可以通过time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。
*/
func demo20() {
	now := time.Now()
	fmt.Println(now)         //2020-12-10 00:29:27.9871901 +0800 CST m=+0.003017501
	fmt.Println(now.Year())  //2020
	fmt.Println(now.Month()) // December
	fmt.Println(now.Day())   //10
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())           //秒
	fmt.Println(now.UnixNano() / 1e6) // 毫秒
	fmt.Println(now.UnixNano())       //纳秒
}

// 使用 time.Unix()函数可以将时间戳转为时间格式
/*
是从秒时间开始转的，如果是毫秒记得再除3
*/
func demo21() {
	unix := time.Now().UnixNano() / 1e9
	date := time.Unix(unix, 0)
	fmt.Println(date)

}

// 时间间隔
/*
time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
time.Duration表示一段时间间隔，可表示的最长时间段大约290年。
*/
func demo22() {
	fmt.Println(time.Nanosecond)  // 1ns
	fmt.Println(time.Millisecond) // 1ms
	fmt.Println(time.Second)      // 1s
	fmt.Println(time.Minute)      // 1m0s
	fmt.Println(time.Hour)        // 1h0m0s
}

// 时间操作
/*
我们在日常的编码过程中可能会遇到要求时间+时间间隔的需求，Go语言的时间对象有提供Add方法如下
*/
func demo23() {
	now := time.Now()
	// add
	fmt.Println(now.Add(time.Hour))  //当前时间加1小时后的时间
	fmt.Println(now.Add(-time.Hour)) //当前时间减1小时
	// sub
	fmt.Println(now.Sub(now.Add(-time.Hour))) //求两个时间之差 1h0m0s
	// equal
	fmt.Println(now.Equal(now.Add(time.Hour))) //false
	// before after
	parse, _ := time.Parse("2006-04-02", "2020-12-11")
	fmt.Println(now.After(parse))
}

// 定时器
/*
使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）。
*/
func demo24() {
	var wait = time.Second
	tick := time.Tick(wait)
	for range tick { //每秒执行一次
		fmt.Println(time.Now())
	}
}
func demo25() {
	wait := time.Second
	count := 10
	for {
		count++
		if count%20 == 0 {
			wait = time.Second * 10
			fmt.Println("哈哈，我进来啦")
			someThing()
			wait = time.Second
		}
		fmt.Println(time.Now())
		time.Sleep(wait)
	}
}
func someThing() {
	time.Sleep(time.Second * 5)
	fmt.Println("做完了")
}

// 时间格式化
/*
时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧时间类型有一个自带的方法Format进行格式化，
需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧

补充：如果想格式化为12小时方式，需指定PM。
*/
func demo26() {
	now := time.Now()
	// 24小时
	fmt.Println(now.Format("2006-01-02 15:04:05.000 pm Mon Jan"))
	// 12小时
	fmt.Println(now.Format("2006-01-02 3:04:02.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/04/02 15:04"))
}

// 解析字符串格式的时间
func demo27() {
	begin := time.Now()
	now := time.Now()
	fmt.Println(now)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/12/10 01:46:28", location)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(time.Now().Sub(begin).Nanoseconds() / 10e5) // 用时 微秒

}
func main() {
	demo27()
}
