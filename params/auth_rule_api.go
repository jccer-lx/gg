package params

import "encoding/json"

type AuthRuleParams struct {
	Pid    json.Number `json:"pid" `                                   // 父ID
	Name   string      `json:"name"`                                   // 规则值（路由）
	Title  string      `json:"title" validate:"required,min=1,max=10"` // 规则名称
	Remark string      `json:"remark"`                                 // 备注
	Weigh  json.Number `json:"weigh" `                                 // 权重
}

func (p *AuthRuleParams) NewParams() GGParams {
	return new(AuthRuleParams)
}

type AuthRuleUpdateParams struct {
	Name   string      `json:"name"`                                   // 规则值（路由）
	Title  string      `json:"title" validate:"required,min=1,max=10"` // 规则名称
	Remark string      `json:"remark"`                                 // 备注
	Weigh  json.Number `json:"weigh" `                                 // 权重
}

func (p *AuthRuleUpdateParams) NewParams() GGParams {
	return new(AuthRuleUpdateParams)
}
