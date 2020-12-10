package modules

import "time"

type User struct {
	ID            int
	Pid           int
	Username      string
	Nickname      string
	Password      string
	Icon          string
	usertype      int `gorm:"column:type"`
	state         int
	Gender        int
	Phone         string
	Email         string
	JoinIp        string
	LastLoginTime time.Time
	LastLoginIp   string
	CreateTime    time.Time
}

func (u User) TableName() string {
	return "us_user"
}
