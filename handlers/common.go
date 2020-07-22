package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"net/http"
)

func errorView(c *gin.Context) {
	c.HTML(http.StatusNotFound, "layout/404.html", nil)
}

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "layout/common.html", nil)
}

func apiReturn(c *gin.Context, op *helper.Output) {
	op.ReturnOutput()
	c.JSON(http.StatusOK, op)
}
