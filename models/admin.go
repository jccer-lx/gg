package models

import (
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/impl"
)

type Admin struct {
	Id           uint32   `gorm:"type:INT(10) UNSIGNED;AUTO_INCREMENT;NOT NULL" json:"id"`
	Username     string   `gorm:"type:VARCHAR(20);NOT NULL" json:"username"`
	Nickname     string   `gorm:"type:VARCHAR(50);NOT NULL" json:"nickname"`
	Password     string   `gorm:"type:VARCHAR(32);NOT NULL" json:"password"`
	Salt         string   `gorm:"type:VARCHAR(30);NOT NULL" json:"salt"`
	Avatar       string   `gorm:"type:VARCHAR(255);NOT NULL" json:"avatar"`
	Email        string   `gorm:"type:VARCHAR(100);NOT NULL" json:"email"`
	Loginfailure uint8    `gorm:"type:TINYINT(1) UNSIGNED;NOT NULL" json:"loginfailure"`
	Logintime    int32    `gorm:"type:INT(10);" json:"logintime"`
	Loginip      string   `gorm:"type:VARCHAR(50);" json:"loginip"`
	Createtime   int32    `gorm:"type:INT(10);" json:"createtime"`
	Updatetime   int32    `gorm:"type:INT(10);" json:"updatetime"`
	Token        string   `gorm:"type:VARCHAR(59);NOT NULL" json:"token"`
	Status       string   `gorm:"type:VARCHAR(30);NOT NULL" json:"status"`
	Res          []*Admin `gorm:"-"`
}

func (t *Admin) TableName() string {
	return "at_admin"
}

func (t *Admin) GormFindOut() interface{} {
	return &t.Res
}

func (t *Admin) GetTableFields() ([]*impl.DataField, error) {
	return helper.GetStructFields(Admin{})
}
