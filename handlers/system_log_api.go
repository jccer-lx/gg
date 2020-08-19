package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
)

func SystemLogListApi(c *gin.Context) {
	pagination := new(helper.Pagination)
	pagination.Data = &[]models.SystemLog{}
	ggList(c, &models.SystemLog{}, pagination)
}
