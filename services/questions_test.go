package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
	"github.com/sirupsen/logrus"
	"math/rand"
	"reflect"
	"testing"
	"time"
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
	//打乱顺序
	randSlice(choiceQuestionList)

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

//打乱切片顺序
func randSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		fmt.Println("rv.Type().Kind() != reflect.Slice")
		return
	}

	length := rv.Len()
	if length < 2 {
		fmt.Println("length := rv.Len()")
		return
	}

	swap := reflect.Swapper(slice)
	rand.Seed(time.Now().Unix())
	for i := length - 1; i >= 0; i-- {
		j := rand.Intn(length)
		swap(i, j)
	}
	return
}
