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
	output := ggOutput(c)
	goodsModel := new(models.Goods)
	//分页参数
	pagination := new(helper.Pagination)
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	var goodsList []*models.Goods
	pagination.Data = &goodsList
	err = services.GetList(goodsModel, pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = goodsList
	output.Count = pagination.Count
}

func GoodsAddApi(c *gin.Context) {
	output := ggOutput(c)
	goodsModel := new(models.Goods)
	goodsUpdateParams := ggParams(c).(*params.GoodsAddParams)
	goodsModel.Name = goodsUpdateParams.Name
	goodsModel.MainInfo = goodsUpdateParams.MainInfo
	goodsModel.MainImage = goodsUpdateParams.MainImage
	goodsModel.SliderImageJson = _sliderImageJson(goodsUpdateParams.SliderImage)
	adminId, _ := services.GetAdminIdByToken(c.Keys["token"].(string))
	goodsModel.AdminId = adminId
	goodsModel.Keyword = goodsUpdateParams.Keyword
	goodsModel.BarCode = goodsUpdateParams.BarCode
	goodsModel.CategoryId = helper.JsonNumber2Uint(goodsUpdateParams.CategoryId)
	goodsModel.Price = helper.JsonNumber2Float64(goodsUpdateParams.Price)
	goodsModel.VipPrice = helper.JsonNumber2Float64(goodsUpdateParams.VipPrice)
	goodsModel.OtPrice = helper.JsonNumber2Float64(goodsUpdateParams.OtPrice)
	goodsModel.Postage = helper.JsonNumber2Float64(goodsUpdateParams.Postage)
	goodsModel.UnitName = goodsUpdateParams.UnitName
	goodsModel.Sort = helper.JsonNumber2Int(goodsUpdateParams.Sort)
	goodsModel.Sales = helper.JsonNumber2Int(goodsUpdateParams.Sales)
	goodsModel.Stock = helper.JsonNumber2Int(goodsUpdateParams.Stock)
	goodsModel.IsShow = helper.Switch2Int(goodsUpdateParams.IsShow)
	goodsModel.IsHot = helper.Switch2Int(goodsUpdateParams.IsHot)
	goodsModel.IsBenefit = helper.Switch2Int(goodsUpdateParams.IsBenefit)
	goodsModel.IsBest = helper.Switch2Int(goodsUpdateParams.IsBest)
	goodsModel.IsNew = helper.Switch2Int(goodsUpdateParams.IsNew)
	goodsModel.IsPostage = helper.Switch2Int(goodsUpdateParams.IsPostage)
	goodsModel.GiveIntegral = helper.JsonNumber2Int(goodsUpdateParams.GiveIntegral)
	goodsModel.Cost = helper.JsonNumber2Float64(goodsUpdateParams.Cost)
	goodsModel.IsGood = helper.Switch2Int(goodsUpdateParams.IsGood)
	goodsModel.VirtualSales = helper.JsonNumber2Int(goodsUpdateParams.VirtualSales)
	goodsModel.Browse = helper.JsonNumber2Int(goodsUpdateParams.Browse)
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
	goodsModel.Name = goodsUpdateParams.Name
	goodsModel.Keyword = goodsUpdateParams.Keyword
	goodsModel.BarCode = goodsUpdateParams.BarCode
	goodsModel.Price = helper.JsonNumber2Float64(goodsUpdateParams.Price)
	goodsModel.VipPrice = helper.JsonNumber2Float64(goodsUpdateParams.VipPrice)
	goodsModel.OtPrice = helper.JsonNumber2Float64(goodsUpdateParams.OtPrice)
	goodsModel.Postage = helper.JsonNumber2Float64(goodsUpdateParams.Postage)
	goodsModel.UnitName = goodsUpdateParams.UnitName
	goodsModel.Sort = helper.JsonNumber2Int(goodsUpdateParams.Sort)
	goodsModel.Sales = helper.JsonNumber2Int(goodsUpdateParams.Sales)
	goodsModel.Stock = helper.JsonNumber2Int(goodsUpdateParams.Stock)
	goodsModel.GiveIntegral = helper.JsonNumber2Int(goodsUpdateParams.GiveIntegral)
	goodsModel.Cost = helper.JsonNumber2Float64(goodsUpdateParams.Cost)
	goodsModel.VirtualSales = helper.JsonNumber2Int(goodsUpdateParams.VirtualSales)
	goodsModel.Browse = helper.JsonNumber2Int(goodsUpdateParams.Browse)
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
