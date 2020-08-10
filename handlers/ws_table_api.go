package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
)

func init() {
	params.InitParams("github.com/lvxin0315/gg/handlers.WsTableAddApi", &params.WsTableAddParams{})
}

//{field: "f1", title: "字段9", edit: "text",event: "f9"},
type WsTableField struct {
	Field string `json:"field"`
	Title string `json:"title"`
	Edit  string `json:"edit"`
	Event string `json:"event"`
}

//ws表格列表api
func WsTableListApi(c *gin.Context) {
	output := ggOutput(c)
	wsTableModel := new(models.WsTable)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	var wsTableList []*models.WsTable
	pagination.Data = &wsTableList
	err = services.GetList(wsTableModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = wsTableList
	output.Count = pagination.Count
}

//添加ws表格api
func WsTableAddApi(c *gin.Context) {
	output := ggOutput(c)
	wsTableModel := new(models.WsTable)
	wsTableAddParams := ggParams(c).(*params.WsTableAddParams)
	helper.ReflectiveStructToStruct(wsTableModel, wsTableAddParams)
	logrus.Info("wsTableModel.Title:", wsTableModel.Title)
	err := services.SaveOne(wsTableModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = wsTableModel
}

//获取ws表格数据
func WsTableDataApi(c *gin.Context) {
	output := ggOutput(c)
	wsTableDataModel := new(models.WsTableData)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	var wsTableDataList []*models.WsTableData
	pagination.Data = &wsTableDataList
	err = services.GetList(wsTableDataModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = wsTableDataList
	output.Count = pagination.Count
}

//获取ws表格字段
func WsTableFieldsApi(c *gin.Context) {
	wsTableDataModel := new(models.WsTableData)
	var wsTableFieldList []*WsTableField

	for _, f := range databases.NewDB().NewScope(wsTableDataModel).Fields() {
		if !checkIgnoreFiled(f.DBName) {
			continue
		}
		//id 不可编辑
		if f.DBName == "id" {
			wsTableFieldList = append(wsTableFieldList, &WsTableField{
				Field: f.DBName,
				Title: f.Name,
				Edit:  "",
				Event: f.DBName,
			})
		} else {
			wsTableFieldList = append(wsTableFieldList, &WsTableField{
				Field: f.DBName,
				Title: f.Name,
				Edit:  "text",
				Event: f.DBName,
			})
		}
	}
	output := ggOutput(c)
	output.Data = wsTableFieldList
}

func checkIgnoreFiled(fName string) bool {
	ignoreFields := []string{
		"created_at", "updated_at", "deleted_at",
	}
	for _, ignore := range ignoreFields {
		if ignore == fName {
			return false
		}
	}
	return true
}
