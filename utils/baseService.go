package service

import (
	"errors"
	"gin-vue-admin/constant"
	"gin-vue-admin/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type baseService struct {
}

func (baseService) getDb() *gorm.DB {
	return global.GVA_DB
}
func (baseService) GetLog() *zap.Logger {
	return global.GVA_LOG
}
func (baseService) DbNotFond(err error) error {
	if err == gorm.ErrRecordNotFound {
		err = errors.New(constant.ErrNotFond)
	} else {
		err = errors.New(constant.ErrService)
	}
	return err
}
