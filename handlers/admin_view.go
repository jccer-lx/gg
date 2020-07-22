package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminListView(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/list.tpl", nil)
}

func AdminAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/add.tpl", nil)
}

func AdminEditView(c *gin.Context) {
	data := map[string]interface{}{
		"ID": c.Param("id"),
	}
	c.HTML(http.StatusOK, "admin/edit.tpl", data)
}
