package helper

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetUriStringList(c *gin.Context) []string {
	return strings.Split(c.Request.RequestURI, "/")
}
