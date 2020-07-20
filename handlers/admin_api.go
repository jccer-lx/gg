package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/databases"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Output struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func AdminApi(c *gin.Context) {
	//参数
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")
	output := new(Output)
	defer c.JSON(http.StatusOK, output)
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		logrus.Error(err)
		output.Code = 1
		return
	}
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		logrus.Error(err)
		output.Code = 1
		return
	}
	m := c.Param("view")
	//map中取出对应结构
	model := adminTableMap[m]
	//resultData := reflect.New(model.GetSliceType()).Elem()

	err = databases.NewDB().Table(model.TableName()).Offset(limit * (page - 1)).Limit(limit).Find(model.GormFindOut()).Error
	if err != nil {
		logrus.Error(err)
		output.Code = 1
		return
	}
	output.Data = model.GormFindOut()
	output.Count = 100
	output.Msg = "nb"
}
