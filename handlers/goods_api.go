package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/services"
)

func GoodsListApi(c *gin.Context) {
	output := ggOutput(c)
	GoodsModel := new(models.Goods)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	var GoodsList []*models.Goods
	pagination.Data = &GoodsList
	err = services.GetList(GoodsModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = GoodsList
	output.Count = pagination.Count
}
