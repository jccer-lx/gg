package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorPage(c *gin.Context, err error) {
	data := make(map[string]interface{})
	data["ErrorMsg"] = err.Error()
	c.HTML(http.StatusInternalServerError, "500.tpl", data)
}
