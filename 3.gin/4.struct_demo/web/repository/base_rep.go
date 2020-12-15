package repository

import (
	"github.com/jinzhu/gorm"
	"lean-go/3.gin/4.struct_demo/basic/db"
)

type baseRep struct {
	apiTx *gorm.DB
}

func (r *baseRep) getDb() *gorm.DB {
	if r.apiTx != nil {
		return r.apiTx
	}
	return db.GetDbClient()
}
func (*baseRep) getMainDbSelect() *gorm.DB {
	return db.GetMainDbSelect()
}
func (*baseRep) getFansDb() *gorm.DB {
	return db.GetFansDb()
}
