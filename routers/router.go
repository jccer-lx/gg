package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/handlers"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/", handlers.IndexView)
	admin := engine.Group("/admin")
	admin.GET("/view/list", handlers.AdminListView)
	admin.GET("/view/add", handlers.AdminAddView)
	admin.GET("/view/edit/:id", handlers.AdminEditView)

	admin.GET("/api/list", handlers.AdminListApi)
	admin.POST("/api/add", handlers.AdminAddApi)
	admin.GET("/api/edit/:id", handlers.AdminGetApi)
	admin.PUT("/api/edit/:id", handlers.AdminUpdateApi)

}
