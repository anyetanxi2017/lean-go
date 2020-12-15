package repository

import (
	"github.com/jinzhu/gorm"
	"lean-go/3.gin/4.struct_demo/web/common/constants"
	"lean-go/3.gin/4.struct_demo/web/models"
)

type UserRep struct {
	baseRep
}

func NewUserRep() *UserRep {
	return &UserRep{}
}
func (r *UserRep) GetByUsername(username string) (*models.User, error) {
	db := r.getDb()
	user := &models.User{}
	if err := db.Where("username = ? and status = ?", username, constants.UserNormal).First(user).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			// 日志记录下
		}
		return nil, err
	}
	return user, nil
}
