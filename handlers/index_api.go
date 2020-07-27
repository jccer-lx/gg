package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
)

//首页数据api
func IndexApi(c *gin.Context) {
	indexData := make(map[string]interface{})
	indexData["admin_nickname"] = getSessionNickname(c)
	output := ggOutput(c)
	output.Data = indexData
}

func getSessionNickname(c *gin.Context) string {
	token := c.Keys["token"].(string)
	if token == "" {
		return ""
	}
	adminModel, err := services.FindAdminByToken(token)
	if err != nil {
		logrus.Error("getSessionNickname error: ", err)
		return ""
	}
	return adminModel.Nickname
}
