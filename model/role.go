package model

import (
	"errors"
	"gorm.io/gorm"
)

/*
	应用角色表
*/
func init() {
	M.AutoMigrate(&Role{})
}

type Role struct {
	gorm.Model
	Name           string `gorm:"type:varchar(256);not null;default:'';comment:'名称'"`
	Remark         string `gorm:"size:2048;not null;default:'';comment:'备注'"`
	Sequence       uint32 `gorm:"comment:'排序'"`
	Status         *bool  `gorm:"default:true"`
	AppKey         string `gorm:"not null;`
	RoleAuthoritys []RoleAuthority
	UserCount      int64       `gorm:"-"`
	Other          interface{} `gorm:"-"`
	AssUserIds     []uint      `gorm:"-"`
}

func (this *Role) Create() error {
	return M.Create(this).Error
}

func (this *Role) Update(roleId uint) error {
	// 返回任何错误都会回滚事务
	return M.Transaction(func(tx *gorm.DB) error {
		var total int64
		tx.Model(&Role{}).Where("id = ? AND app_key = ?", roleId, this.AppKey).Count(&total)
		if total <= 0 {
			return errors.New("未找到角色")
		}

		if err := tx.Model(Role{}).Where("id = ? AND app_key = ?", roleId, this.AppKey).Updates(map[string]interface{}{
			"name":     this.Name,
			"remark":   this.Remark,
			"sequence": this.Sequence,
			"status":   this.Status,
		}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Where("role_id = ?", roleId).Delete(RoleAuthority{}).Error; err != nil {
			return err
		}

		for _, authoritys := range this.RoleAuthoritys {
			authoritys.RoleID = roleId
			if err := tx.Create(&authoritys).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (this *Role) Delete(roleId uint, appKey string) error {
	return M.Unscoped().Select("RoleAuthoritys").Delete(&Role{Model: gorm.Model{ID: roleId}, AppKey: appKey}).Error
}
