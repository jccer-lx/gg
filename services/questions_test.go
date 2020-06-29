package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestAddChoiceQuestionForBank(t *testing.T) {
	databases.InitMysqlDB()
	logrus.SetLevel(logrus.DebugLevel)
	err := AddChoiceQuestionForBank(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestAddAllChoiceQuestionForBank(t *testing.T) {
	databases.InitMysqlDB()
	logrus.SetLevel(logrus.DebugLevel)
	var choiceQuestionList []*models.ChoiceQuestion
	err := databases.NewDB().Where("id > 1").Find(&choiceQuestionList).Error
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for _, cq := range choiceQuestionList {
		err := AddChoiceQuestionForBank(cq.ID)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
	}
}

func TestGetQuestionByBankId(t *testing.T) {
	databases.InitMysqlDB()
	logrus.SetLevel(logrus.DebugLevel)
	questionBankModel, err := GetQuestionByBankId(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(questionBankModel.ID)
	fmt.Println(questionBankModel.QuestionType)
	fmt.Println(questionBankModel.Question.GetId())
	fmt.Println(questionBankModel.Question.(*models.ChoiceQuestion).Stem)
}
