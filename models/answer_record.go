package models

import "github.com/jinzhu/gorm"

const (
	AnswerRecordUndefined = iota //未知结果，可能是跳过这题了
	AnswerRecordTrue             //结果正确
	AnswerRecordFalse            //结果错误
)

type AnswerRecord struct {
	gorm.Model
	UserId         uint   `gorm:"NOT NULL"`
	QuestionBankId uint   `gorm:"NOT NULL"`
	Result         int    `gorm:"type:TINYINT(2);"` //答题结果
	Answer         string `gorm:"type:TEXT;"`       //答题答案
}

func (r *AnswerRecord) TableName() string {
	return "answer_record"
}
