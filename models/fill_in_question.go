package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//填空替换符号
const FillInReplace = "___"

//填空题
type FillInQuestion struct {
	BaseQuestion
	FillInNum  int      `gorm:"type:INT(10);"` //填空数
	AnswerList string   `gorm:"type:TEXT;"`    //答案
	Answers    []string `gorm:"-"`             //答案
}

func (q *FillInQuestion) TableName() string {
	return "question_fill_in"
}

func (q *FillInQuestion) GetType() int {
	return Choice
}

//填充answers内容保存到answer_list
func (q *FillInQuestion) BeforeCreate(scope *gorm.Scope) error {
	//如果answer_list有值，这个操作不进行
	if len(q.AnswerList) > 0 {
		return nil
	}
	//根据答案计算填空数量
	q.FillInNum = len(q.Answers)
	return q.Answers2AnswerList()
}

//填充answers
func (q *FillInQuestion) AfterFirst(scope *gorm.Scope) error {
	return q.AnswerList2Answers()
}

func (q *FillInQuestion) AfterFind(scope *gorm.Scope) error {
	return q.AnswerList2Answers()
}

func (q *FillInQuestion) Answers2AnswerList() error {
	oByte, err := json.Marshal(q.Answers)
	if err != nil {
		logrus.Error("Answers2AnswerList:", err)
		return err
	}
	logrus.Debug("Answers2AnswerList string(oByte):", string(oByte))
	q.AnswerList = string(oByte)
	return nil
}

func (q *FillInQuestion) AnswerList2Answers() error {
	var os []string
	err := json.Unmarshal([]byte(q.AnswerList), &os)
	if err != nil {
		return err
	}
	q.Answers = os
	return nil
}
