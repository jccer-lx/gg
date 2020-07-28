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
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsCategoryAddApi", &params.GoodsCategoryParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsCategoryUpdateApi", &params.GoodsCategoryUpdateParams{})
}

func GoodsCategoryListApi(c *gin.Context) {
	output := ggOutput(c)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	goodsCategoryModel := new(models.GoodsCategory)
	var goodsCategoryList []*models.GoodsCategory
	pagination.Data = &goodsCategoryList
	err = services.GetGoodsCategoryList(goodsCategoryModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = goodsCategoryList
	output.Count = pagination.Count
}

func GoodsCategoryAllListApi(c *gin.Context) {
	output := ggOutput(c)
	GoodsCategoryModel := new(models.GoodsCategory)
	var GoodsCategoryList []*models.GoodsCategory
	err := services.GetAllList(GoodsCategoryModel, &GoodsCategoryList)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = GoodsCategoryList
}

func GoodsCategoryAddApi(c *gin.Context) {
	GoodsCategoryModel := new(models.GoodsCategory)
	GoodsCategoryParams := c.Keys["params"].(*params.GoodsCategoryParams)
	GoodsCategoryModel.CateName = GoodsCategoryParams.CateName
	GoodsCategoryModel.Pic = GoodsCategoryParams.Pic
	GoodsCategoryModel.Pid = helper.JsonNumber2Uint(GoodsCategoryParams.Pid)
	GoodsCategoryModel.Sort = helper.JsonNumber2Int(GoodsCategoryParams.Sort)
	isShowSwitch := GoodsCategoryParams.IsShow
	if isShowSwitch == "on" {
		GoodsCategoryModel.IsShow = models.GoodsCategoryIsShowTrue
	} else {
		GoodsCategoryModel.IsShow = models.GoodsCategoryIsShowFalse
	}
	err := services.AddGoodsCategory(GoodsCategoryModel)
	if err != nil {
		setGGError(c, err)
		return
	}
}

func GoodsCategoryUpdateApi(c *gin.Context) {
	output := ggOutput(c)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	GoodsCategoryModel := new(models.GoodsCategory)
	GoodsCategoryUpdateParams := c.Keys["params"].(*params.GoodsCategoryUpdateParams)
	GoodsCategoryModel.ID = id
	GoodsCategoryModel.CateName = GoodsCategoryUpdateParams.CateName
	GoodsCategoryModel.Pic = GoodsCategoryUpdateParams.Pic
	GoodsCategoryModel.IsShow = helper.JsonNumber2Int(GoodsCategoryUpdateParams.IsShow)
	GoodsCategoryModel.Sort = helper.JsonNumber2Int(GoodsCategoryUpdateParams.Sort)
	err := services.UpdateOne(GoodsCategoryModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = GoodsCategoryModel
}
