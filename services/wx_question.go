package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const wxMessageBr = "\r\n"

/**
判断是否微信答题状态
@param string openid
@return bool
*/
func IsExamination(openid string) bool {
	logrus.Info(openid)
	return true
}

/**
保存微信答题状态，openid => questionId
@param string openid
@param uint questionId 问题ID
*/
func SetQuestionForOpenid(openid string, questionId uint) error {
	//TODO 具体保存方式，待定，mysql容易，redis靠谱
	return nil
}

/**
读取微信答题状态，openid => questionId
@param string openid
@return uint questionId 问题ID
*/
func GetQuestionForOpenid(openid string) (questionId uint, err error) {
	//TODO
	return 1, nil
}

/**
下一题
*/
func NextQuestion(openid string) (string, error) {
	//TODO
	nextQuestionId := 2
	err := SetQuestionForOpenid(openid, uint(nextQuestionId))
	if err != nil {
		return "", err
	}
	return "以下哪些时候1111111111", nil
}

/**
微信答题
@param string openid
@param string answer 答案
@return string
*/
func Answer(openid string, answer string) (string, error) {
	if !IsExamination(openid) {
		errorMsg := "未开启答题状态"
		return errorMsg, fmt.Errorf(errorMsg)
	}
	res := ""
	//判断答案 TODO
	res += fmt.Sprintf("回答正确~~~%s", wxMessageBr)
	//下一题
	nextQuestionStr, err := NextQuestion(openid)
	if err != nil {
		return "查询题目错误", err
	}
	res += nextQuestionStr
	return res, nil
}
