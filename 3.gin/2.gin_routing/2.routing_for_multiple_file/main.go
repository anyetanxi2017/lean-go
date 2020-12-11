package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lean-go/3.gin/3.1.gin_routing/2.routing_for_multiple_file/routers"
)

func main() {
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startuup service failed,err:%v\n", err)
	}
}
