package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
)

func JdGoodsListApi(c *gin.Context) {
	pagination := new(helper.Pagination)
	pagination.Data = &[]models.JdGoods{}
	ggList(c, &models.JdGoods{}, pagination, "sku_id")
}
