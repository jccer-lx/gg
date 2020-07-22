package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/services"
	"net/http"
)

//使用自定义结构体接受参数，防止直接操作model的意外惊喜
type AdminParams struct {
	Username string `from:"username" json:"username"`
	Nickname string `from:"nickname" json:"nickname"`
	Password string `from:"password" json:"password"`
	Salt     string `from:"salt" json:"salt"`
	Avatar   string `from:"avatar" json:"avatar"`
	Email    string `from:"email" json:"email"`
}

func AdminListApi(c *gin.Context) {
	output := new(helper.Output)
	defer c.JSON(http.StatusOK, output.ReturnOutput())
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		output.Err = err
		return
	}
	adminModel := new(models.Admin)
	var adminList []*models.Admin
	pagination.Data = &adminList
	err = services.GetList(adminModel, pagination)
	if err != nil {
		output.Err = err
		return
	}
	output.Data = adminList
	output.Count = pagination.Count
}

func AdminAddApi(c *gin.Context) {
	output := new(helper.Output)
	defer c.JSON(http.StatusOK, output.ReturnOutput())
	adminModel := new(models.Admin)
	adminParams := new(AdminParams)
	err := c.ShouldBind(adminParams)
	if err != nil {
		output.Err = err
		return
	}
	adminModel.Username = adminParams.Username
	adminModel.Nickname = adminParams.Nickname
	adminModel.Password = adminParams.Password
	adminModel.Email = adminParams.Email
	err = services.AddAdmin(adminModel)
	if err != nil {
		output.Err = err
		return
	}
}

func AdminGetApi(c *gin.Context) {
	output := new(helper.Output)
	defer c.JSON(http.StatusOK, output.ReturnOutput())
	id := c.Param("id")
	if id == "" {
		output.Err = fmt.Errorf("参数异常")
		return
	}
	adminModel := new(models.Admin)
	output.Data = adminModel
	err := services.GetOne(adminModel, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		output.Err = err
		return
	}
	//密码和盐去掉哦~
	adminModel.Password = ""
	adminModel.Salt = ""
}

func AdminUpdateApi(c *gin.Context) {
	output := new(helper.Output)
	defer apiReturn(c, output)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		output.Err = fmt.Errorf("参数异常")
		return
	}
	adminModel := new(models.Admin)
	adminParams := new(AdminParams)
	err := c.ShouldBind(adminParams)
	if err != nil {
		fmt.Println(err)
		output.Err = err
		return
	}
	adminModel.ID = id
	adminModel.Nickname = adminParams.Nickname
	adminModel.Email = adminParams.Email
	err = services.UpdateOne(adminModel)
	if err != nil {
		output.Err = err
		return
	}
	output.Data = adminModel
}
