package main

import "fmt"

//gin框架路由拆分与注册
/*
/*
当项目的规模增大后就不太适合继续在项目的main.go文件中去实现路由注册相关逻辑了，我们会倾向于把路由部分的代码都拆分出来，形成一个单独的文件或包：
我们在routers.go文件中定义并注册路由信息：
*/
func main() {
	r := SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed,err:%v\n", err)
	}
}
