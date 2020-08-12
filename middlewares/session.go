package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/sirupsen/logrus"
	"net/http"
)

//处理会话情况
func SessionMiddleware(c *gin.Context) {
	//public的内容不考虑
	uriList := helper.GetUriStringList(c)
	logrus.Info("uriList[1]:", uriList[1])
	if uriList[1] == "public" || uriList[1] == "assets" {
		return
	}
	session := sessions.Default(c)
	token := session.Get("token")
	logrus.Info("token:", token)
	if token == nil {
		redirectMiddleware(c)
		return
	}
	//session，具体是把token放到keys里
	c.Keys["token"] = token.(string)
}

//重定向到登录或ajax返回未登录
func redirectMiddleware(c *gin.Context) {
	if isAjax(c) {
		output := new(helper.Output)
		output.UnLogin()
		c.JSON(http.StatusTemporaryRedirect, output)
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/public/view/login")
	}
	c.Abort()
}

func isAjax(c *gin.Context) bool {
	return c.GetHeader("X-Requested-With") == "XMLHttpRequest"
}
