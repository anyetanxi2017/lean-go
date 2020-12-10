package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lean-go/3_gorm/modules"
)

var db *gorm.DB

// 查询
func demoSelect() {
	var (
		u4, u1, u2, u3 modules.User
	)
	db.Model(modules.User{}).First(&u1)
	fmt.Println(u1)
	db.Model(modules.User{}).Where("username=?", "user1").First(&u2)
	fmt.Println(u2)
	db.Model(modules.User{Username: "user2"}).First(&u3)
	fmt.Println(u3)
	db.Model(modules.User{ID: 13}).First(&u4)
	fmt.Println(u4)
}

// 更新
func demoUpdate() {
	// 方式一
	db.Model(&modules.User{}).Where("id=?", 13).Update("nickname", "yy2")
	// 方式二
	update := db.Model(&modules.User{ID: 13}).Update("nickname", "yy3")
	fmt.Println("更新条数:", update.RowsAffected)
	var u1 modules.User
	db.Model(modules.User{ID: 13}).First(&u1)
	fmt.Println(u1)
}
func main() {
	demoUpdate()
}

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/rita_stock?charset=utf8mb4&parseTime=True&loc=Local"
	dbs, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = dbs
}
