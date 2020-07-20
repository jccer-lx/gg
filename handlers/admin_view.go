package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/impl"
	"github.com/lvxin0315/gg/models"
	"net/http"
	"time"
)

type AdminTableData struct {
	ApiUrl    string            `json:"api_url"`
	Id        string            `json:"id"`
	FieldList []*AdminFieldData `json:"field_list"`
}

type AdminFieldData struct {
	Field string `json:"field"`
	Title string `json:"title"`
	Sort  string `json:"sort"`
}

var adminTableMap map[string]impl.AdminModelImpl

func init() {
	adminTableMap = make(map[string]impl.AdminModelImpl)
	adminTableMap["list"] = &models.Admin{}
	adminTableMap["list1"] = &models.Admin{}
	adminTableMap["list2"] = &models.Admin{}
}

func AdminIndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "layout/common.html", nil)
}

func AdminView(c *gin.Context) {
	m := c.Param("view")
	adminTableData := new(AdminTableData)
	//map中取出对应结构
	fl, err := adminTableMap[m].GetTableFields()
	if err != nil {
		ErrorView(c)
		return
	}
	//字段
	for _, f := range fl {
		adminTableData.FieldList = append(adminTableData.FieldList, &AdminFieldData{
			Field: f.JsonTitle,
			Title: f.Title,
			Sort:  "true",
		})
	}
	//接口地址
	adminTableData.ApiUrl = fmt.Sprintf("/admin/api/%s", m)
	adminTableData.Id = fmt.Sprintf("%X", time.Now().UnixNano())
	c.HTML(http.StatusOK, "admin/list.html", adminTableData)
}
