package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/etc"
	"github.com/lvxin0315/gg/middlewares"
	"github.com/lvxin0315/gg/routers"
	"github.com/lvxin0315/gg/validation"
	"github.com/sirupsen/logrus"
)

func main() {
	engine := gin.Default()

	//中间件-跨域
	engine.Use(middlewares.CorsMiddleware)
	//中间件-session
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("gg-session", store))
	engine.Use(middlewares.SessionMiddleware)
	//加载路由
	routers.InitRouter(engine)
	//自定义函数
	//engine.SetFuncMap(template.FuncMap{})
	//静态资源
	engine.Static("/assets", "assets")
	engine.LoadHTMLGlob("views/**/*")
	//加载db
	databases.InitMysqlDB()
	databases.InitMemDB()
	//validate init
	err := validation.InitValidate()
	if err != nil {
		panic(err)
	}
	//debug?
	logrus.SetLevel(logrus.DebugLevel)
	engine.Run(fmt.Sprintf(":%s", etc.Config.Port))
}
