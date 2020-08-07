package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
)

func JdGoodsListApi(c *gin.Context) {
	output := ggOutput(c)
	goodsModel := new(models.JdGoods)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	var goodsList []*models.JdGoods
	pagination.Data = &goodsList
	db := databases.NewDB()
	err = db.Model(goodsModel).
		Offset(pagination.Limit * (pagination.Page - 1)).
		Limit(pagination.Limit).
		Find(pagination.Data).Error
	if err != nil {
		setGGError(c, err)
		return
	}
	err = db.Model(goodsModel).Count(&pagination.Count).Error
	if err != nil {
		setGGError(c, err)
		return
	}

	output.Data = goodsList
	output.Count = pagination.Count
}
