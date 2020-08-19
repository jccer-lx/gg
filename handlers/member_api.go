package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/services"
)

func init() {
	params.InitParams("github.com/lvxin0315/gg/handlers.AddMoneyApi", &addMoneyApiParams{})
}

type addMoneyApiParams struct {
	MemberId json.Number `json:"member_id" validate:"required"`
	Money    json.Number `json:"money" validate:"required"`
}

func (p *addMoneyApiParams) NewParams() params.GGParams {
	return &addMoneyApiParams{}
}

func MemberListApi(c *gin.Context) {
	output := ggOutput(c)
	memberModel := new(models.Member)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	var memberList []*models.Member
	pagination.Data = &memberList
	err = services.GetList(memberModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = memberList
	output.Count = pagination.Count
}

//添加余额
func AddMoneyApi(c *gin.Context) {
	p := ggParams(c).(*addMoneyApiParams)
	_, err := services.Recharge(helper.JsonNumber2Uint(p.MemberId), helper.JsonNumber2Float64(p.Money))
	if err != nil {
		setGGError(c, err)
		return
	}
}
