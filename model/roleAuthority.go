package model

import (
	"gorm.io/gorm"
)

/*
	应用角色权限表
*/
func init() {
	M.AutoMigrate(&RoleAuthority{})
}

type RoleAuthority struct {
	gorm.Model
	Code        string `gorm:"type:varchar(256);not null;comment:'编码'"`
	ActionValue uint32 `gorm:"not null;default:0;comment:'权限值'"`
	RoleID      uint   // Has Many
}
