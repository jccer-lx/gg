package handlers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
)

func init() {
	params.InitParams("github.com/lvxin0315/gg/handlers.AdminAddApi", &params.AdminAddApiParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.AdminUpdateApi", &params.AdminUpdateApiParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.LoginApi", &loginApiParams{})
}

type loginApiParams struct {
	Username string `from:"username" json:"username" validate:"required,min=6,max=30" label:"账号"`
	Password string `from:"password" json:"password" validate:"required,min=6,max=20" label:"密码"`
}

func (p *loginApiParams) NewParams() params.GGParams {
	return &loginApiParams{}
}

func AdminListApi(c *gin.Context) {
	output := ggOutput(c)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	adminModel := new(models.Admin)
	var adminList []*models.Admin
	pagination.Data = &adminList
	err = services.GetList(adminModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = adminList
	output.Count = pagination.Count
}

func AdminAddApi(c *gin.Context) {
	adminModel := new(models.Admin)
	adminParams := c.Keys["params"].(*params.AdminAddApiParams)
	adminModel.Username = adminParams.Username
	adminModel.Nickname = adminParams.Nickname
	adminModel.Password = adminParams.Password
	adminModel.Email = adminParams.Email
	err := services.AddAdmin(adminModel)
	if err != nil {
		setGGError(c, err)
		return
	}
}

func AdminGetApi(c *gin.Context) {
	output := ggOutput(c)
	id := c.Param("id")
	if id == "" {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	adminModel := new(models.Admin)
	output.Data = adminModel
	err := services.GetOne(adminModel, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		setGGError(c, err)
		return
	}
	//密码和盐去掉哦~
	adminModel.Password = ""
	adminModel.Salt = ""
}

func AdminUpdateApi(c *gin.Context) {
	output := ggOutput(c)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	adminModel := new(models.Admin)
	adminParams := c.Keys["params"].(*params.AdminUpdateApiParams)
	adminModel.ID = id
	adminModel.Nickname = adminParams.Nickname
	adminModel.Email = adminParams.Email
	err := services.UpdateOne(adminModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = adminModel
}

//登录
func LoginApi(c *gin.Context) {
	output := ggOutput(c)
	loginParams := c.Keys["params"].(*loginApiParams)
	adminModel, err := services.Login(loginParams.Username, loginParams.Password)
	if err != nil {
		setGGError(c, err)
		return
	}
	//密码和盐去掉哦~
	adminModel.Password = ""
	adminModel.Salt = ""
	//保存session
	session := sessions.Default(c)
	session.Set("token", adminModel.Token)
	session.Save()
	logrus.Info("login token:", session.Get("token"))
	output.Data = adminModel
}

func LogoutApi(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("token")
	session.Clear()
	session.Save()
	helper.DeleteToken(token.(string))
}
