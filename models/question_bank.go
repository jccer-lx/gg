package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/*每日题库*/
type QuestionBank struct {
	gorm.Model
	QuestionType int
	QuestionId   uint
}

func (q *QuestionBank) TableName() string {
	return "question_bank"
}

func (q *QuestionBank) AddQuestion(question Questions) error {
	if question.GetId() == 0 {
		return fmt.Errorf("question.GetId error")
	}
	q.QuestionId = question.GetId()
	q.QuestionType = question.GetType()
	return nil
}
