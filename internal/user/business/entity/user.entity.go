package entity

import (
	"go-clean-architecture/common/base"
)

const EntityName = "User"

type User struct {
	base.BaseEntity `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	Email           string `json:"email" gorm:"column:email;"`
	Role            string `json:"role" gorm:"column:role;defaul:user"`
	Addr            string `json:"addr" gorm:"column:addr;"`
	Status          string `json:"status" gorm:"column:status;defaul:inactive"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	OTP             string `json:"otp" gorm:"column:otp;default:999999"`
}

func (User) TableName() string {
	return "users"
}
