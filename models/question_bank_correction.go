package models

import "github.com/jinzhu/gorm"

type QuestionBankCorrection struct {
	gorm.Model
	QuestionBankId uint
	UserId         uint
}

func (q *QuestionBankCorrection) TableName() string {
	return "question_bank_correction"
}
