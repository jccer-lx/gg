package models

import "github.com/jinzhu/gorm"

//题分类
type QuestionCategory struct {
	gorm.Model
	Name     string `gorm:"type:VARCHAR(64);NOT NULL"`      //分类名称
	ParentId uint   `gorm:"type:INT(10) UNSIGNED;NOT NULL"` //父级分类id
}
