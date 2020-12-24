package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Class struct {
	gorm.Model
	Name     string
	Students []Student
}
type Student struct {
	gorm.Model
	ClassID  uint
	Name     string
	IDCard   IDCard
	Teachers []Teacher `gorm:"many2many:student_teachers"`
}
type IDCard struct {
	StudentID uint
	Num       int
}
type Teacher struct {
	gorm.Model
	Name     string
	Students []Student `gorm:"many2many:student_teachers"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/bhmc_common?charset=utf8mb4&parseTime=True&loc=Local&timeout=20s&readTimeout=20s&writeTimeout=20s")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var s Student
		db.Preload("Teachers").Preload("IDCard").First(&s, "id = ?", id)
		c.JSON(http.StatusOK, gin.H{"s": s})
	})
	defer db.Close()
	r.Run(":8888")
}
