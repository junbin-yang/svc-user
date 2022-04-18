package model

import (
	"gorm.io/gorm"
)

/*
	菜单表，权限模板
*/
func init() {
	M.AutoMigrate(&Menu{})
}

type Menu struct {
	gorm.Model
	PID      uint   `gorm:"not null;comment:'上级Id'"`
	Code     string `gorm:"<-:create;index;type:varchar(256);not null;default:'';comment:'编码'"`
	Name     string `gorm:"type:varchar(256);not null;default:'';comment:'名称'"`
	Uri      string `gorm:"type:varchar(256);not null;default:'';comment:'路由'"`
	Remark   string `gorm:"size:4096;not null;default:'';comment:'备注'"`
	Sequence uint32 `gorm:"comment:'排序'"`
	Status   bool   `gorm:"default:true"`
	Actions  string `gorm:"type:mediumtext;not null;default:'';comment:'操作'"`
	AppKey   string `gorm:"not null;`
}

func (this *Menu) Create() error {
	return M.Create(this).Error
}
