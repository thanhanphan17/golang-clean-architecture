package entity

import (
	"go-clean-architecture/common"
)

const EntityName string = "User"

type User struct {
	common.Entity `json:",inline"`
	Name          string `json:"name" gorm:"column:name;"`
	Phone         string `json:"phone" gorm:"column:phone;"`
	Email         string `json:"email" gorm:"column:email;"`
	Role          string `json:"role" gorm:"column:role;default:user"`
	Status        string `json:"status" gorm:"column:status;default:unverified"`
	Point         int    `json:"point" gorm:"column:point;default:0"`
	OTP           int    `json:"-" gorm:"column:otp;default:999999"`
	Password      string `json:"-" gorm:"column:password;"`
	Salt          string `json:"-" gorm:"column:salt;"`
}

func (User) TableName() string {
	return "users"
}
