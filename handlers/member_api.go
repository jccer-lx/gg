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
	pagination := new(helper.Pagination)
	pagination.Data = &[]models.Member{}
	ggList(c, &models.Member{}, pagination)
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
