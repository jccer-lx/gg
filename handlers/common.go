package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/middlewares"
	"github.com/lvxin0315/gg/params"
	"net/http"
	"strings"
)

//通用404页面
func errorView(c *gin.Context) {
	c.HTML(http.StatusNotFound, "layout/404.html", nil)
}

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "layout/common.html", nil)
}

//api通用返回值
func apiReturn(c *gin.Context, op *helper.Output) {
	op.ReturnOutput()
	c.JSON(http.StatusOK, op)
}

//通用接受参数方法
func ggBindParams(c *gin.Context, data params.GGParams) error {
	return middlewares.GGBindParams(c, data)
}

//通用的view
func GGView(c *gin.Context) {
	uriList := strings.Split(c.Request.RequestURI, "/")
	if len(uriList) < 3 || uriList[2] != "view" {
		return
	}
	//参数[id]
	data := map[string]interface{}{
		"ID": c.Param("id"),
	}
	c.HTML(http.StatusOK, fmt.Sprintf("%s/%s.tpl", uriList[1], uriList[3]), data)
}

//通用的获取output
func ggOutput(c *gin.Context) *helper.Output {
	return c.Keys["output"].(*helper.Output)
}

//通用的设置返回error
func setGGError(c *gin.Context, err error) {
	middlewares.SetGGError(c, err)
}
