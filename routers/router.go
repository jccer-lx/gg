package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/handlers"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/", handlers.IndexView)
	engine.GET("/index/api/data", handlers.IndexApi)
	//public
	public := engine.Group("/public")
	public.GET("/view/login", handlers.GGView)

	public.POST("/api/login", handlers.LoginApi)
	public.GET("/api/logout", handlers.LogoutApi)

	//附件
	upload := engine.Group("/upload")
	upload.POST("/api/pic", handlers.UploadPic)

	//管理员
	admin := engine.Group("/admin")
	admin.GET("/view/list", handlers.GGView)
	admin.GET("/view/add", handlers.GGView)

	admin.GET("/api/list", handlers.AdminListApi)
	admin.POST("/api/add", handlers.AdminAddApi)
	admin.GET("/api/edit/:id", handlers.AdminGetApi)
	admin.PUT("/api/edit/:id", handlers.AdminUpdateApi)

	//权限&菜单
	authRule := engine.Group("/auth_rule")
	authRule.GET("/view/list", handlers.GGView)
	authRule.GET("/view/add", handlers.GGView)

	authRule.GET("/api/list", handlers.AuthRuleListApi)
	authRule.GET("/api/all", handlers.AuthRuleAllListApi)
	authRule.POST("/api/add", handlers.AuthRuleAddApi)
	authRule.PUT("/api/edit/:id", handlers.AuthRuleUpdateApi)
	authRule.GET("/api/get_menus", handlers.MenuApi)

	//商品分类
	goodsCategory := engine.Group("/goods_category")
	goodsCategory.GET("/view/list", handlers.GGView)
	goodsCategory.GET("/view/add", handlers.GGView)

	goodsCategory.GET("/api/list", handlers.GoodsCategoryListApi)
	goodsCategory.GET("/api/all", handlers.GoodsCategoryAllListApi)
	goodsCategory.POST("/api/add", handlers.GoodsCategoryAddApi)
	goodsCategory.PUT("/api/edit/:id", handlers.GoodsCategoryUpdateApi)
}
