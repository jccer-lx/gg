package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
	"strings"
)

func init() {
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsAddApi", &params.GoodsAddParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.UpdateGoodsForFieldApi", &params.UpdateGoodsForFieldParams{})
	params.InitParams("github.com/lvxin0315/gg/handlers.GoodsUpdateApi", &params.GoodsUpdateParams{})
}

func GoodsListApi(c *gin.Context) {
	pagination := new(helper.Pagination)
	pagination.Data = &[]models.Goods{}
	ggList(c, &models.Goods{}, pagination)
}

func GoodsAddApi(c *gin.Context) {
	output := ggOutput(c)
	goodsModel := new(models.Goods)
	goodsUpdateParams := ggParams(c).(*params.GoodsAddParams)
	helper.ReflectiveStructToStructWithJson(goodsModel, goodsUpdateParams)
	goodsModel.SliderImageJson = _sliderImageJson(goodsUpdateParams.SliderImage)
	adminId, _ := services.GetAdminIdByToken(c.Keys["token"].(string))
	goodsModel.AdminId = adminId
	goodsModel.IsShow = helper.Switch2Int(goodsUpdateParams.IsShow)
	goodsModel.IsHot = helper.Switch2Int(goodsUpdateParams.IsHot)
	goodsModel.IsBenefit = helper.Switch2Int(goodsUpdateParams.IsBenefit)
	goodsModel.IsBest = helper.Switch2Int(goodsUpdateParams.IsBest)
	goodsModel.IsNew = helper.Switch2Int(goodsUpdateParams.IsNew)
	goodsModel.IsPostage = helper.Switch2Int(goodsUpdateParams.IsPostage)
	goodsModel.IsGood = helper.Switch2Int(goodsUpdateParams.IsGood)
	err := services.AddGoods(goodsModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = goodsModel
}

//逗号分隔的轮播图地址，变成json
func _sliderImageJson(str string) string {
	strList := strings.Split(str, ",")
	b, err := json.Marshal(strList)
	if err != nil {
		logrus.Error("_sliderImageJson", err)
		return ""
	}
	return string(b)
}

//根据字段修改数据
func UpdateGoodsForFieldApi(c *gin.Context) {
	updateGoodsForFieldParams := ggParams(c).(*params.UpdateGoodsForFieldParams)
	err := services.UpdateGoodsForField(updateGoodsForFieldParams.ID, updateGoodsForFieldParams.Field, updateGoodsForFieldParams.Data)
	if err != nil {
		setGGError(c, err)
		return
	}
}

//商品部分信息编辑，给可文本信息编辑内容
func GoodsUpdateApi(c *gin.Context) {
	output := ggOutput(c)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	goodsModel := new(models.Goods)
	goodsModel.ID = id
	goodsUpdateParams := ggParams(c).(*params.GoodsUpdateParams)
	helper.ReflectiveStructToStructWithJson(goodsModel, goodsUpdateParams)
	err := services.UpdateOne(goodsModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = goodsModel
}

func GoodsGetApi(c *gin.Context) {
	output := ggOutput(c)
	id := helper.String2Uint(c.Param("id"))
	if id <= 0 {
		setGGError(c, fmt.Errorf("参数异常"))
		return
	}
	goodsModel := new(models.Goods)
	err := services.GetOne(goodsModel, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		setGGError(c, err)
	}
	output.Data = goodsModel
}
