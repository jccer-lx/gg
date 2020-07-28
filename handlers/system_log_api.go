package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/services"
)

func SystemLogListApi(c *gin.Context) {
	output := ggOutput(c)
	systemLogModel := new(models.SystemLog)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	var systemLogList []*models.SystemLog
	pagination.Data = &systemLogList
	err = services.GetList(systemLogModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = systemLogList
	output.Count = pagination.Count
}
