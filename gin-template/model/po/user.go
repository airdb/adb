package po

import (
	"github.com/jinzhu/gorm"
)

const (
	DBName = "dev_db"
)

type User struct {
	gorm.Model

	Nickname   string `gorm:"type:varchar(64)"`
	Headimgurl string `gorm:"type:varchar(128)"`
	Token      string `gorm:"type:varchar(128)"`
}

func ListProvider() (secret []*User) {
	return
}

func List(voUser string) *User {
	var user User

	return &user
}
