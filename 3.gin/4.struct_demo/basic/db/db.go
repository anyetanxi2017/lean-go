package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lean-go/3.gin/4.struct_demo/basic/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	dbc    = make(map[string]*gorm.DB)
	models = []interface{}{}
)

func Init(config config.Mysql) {
	fmt.Println("配置数据库:", config.Addr)
	db := GetMysql(config.NameSpace, config.Addr)
	// 开启 Logger,以展示详细的日志
	db.LogMode(true)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetConnMaxLifetime(30 * time.Minute)
	if err := createTable(db); err != nil {
		panic(err)
	}
}
func GetMysql(nameSpace, addr string) *gorm.DB {
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(err)
	}
	dbc[nameSpace] = db
	return dbc[nameSpace]
}
func GetDbClient(nameSpace ...string) *gorm.DB {
	var s string
	switch nameLen := len(nameSpace); nameLen {
	case 0:
		s = "default"
	case 1:
		s = nameSpace[0]
	default:
		panic("nameSpace receive at most one parameter")
	}
	if _, ok := dbc[s]; ok {
		return dbc[s]
	}
	panic(fmt.Sprintf("the redis connect(%s) is not exist", s))
}

// 解决主从延迟问题，主库查询较为实时的数据（仅有查询权限）
func GetMainDbSelect() *gorm.DB {
	return dbc["maindbselect"]
}

// fans
func GetFansDb() *gorm.DB {
	return dbc["fans"]
}

func createTable(dbc *gorm.DB) error {
	for _, m := range models {
		if dbc.HasTable(m) {
			continue
		}
		// todo 这是在创建Table？
		if err := dbc.CreateTable(m).Error; err != nil {
			return err
		}
	}
	return nil
}
