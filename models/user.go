package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Openid             string `gorm:"type:VARCHAR(64);NOT NULL;UNIQUE"`
	LastQuestionBankId uint   //参与题库最后一条数据ID
}

func (u *User) TableName() string {
	return "user"
}
