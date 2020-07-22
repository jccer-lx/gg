package helper

import (
	"github.com/sirupsen/logrus"
)

type Output struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
	Err   error
}

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
	Where map[string]interface{}
	Count int
	Data  interface{}
}

func (op *Output) ReturnOutput() *Output {
	//如果有错误信息
	if op.Err != nil {
		logrus.Error(op.Err)
		if op.Code < 1 {
			op.Code = 1
		}
		if op.Msg == "" {
			op.Msg = op.Err.Error()
		}
	}
	return op
}
