package models

import "github.com/jinzhu/gorm"

const (
	Choice         = iota //单选题
	MultipleChoice        //多选题
)

const (
	Level1 = iota //难度1-最简单
	Level2
	Level3
	Level4
	Level5
	Level6
	Level7
	Level8
	Level9
	Level10
)

//试题
type Questions interface {
	GetType() int //试题类型
}

type BaseQuestion struct {
	gorm.Model
	Stem       string  `gorm:"type:TEXT;"`                            //题干
	Score      float64 `gorm:"type:DECIMAL(10, 2) UNSIGNED;NOT NULL"` //分数
	Answer     string  `gorm:"type:TEXT;"`                            //参考答案
	Analysis   string  `gorm:"type:TEXT;"`                            //解析
	CategoryId uint    `gorm:"type:INT(10) UNSIGNED;NOT NULL"`        //类别
	Difficulty int     `gorm:"type:INT(10);NOT NULL"`                 //难度
}
