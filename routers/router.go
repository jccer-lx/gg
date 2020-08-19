package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/handlers"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/", handlers.IndexView)
	engine.GET("/view/:module/:action", handlers.GGView)
	engine.GET("/view/:module/:action/:id", handlers.GGView)

	engine.GET("/index/api/data", handlers.IndexApi)
	//public
	public := engine.Group("/public")
	public.POST("/api/login", handlers.LoginApi)
	public.GET("/api/logout", handlers.LogoutApi)
	//系统日志
	systemLog := engine.Group("/system_log")
	systemLog.GET("/api/list", handlers.SystemLogListApi)
	//附件
	upload := engine.Group("/upload")
	upload.POST("/api/pic", handlers.UploadPicApi)
	upload.POST("/api/layedit_pic", handlers.LayEditUploadPicApi)
	//管理员
	admin := engine.Group("/admin")
	admin.GET("/api/list", handlers.AdminListApi)
	admin.POST("/api/add", handlers.AdminAddApi)
	admin.GET("/api/edit/:id", handlers.AdminGetApi)
	admin.PUT("/api/edit/:id", handlers.AdminUpdateApi)

	//权限&菜单
	authRule := engine.Group("/auth_rule")
	authRule.GET("/api/list", handlers.AuthRuleListApi)
	authRule.GET("/api/all", handlers.AuthRuleAllListApi)
	authRule.POST("/api/add", handlers.AuthRuleAddApi)
	authRule.PUT("/api/edit/:id", handlers.AuthRuleUpdateApi)
	authRule.GET("/api/get_menus", handlers.MenuApi)

	//商品分类
	goodsCategory := engine.Group("/goods_category")
	goodsCategory.GET("/api/list", handlers.GoodsCategoryListApi)
	goodsCategory.GET("/api/all", handlers.GoodsCategoryAllListApi)
	goodsCategory.POST("/api/add", handlers.GoodsCategoryAddApi)
	goodsCategory.PUT("/api/edit/:id", handlers.GoodsCategoryUpdateApi)
	goodsCategory.PUT("/api/show_pic/:id", handlers.GoodsCategoryUpdatePicApi)
	goodsCategory.GET("/api/get/:id", handlers.GoodsCategoryGetApi)
	goodsCategory.DELETE("/api/delete", handlers.GoodsCategoryDeleteApi)
	//商品
	goods := engine.Group("/goods")
	goods.GET("/api/list", handlers.GoodsListApi)
	goods.POST("/api/add", handlers.GoodsAddApi)
	goods.PUT("/api/edit/:id", handlers.GoodsUpdateApi)
	goods.PUT("/api/update_for_field", handlers.UpdateGoodsForFieldApi)
	goods.GET("/api/get/:id", handlers.GoodsGetApi)
	//会员
	member := engine.Group("/member")
	member.GET("/api/list", handlers.MemberListApi)

	money := engine.Group("/money")
	money.POST("/api/add", handlers.AddMoneyApi)

	//jd商品
	jdGoods := engine.Group("/jd_goods")
	jdGoods.GET("/api/list", handlers.JdGoodsListApi)
	//wsTable
	wsTable := engine.Group("/ws_table")
	//ws
	wsTable.GET("/wst", handlers.WsTable)
	wsTable.GET("/api/list", handlers.WsTableListApi)
	wsTable.POST("/api/add", handlers.WsTableAddApi)
	wsTable.GET("/api/wt_data/:id", handlers.WsTableDataApi)
	wsTable.GET("/api/get_wt_fields/:id", handlers.WsTableFieldsApi)
	//plane
	plane := engine.Group("/plane")
	plane.GET("/api/list", handlers.PlaneListApi)
	plane.GET("/api/user_plane_col_list", handlers.GetUserAllPlaneListApi)
	plane.POST("/api/plane_coordinate", handlers.PlaneCoordinateApi)
	plane.POST("/api/save_plane", handlers.SavePlaneApi)
	plane.GET("/ws_plane", handlers.PlaneWs)

}
