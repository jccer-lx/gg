package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const (
	StringOption = "string"
	ImageOption  = "img"
)

type ChoiceOption struct {
	OptionType string `json:"option_type"`
	Item       string `json:"item"`
}

//单选题
type ChoiceQuestion struct {
	BaseQuestion
	OptionList  string          `gorm:"type:TEXT;"`    //选项
	Options     []*ChoiceOption `gorm:"-"`             //选项
	AnswerIndex uint            `gorm:"type:INT(10);"` //参考答案对应索引
}

func (q *ChoiceQuestion) TableName() string {
	return "question_choice"
}

func (q *ChoiceQuestion) GetType() int {
	return Choice
}

//填充options内容保存到option_list
func (q *ChoiceQuestion) BeforeCreate(scope *gorm.Scope) error {
	//如果option_list有值，这个操作不进行
	if len(q.OptionList) > 0 {
		return nil
	}
	return q.Options2OptionList()
}

//填充options
func (q *ChoiceQuestion) AfterFirst(scope *gorm.Scope) error {
	return q.OptionList2Options()
}

func (q *ChoiceQuestion) AfterFind(scope *gorm.Scope) error {
	return q.OptionList2Options()
}

func (q *ChoiceQuestion) Options2OptionList() error {
	oByte, err := json.Marshal(q.Options)
	if err != nil {
		logrus.Error("Options2OptionList:", err)
		return err
	}
	logrus.Debug("Options2OptionList string(oByte):", string(oByte))
	q.OptionList = string(oByte)
	return nil
}

func (q *ChoiceQuestion) OptionList2Options() error {
	var os []*ChoiceOption
	err := json.Unmarshal([]byte(q.OptionList), &os)
	if err != nil {
		return err
	}
	q.Options = os
	return nil
}
