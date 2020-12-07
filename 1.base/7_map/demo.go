package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	demoBaseUse() //基本使用
}
func demoBaseUse() {
	// 声明 方式一
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 95
	scoreMap["yy"] = 99

	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	// 声明 方式二
	userInfo := map[string]string{
		"username": "yy",
		"pwd":      "123123",
	}
	fmt.Println(userInfo)

	// 判断某个键是否存在
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	// 遍历
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 删除键值对
	delete(scoreMap, "小明")
	fmt.Println(scoreMap)

	// 指定顺序遍历
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	sort.Strings(keys) // 对切片进行排序
	fmt.Println(keys)

	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Println(index, value)
		mapSlice[0] = make(map[string]string, 10)
		mapSlice[0]["name"] = "小王子"
		mapSlice[0]["pwd"] = "123123"
		mapSlice[0]["addr"] = "成都"
		for index, value := range mapSlice {
			fmt.Println(index, value)
		}
	}
	// 值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}
