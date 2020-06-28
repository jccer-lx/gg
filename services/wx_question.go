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
	return `①足球进校园，让孩子们能在绿茵场上尽情奔跑，同时，也要尊重一部分孩子不喜欢足球的_____________。
②老同学找他办事，他很为难地说：“我就这么点______，解决不了那么大的问题，你还是想想别的办法吧。”
③中国保监会近日召开会议，强调通过开展治理理赔难等三项措施，切实保护好保险消费者合法______。
A.权力
权益
权利
B.权力
权利
权益
C.权利
权益
权力
D.权利
权力
权益`, nil
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
