package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/handlers"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/", handlers.AdminIndexView)
	admin := engine.Group("/admin")
	admin.GET("/tpl/:view/list", handlers.AdminListView)
	admin.GET("/tpl/:view/add", handlers.AdminAddView)
	admin.GET("/tpl/:view/edit", handlers.AdminEditView)
	admin.GET("/api/:view", handlers.AdminListApi)
	admin.POST("/api/:view", handlers.AdminAddApi)
	admin.PUT("/api/:view", handlers.AdminUpdateApi)
}
