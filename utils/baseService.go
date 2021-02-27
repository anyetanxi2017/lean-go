package service

import (
	"errors"
	"gin-vue-admin/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ErrNotFond = "记录不存在"
	ErrService = "服务器异常，请稍后再试"
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
		err = errors.New(ErrNotFond)
	} else {
		err = errors.New(ErrService)
	}
	return err
}
