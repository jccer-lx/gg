package params

import "github.com/sirupsen/logrus"

type GGParams interface {
	NewParams() GGParams
}

type NullParams struct {
}

func (p *NullParams) NewParams() GGParams {
	return &NullParams{}
}

var ggRouterParamsList = make(map[string]GGParams)

//handler 与 参数匹配
func InitParams(handlerFuncName string, p GGParams) {
	ggRouterParamsList[handlerFuncName] = p
}

func NewRouterParamInterface(handlerFuncName string) GGParams {
	logrus.Info("handlerFuncName", handlerFuncName)
	if ggRouterParamsList[handlerFuncName] == nil {
		return nil
	}
	return ggRouterParamsList[handlerFuncName].(GGParams).NewParams()
}
