package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorView(c *gin.Context) {
	c.HTML(http.StatusNotFound, "layout/404.html", nil)
}
