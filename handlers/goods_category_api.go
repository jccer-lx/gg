package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
)

func init() {
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsCategoryAddApi", &params.GoodsCategoryParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsCategoryUpdateApi", &params.GoodsCategoryUpdateParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsCategoryUpdatePicApi", &params.GoodsCategoryUpdatePicParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsCategoryDeleteApi", &params.GoodsCategoryDeleteParams{})
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
	goodsCategoryModel := new(models.GoodsCategory)
	var goodsCategoryList []*models.GoodsCategory
	err := services.GetAllList(goodsCategoryModel, &goodsCategoryList)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = goodsCategoryList
}

func GoodsCategoryAddApi(c *gin.Context) {
	goodsCategoryModel := new(models.GoodsCategory)
	goodsCategoryParams := c.Keys["params"].(*params.GoodsCategoryParams)
	goodsCategoryModel.CateName = goodsCategoryParams.CateName
	goodsCategoryModel.Pic = goodsCategoryParams.Pic
	goodsCategoryModel.Pid = helper.JsonNumber2Uint(goodsCategoryParams.Pid)
	goodsCategoryModel.Sort = helper.JsonNumber2Int(goodsCategoryParams.Sort)
	isShowSwitch := goodsCategoryParams.IsShow
	if isShowSwitch == "on" {
		goodsCategoryModel.IsShow = models.GoodsCategoryIsShowTrue
	} else {
		goodsCategoryModel.IsShow = models.GoodsCategoryIsShowFalse
	}
	err := services.AddGoodsCategory(goodsCategoryModel)
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
	goodsCategoryModel := new(models.GoodsCategory)
	goodsCategoryUpdateParams := c.Keys["params"].(*params.GoodsCategoryUpdateParams)
	goodsCategoryModel.ID = id
	goodsCategoryModel.CateName = goodsCategoryUpdateParams.CateName
	goodsCategoryModel.Pic = goodsCategoryUpdateParams.Pic
	goodsCategoryModel.IsShow = helper.JsonNumber2Int(goodsCategoryUpdateParams.IsShow)
	goodsCategoryModel.Sort = helper.JsonNumber2Int(goodsCategoryUpdateParams.Sort)
	err := services.UpdateOne(goodsCategoryModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = goodsCategoryModel
}

func GoodsCategoryGetApi(c *gin.Context) {
	output := ggOutput(c)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	goodsCategoryModel := new(models.GoodsCategory)
	err := services.GetOne(goodsCategoryModel, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		setGGError(c, err)
	}
	output.Data = goodsCategoryModel
}

func GoodsCategoryUpdatePicApi(c *gin.Context) {
	output := ggOutput(c)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	goodsCategoryUpdatePicParams := c.Keys["params"].(*params.GoodsCategoryUpdatePicParams)
	goodsCategoryModel := new(models.GoodsCategory)
	goodsCategoryModel.ID = id
	goodsCategoryModel.Pic = goodsCategoryUpdatePicParams.Pic
	logrus.Info("pic:", goodsCategoryModel.Pic)
	err := services.UpdateOne(goodsCategoryModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = goodsCategoryModel
}

func GoodsCategoryDeleteApi(c *gin.Context) {
	goodsCategoryDeleteParams := c.Keys["params"].(*params.GoodsCategoryDeleteParams)
	goodsCategoryModel := new(models.GoodsCategory)
	err := services.DeleteByIds(goodsCategoryModel, goodsCategoryDeleteParams.IDList)
	if err != nil {
		setGGError(c, err)
		return
	}
}
