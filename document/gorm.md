## 连接数据库
```
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func GetDb() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/qmplus?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
```
