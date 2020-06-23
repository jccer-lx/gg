package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//多选题
type MultipleChoiceQuestion struct {
	BaseQuestion
	OptionList  string   `gorm:"type:TEXT;"` //选项
	Options     []string `gorm:"-"`          //选项
	AnswerIndex string   `gorm:"type:TEXT;"` //参考答案对应索引
}

func (q *MultipleChoiceQuestion) TableName() string {
	return "multiple_questions_choice"
}

func (q *MultipleChoiceQuestion) GetType() int {
	return Choice
}

//填充options内容保存到option_list
func (q *MultipleChoiceQuestion) BeforeCreate(scope *gorm.Scope) error {
	//如果option_list有值，这个操作不进行
	if len(q.OptionList) > 0 {
		return nil
	}
	return q.Options2OptionList()
}

//填充options
func (q *MultipleChoiceQuestion) AfterFirst(scope *gorm.Scope) error {
	return q.OptionList2Options()
}

func (q *MultipleChoiceQuestion) AfterFind(scope *gorm.Scope) error {
	return q.OptionList2Options()
}

func (q *MultipleChoiceQuestion) Options2OptionList() error {
	oByte, err := json.Marshal(q.Options)
	if err != nil {
		logrus.Error("Options2OptionList:", err)
		return err
	}
	logrus.Debug("Options2OptionList string(oByte):", string(oByte))
	q.OptionList = string(oByte)
	return nil
}

func (q *MultipleChoiceQuestion) OptionList2Options() error {
	var os []string
	err := json.Unmarshal([]byte(q.OptionList), &os)
	if err != nil {
		return err
	}
	q.Options = os
	return nil
}
