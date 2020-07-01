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

/*
根据ID返回题库内容
@param uint id 选择题ID
@return
*/
func GetQuestionByBankId(id uint) (questionBankModel *models.QuestionBank, err error) {
	questionBankModel = new(models.QuestionBank)
	err = databases.NewDB().First(questionBankModel, map[string]interface{}{
		"id": id,
	}).Error
	if err != nil {
		logrus.Debug("databases.NewDB().First(questionBankModel)", err)
		return
	}
	err = FindQuestionByType(questionBankModel)
	return
}

/*
根据ID返回下一题题库内容，未添加条件
@param uint id 选择题ID
@return
*/
func GetNextQuestionByBankId(id uint) (questionBankModel *models.QuestionBank, err error) {
	questionBankModel = new(models.QuestionBank)
	err = databases.NewDB().Where("id > ?", id).First(questionBankModel).Error
	if err != nil {
		logrus.Debug("databases.NewDB().First(questionBankModel)", err)
		return
	}
	err = FindQuestionByType(questionBankModel)
	return
}

func FindQuestionByType(questionBankModel *models.QuestionBank) error {
	//查询对应题目内容
	switch questionBankModel.QuestionType {
	case models.Choice:
		questionBankModel.Question = new(models.ChoiceQuestion)
	case models.Judgment:
		questionBankModel.Question = new(models.JudgmentQuestion)
	case models.MultipleChoice:
		questionBankModel.Question = new(models.MultipleChoiceQuestion)
	default:
		return fmt.Errorf("暂时不支持的类型")
	}
	//查询
	err := databases.NewDB().First(questionBankModel.Question, map[string]interface{}{
		"id": questionBankModel.QuestionId,
	}).Error
	if err != nil {
		logrus.Debug("questionBankModel 类型查询：", err)
		return err
	}
	if questionBankModel.Question.GetId() == 0 {
		return fmt.Errorf("未查询到对应题目数据")
	}
	return nil
}

/*
判断答案是否正确
@param uint id 题库ID
@param uint userId 用户ID
@param string answer 答案
@return bool res 回答是否正确
@return *models.QuestionBank questionBankModel 具体题目信息
*/
func CheckQuestionAnswer(id, userId uint, answer string) (res bool, questionBankModel *models.QuestionBank, err error) {
	questionBankModel, err = GetQuestionByBankId(id)
	if err != nil {
		return
	}
	//答案判断
	res = questionBankModel.Question.GetAnswer() != answer
	//记录
	if res {
		addAnswerRecord(id, userId, answer, models.AnswerRecordTrue)
	} else {
		addAnswerRecord(id, userId, answer, models.AnswerRecordFalse)
	}
	return
}

/*
题目纠错，每个用户只能对题目纠错一次
@param uint id 题库ID
@param uint userId 用户ID
*/
func CorrectionQuestion(id, userId uint) (err error) {
	logrus.Debug("CorrectionQuestion params:", id, userId)
	//用户存在？
	userModel := new(models.User)
	userModel.ID = userId
	err = databases.NewDB().First(userModel).Error
	if err != nil {
		err = fmt.Errorf("用户ID错误")
		return
	}
	//题库存在？
	questionBankModel := new(models.QuestionBank)
	questionBankModel.ID = id
	err = databases.NewDB().First(questionBankModel).Error
	if err != nil {
		err = fmt.Errorf("题库ID错误")
		return err
	}
	questionBankCorrectionModel := new(models.QuestionBankCorrection)
	questionBankCorrectionModel.QuestionBankId = id
	questionBankCorrectionModel.UserId = userId
	err = databases.NewDB().FirstOrCreate(questionBankCorrectionModel, map[string]interface{}{
		"question_bank_id": id,
		"user_id":          userId,
	}).Error
	//添加记录
	addAnswerRecord(id, userId, "", models.AnswerRecordUndefined)
	return
}

//添加答题记录
func addAnswerRecord(id, userId uint, answer string, result int) {
	answerRecord := new(models.AnswerRecord)
	answerRecord.UserId = userId
	answerRecord.QuestionBankId = id
	answerRecord.Answer = answer
	answerRecord.Result = result
	recordErr := databases.NewDB().Save(answerRecord).Error
	//添加记录的error不做返回值了
	if recordErr != nil {
		logrus.Error("recordErr: ", recordErr)
	}
	return
}
