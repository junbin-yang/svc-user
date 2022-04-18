package model

import (
	"gorm.io/gorm"
)

/*
	用户信息表
*/
func init() {
	M.AutoMigrate(&Info{})
}

type Info struct {
	gorm.Model
	RealName string `gorm:"type:varchar(256);not null;default:'';comment:'姓名'"`
	Age      uint   `gorm:"type:tinyint(1);not null;default:0"`
	Sex      uint   `gorm:"type:tinyint(1);not null;default:0;comment:'男性1，女性2，未知0'"`
	Nation   string `gorm:"not null;default:'';comment:'民族'"`
	Address  string `gorm:"size:2048;not null;default:''"`
	IdCard   string `gorm:"not null;index;default:'';comment:'身份证号码'"`
	UserID   uint   // Has One
}
