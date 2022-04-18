package model

import (
	"gorm.io/gorm"
)

/*
	角色用户关系表
*/
func init() {
	M.AutoMigrate(&RoleUserMap{})
}

type RoleUserMap struct {
	gorm.Model
	AppKey string `gorm:"not null;`
	UserID uint
	RoleID uint
}

func (this *RoleUserMap) Create() error {
	return M.Create(this).Error
}
