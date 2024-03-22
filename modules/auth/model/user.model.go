package model

import "github.com/LuanTruongPTIT/tutor-be/common"

type Auth struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
}

func (Auth) TableName() string {
	return "auth"
}
