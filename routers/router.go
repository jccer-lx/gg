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

	//系统日志
	systemLog := engine.Group("/system_log")
	systemLog.GET("/view/list", handlers.GGView)

	systemLog.GET("/api/list", handlers.SystemLogListApi)

	//附件
	upload := engine.Group("/upload")
	upload.POST("/api/pic", handlers.UploadPicApi)
	upload.POST("/api/layedit_pic", handlers.LayEditUploadPicApi)

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
	goodsCategory.GET("/view/show_pic/:id", handlers.GGView)

	goodsCategory.GET("/api/list", handlers.GoodsCategoryListApi)
	goodsCategory.GET("/api/all", handlers.GoodsCategoryAllListApi)
	goodsCategory.POST("/api/add", handlers.GoodsCategoryAddApi)
	goodsCategory.PUT("/api/edit/:id", handlers.GoodsCategoryUpdateApi)
	goodsCategory.PUT("/api/show_pic/:id", handlers.GoodsCategoryUpdatePicApi)
	goodsCategory.GET("/api/get/:id", handlers.GoodsCategoryGetApi)
	goodsCategory.DELETE("/api/delete", handlers.GoodsCategoryDeleteApi)

	//商品
	goods := engine.Group("/goods")
	goods.GET("/view/list", handlers.GGView)
	goods.GET("/view/add", handlers.GGView)
	goods.GET("/view/show_main_image/:id", handlers.GGView)
	goods.GET("/view/show_slider_image/:id", handlers.GGView)

	goods.GET("/api/list", handlers.GoodsListApi)
	goods.POST("/api/add", handlers.GoodsAddApi)
	goods.PUT("/api/edit/:id", handlers.GoodsUpdateApi)
	goods.PUT("/api/update_for_field", handlers.UpdateGoodsForFieldApi)
	goods.GET("/api/get/:id", handlers.GoodsGetApi)

	//jd商品
	jdGoods := engine.Group("/jd_goods")
	jdGoods.GET("/view/list", handlers.GGView)

	jdGoods.GET("/api/list", handlers.JdGoodsListApi)

	//wsTable
	wsTable := engine.Group("/ws_table")
	wsTable.GET("/view/list", handlers.GGView)
	wsTable.GET("/view/add", handlers.GGView)
	wsTable.GET("/view/wt/:id", handlers.GGView)
	//ws
	wsTable.GET("/wst", handlers.WsTable)

	wsTable.GET("/api/list", handlers.WsTableListApi)
	wsTable.POST("/api/add", handlers.WsTableAddApi)

	wsTable.GET("/api/wt_data/:id", handlers.WsTableDataApi)
	wsTable.GET("/api/get_wt_fields/:id", handlers.WsTableFieldsApi)

}
