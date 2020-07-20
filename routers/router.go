package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/handlers"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/", handlers.AdminIndexView)
	admin := engine.Group("/admin")
	admin.GET("/tpl/:view", handlers.AdminView)
	admin.GET("/api/:view", handlers.AdminApi)
}
