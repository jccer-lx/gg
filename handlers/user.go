package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/services"
	"net/http"
)

func UserInfoByOpenid(c *gin.Context) {
	openid := c.Param("openid")
	if openid == "" {
		ErrorPage(c, fmt.Errorf("openid is error"))
		return
	}
	userModel, err := services.GetUserInfoByOpenid(openid)
	if err != nil {
		ErrorPage(c, err)
		return
	}
	//用户的支付码
	paymentCodeModel, err := services.PayCode(userModel.Id)
	if err != nil {
		ErrorPage(c, err)
		return
	}
	data := make(map[string]interface{})
	data["Openid"] = userModel.Openid
	data["Score"] = userModel.Score
	data["Money"] = userModel.Money
	data["QrCodePath"] = paymentCodeModel.QrCodePath
	c.HTML(http.StatusOK, "user/user_info.tpl", data)
}
