package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/etc"
	"github.com/lvxin0315/gg/middlewares"
	"github.com/lvxin0315/gg/routers"
	"github.com/sirupsen/logrus"
)

func main() {
	engine := gin.Default()
	//加载路由
	routers.InitRouter(engine)
	//中间件-跨域
	engine.Use(middlewares.Cors())
	//自定义函数
	//engine.SetFuncMap(template.FuncMap{})
	//静态资源
	engine.Static("/assets", "assets")
	engine.LoadHTMLGlob("views/**/*")

	//加载db
	databases.InitMysqlDB()
	databases.InitMemDB()
	//debug?
	logrus.SetLevel(logrus.DebugLevel)
	engine.Run(fmt.Sprintf(":%s", etc.Config.Port))
}
