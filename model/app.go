package model

import (
	"gorm.io/gorm"
)

/*
	应用注册表
*/
func init() {
	M.AutoMigrate(&App{})
}

type App struct {
	gorm.Model
	Name      string `gorm:"type:varchar(256);not null;default:'';comment:'APP名称'"`
	Remark    string `gorm:"size:2048;not null;default:'';comment:'备注'"`
	AppKey    string `gorm:"<-:create;unique;index;not null"`
	AppSecret string `gorm:"not null"`
	Status    bool   `gorm:"default:true"`
}

func (this *App) Create() error {
	return M.Create(this).Error
}

func (this *App) FirstById(id uint64) (int64, error) {
	result := M.Where("id = ?", id).First(this)
	return result.RowsAffected, result.Error
}
