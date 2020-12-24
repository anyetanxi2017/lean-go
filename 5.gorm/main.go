package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/bhmc_common?charset=utf8mb4&parseTime=True&loc=Local&timeout=20s&readTimeout=20s&writeTimeout=20s")
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
	defer db.Close()
}
