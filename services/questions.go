package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
	"github.com/sirupsen/logrus"
)

/*
添加选择题到题库
@param uint id 选择题ID
@return
*/
func AddChoiceQuestionForBank(id uint) (err error) {
	//验证合法
	choiceQuestionModel := new(models.ChoiceQuestion)
	choiceQuestionModel.ID = id
	err = databases.NewDB().First(choiceQuestionModel).Error
	if err != nil {
		logrus.Debug("databases.NewDB().First(choiceQuestionModel)", err)
		return
	}
	//题库数据暂时不让重复
	questionBankModel := new(models.QuestionBank)
	bankErr := databases.NewDB().First(questionBankModel, map[string]interface{}{
		"question_type": models.Choice,
		"question_id":   id,
	}).Error
	if bankErr != nil && bankErr != gorm.ErrRecordNotFound {
		err = bankErr
		logrus.Debug("databases.NewDB().First(questionBankModel)", err)
		return
	}
	if questionBankModel.ID > 0 {
		err = fmt.Errorf("此题目已经在题库中")
		return
	}
	//保存
	err = questionBankModel.AddQuestion(choiceQuestionModel)
	if err != nil {
		logrus.Debug("questionBankModel.AddQuestion", err)
		return
	}
	err = databases.NewDB().Save(questionBankModel).Error
	return
}
