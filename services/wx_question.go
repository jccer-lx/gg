package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
)

const wxMessageBr = "\r\n"
const choiceStemTpl = `%s
A:%s
B:%s
C:%s
D:%s
`

type WxQuestionService struct {
	Openid    string
	userModel *models.User
}

func NewWxQuestionService(openid string) *WxQuestionService {
	return &WxQuestionService{
		Openid: openid,
	}
}

func (s *WxQuestionService) GetUserModel() (userModel *models.User, err error) {
	if s.userModel == nil {
		userModel, err = SaveOpenid(s.Openid)
		if err != nil {
			return
		}
		s.userModel = userModel
	}
	userModel = s.userModel
	return
}

/**
保存微信答题状态，openid => questionId
@param string openid
@param uint questionId 问题ID
*/
func (s *WxQuestionService) SetQuestionForOpenid(questionId uint) (err error) {
	userModel, err := s.GetUserModel()
	if err != nil {
		return
	}
	userModel.LastQuestionBankId = questionId
	err = databases.NewDB().Save(userModel).Error
	return
}

/**
下一题
*/
func (s *WxQuestionService) NextQuestion() (questionStr string, err error) {
	//当前题号
	userModel, err := s.GetUserModel()
	if err != nil {
		return
	}
	questionBankModel, err := GetNextQuestionByBankId(userModel.LastQuestionBankId)
	if err != nil {
		return
	}
	//根据题目类型整理格式
	switch questionBankModel.QuestionType {
	case models.Choice:
		questionStr = s.formatChoiceQuestion(questionBankModel)
	default:
		err = fmt.Errorf("类型暂时不能支持")
	}
	//记录题号
	err = s.SetQuestionForOpenid(questionBankModel.ID)
	if err != nil {
		return
	}
	return
}

/**
微信答题
@param string openid
@param string answer 答案
@return string
*/
func (s *WxQuestionService) Answer(answer string) (res string, err error) {
	//如果用户没有正在回答题目，直接反馈
	userModel, err := s.GetUserModel()
	if err != nil {
		return
	}
	if userModel.LastQuestionBankId > 0 {
		//判断答案 TODO
		questionRes, _ := CheckQuestionAnswer(userModel.LastQuestionBankId, answer)
		if !questionRes {
			res += fmt.Sprintf("回答错误~~~%s", wxMessageBr)
		} else {
			res += fmt.Sprintf("回答正确~~~%s", wxMessageBr)
		}
	}
	//下一题
	nextQuestionStr, err := s.NextQuestion()
	if err != nil {
		res = "查询题目错误"
		return
	}
	res += nextQuestionStr
	return res, nil
}

//选择题格式整理
func (s *WxQuestionService) formatChoiceQuestion(questionBankModel *models.QuestionBank) string {
	m := questionBankModel.Question.(*models.ChoiceQuestion)
	stem := fmt.Sprintf(choiceStemTpl,
		m.Stem,
		m.Options[0].Item,
		m.Options[1].Item,
		m.Options[2].Item,
		m.Options[3].Item,
	)
	if len(m.Options) > 4 {
		stem = stem + "E:" + m.Options[4].Item
	}
	return stem
}
