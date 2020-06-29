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
