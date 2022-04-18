package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

/*
	用户基础表
*/
func init() {
	M.AutoMigrate(&User{}) //仅会创建表，缺少列和索引，并不会改变现有列的类型或删除未使用的列。
	M.Migrator().AlterColumn(&User{}, "LoginedAt")
	root := User{
		Account:  "root",
		Password: "e10adc3949ba59abbe56e057f20f883e",
		NickName: "超级管理员",
		Admin:    true,
	}
	M.Where(User{Account: "root"}).FirstOrCreate(&root)
}

type User struct {
	gorm.Model
	Account   string     `gorm:"<-:create;index;unique;type:varchar(256);not null;default:'';comment:'用户账号'"` //允许读和创建(禁止更新字段)
	Password  string     `gorm:"type:varchar(256);not null" json:"-"`                                         //隐藏json字段
	NickName  string     `gorm:"type:varchar(256);not null;default:''"`
	Mobile    string     `gorm:"type:varchar(64);unique;not null;default:''"`
	Email     string     `gorm:"type:varchar(64);unique;not null;default:''"`
	Status    bool       `gorm:"default:true"` //像 0、''、false 等零值，不会将这些字段定义的默认值保存到数据库。使用指针类型来避免这个问题
	Admin     bool       `gorm:"default:false;comment:'是否是管理员'"`
	LoginedAt *time.Time `gorm:"default:'0000-00-00 00:00:00.000'"`
	Info      Info       // Has One
	Other     string     `gorm:"-"` //读写会忽略该字段
}

func (this *User) Create() error {
	return M.Create(this).Error
}

func (this *User) Update(data map[string]interface{}) (int64, error) {
	result := M.Model(this).Updates(data)
	return result.RowsAffected, result.Error
}

func (this *User) FirstByAccount(account, password string) (int64, error) {
	result := M.Where("account = ? AND password = ?", account, password).Preload(clause.Associations).First(this)
	return result.RowsAffected, result.Error
}

func (this *User) FirstByEmail(email, password string) (int64, error) {
	result := M.Where("email = ? AND password = ?", email, password).Preload(clause.Associations).First(this)
	return result.RowsAffected, result.Error
}

func (this *User) FirstByMobile(mobile, password string) (int64, error) {
	result := M.Where("mobile = ? AND password = ?", mobile, password).Preload(clause.Associations).First(this)
	return result.RowsAffected, result.Error
}

func (this *User) FirstByUserId(userid uint64) (int64, error) {
	result := M.Where("id = ?", userid).Preload(clause.Associations).First(this)
	return result.RowsAffected, result.Error
}

func (this *User) UpdateInfo(id uint64) error {
	// 返回任何错误都会回滚事务
	return M.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(User{}).Where("id = ?", id).Updates(map[string]interface{}{
			"nick_name": this.NickName,
			"mobile":    this.Mobile,
			"email":     this.Email,
			"status":    this.Status,
		}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Model(Info{}).Where("user_id = ?", id).Updates(map[string]interface{}{
			"real_name": this.Info.RealName,
			"age":       this.Info.Age,
			"sex":       this.Info.Sex,
			"nation":    this.Info.Nation,
			"address":   this.Info.Address,
			"id_card":   this.Info.IdCard,
		}).Error; err != nil {
			return err
		}

		return nil
	})
}
