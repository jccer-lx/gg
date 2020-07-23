package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

//处理会话情况
func SessionMiddleware(c *gin.Context) {
	uriList := strings.Split(c.Request.RequestURI, "/")
	//public的内容不考虑
	logrus.Info("uriList[1]:", uriList[1])
	if uriList[1] == "public" || uriList[1] == "assets" {
		return
	}
	session := sessions.Default(c)
	token := session.Get("token")
	logrus.Info("token:", token)
	if token == nil {
		redirectMiddleware(c)
	}
}

//重定向到登录或ajax返回未登录
func redirectMiddleware(c *gin.Context) {
	if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
		output := new(helper.Output)
		output.UnLogin()
		c.JSON(http.StatusPermanentRedirect, output)
	} else {
		c.Redirect(http.StatusPermanentRedirect, "/public/view/login")
	}
	c.Abort()
}
