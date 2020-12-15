package user

import "lean-go/3.gin/4.struct_demo/web/repository"

type ServiceUser struct {
}

func NewServiceUser() *ServiceUser {
	return &ServiceUser{}
}
func (ServiceUser) Login(username string, pwd string) (data interface{}, err error) {
	data = make(map[string]interface{})
	// 数据层操作
	user, err := repository.NewUserRep().GetByUsername(username)
	if err != nil {
		return data, err
	}
	return user, nil
}
