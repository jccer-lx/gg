package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/handlers"
)

func InitRouter(engine *gin.Engine) {
	//微信
	engine.POST("/wx", handlers.WeChat)
	//页面路由
	view := engine.Group("/v")
	//微信应用
	view.GET("/wx/user_info", handlers.UserInfoView)
	view.GET("/user/openid/:openid", handlers.UserInfoByOpenid)
}
