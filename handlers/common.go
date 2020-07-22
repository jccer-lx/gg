package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/validation"
	"github.com/sirupsen/logrus"
	"net/http"
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
func params(c *gin.Context, data interface{}) error {
	err := c.ShouldBind(data)
	if err != nil {
		logrus.Error("ShouldBind:", err)
		return err
	}
	//validate
	err = validation.Check(data)
	if err != nil {
		logrus.Error("Validation:", err)
		return err
	}
	return nil
}
