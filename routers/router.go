package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/handlers"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/", handlers.IndexView)
	//public
	public := engine.Group("/public")
	public.GET("/view/login", handlers.GGView)

	public.POST("/api/login", handlers.LoginApi)
	public.GET("/api/logout", handlers.LogoutApi)

	//管理员
	admin := engine.Group("/admin")
	admin.GET("/view/list", handlers.GGView)
	admin.GET("/view/add", handlers.GGView)

	admin.GET("/api/list", handlers.AdminListApi)
	admin.POST("/api/add", handlers.AdminAddApi)
	admin.GET("/api/edit/:id", handlers.AdminGetApi)
	admin.PUT("/api/edit/:id", handlers.AdminUpdateApi)

}
