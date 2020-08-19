package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/services"
)

func init() {
	params.InitParams("github.com/lvxin0315/gg/handlers.AuthRuleAddApi", &params.AuthRuleParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.AuthRuleUpdateApi", &params.AuthRuleUpdateParams{})
}

func AuthRuleListApi(c *gin.Context) {
	pagination := new(helper.Pagination)
	pagination.Data = &[]models.AuthRule{}
	ggList(c, &models.AuthRule{}, pagination)
}

func AuthRuleAddApi(c *gin.Context) {
	authRuleModel := new(models.AuthRule)
	authRuleParams := ggParams(c).(*params.AuthRuleParams)
	helper.ReflectiveStructToStructWithJson(authRuleModel, authRuleParams)
	err := services.AddAuthRule(authRuleModel)
	if err != nil {
		setGGError(c, err)
		return
	}
}

//获取所有顶级菜单
func AuthRuleAllListApi(c *gin.Context) {
	output := ggOutput(c)
	authRuleModel := new(models.AuthRule)
	var authRuleList []*models.AuthRule
	err := services.GetAllList(authRuleModel, &authRuleList)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = authRuleList
}

func AuthRuleUpdateApi(c *gin.Context) {
	output := ggOutput(c)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	authRuleModel := new(models.AuthRule)
	authRuleUpdateParams := ggParams(c).(*params.AuthRuleUpdateParams)
	helper.ReflectiveStructToStructWithJson(authRuleModel, authRuleUpdateParams)
	authRuleModel.ID = id
	err := services.UpdateOne(authRuleModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = authRuleModel
}

func MenuApi(c *gin.Context) {
	authRuleModelList, err := services.GetAuthRuleListWithChildren()
	if err != nil {
		setGGError(c, err)
		return
	}
	output := ggOutput(c)
	output.Data = authRuleModelList
}
