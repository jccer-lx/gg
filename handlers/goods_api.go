package handlers

import (
	"encoding/json"
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
	goodsAddParams := ggParams(c).(*params.GoodsAddParams)
	goodsModel.Name = goodsAddParams.Name
	goodsModel.MainInfo = goodsAddParams.MainInfo
	goodsModel.MainImage = goodsAddParams.MainImage
	goodsModel.SliderImageJson = _sliderImageJson(goodsAddParams.SliderImage)
	adminId, _ := services.GetAdminIdByToken(c.Keys["token"].(string))
	goodsModel.AdminId = adminId
	goodsModel.Keyword = goodsAddParams.Keyword
	goodsModel.BarCode = goodsAddParams.BarCode
	goodsModel.CategoryId = helper.JsonNumber2Uint(goodsAddParams.CategoryId)
	goodsModel.Price = helper.JsonNumber2Float64(goodsAddParams.Price)
	goodsModel.VipPrice = helper.JsonNumber2Float64(goodsAddParams.VipPrice)
	goodsModel.OtPrice = helper.JsonNumber2Float64(goodsAddParams.OtPrice)
	goodsModel.Postage = helper.JsonNumber2Float64(goodsAddParams.Postage)
	goodsModel.UnitName = goodsAddParams.UnitName
	goodsModel.Sort = helper.JsonNumber2Int(goodsAddParams.Sort)
	goodsModel.Sales = helper.JsonNumber2Int(goodsAddParams.Sales)
	goodsModel.Stock = helper.JsonNumber2Int(goodsAddParams.Stock)
	goodsModel.IsShow = helper.Switch2Int(goodsAddParams.IsShow)
	goodsModel.IsHot = helper.Switch2Int(goodsAddParams.IsHot)
	goodsModel.IsBenefit = helper.Switch2Int(goodsAddParams.IsBenefit)
	goodsModel.IsBest = helper.Switch2Int(goodsAddParams.IsBest)
	goodsModel.IsNew = helper.Switch2Int(goodsAddParams.IsNew)
	goodsModel.IsPostage = helper.Switch2Int(goodsAddParams.IsPostage)
	goodsModel.GiveIntegral = helper.JsonNumber2Int(goodsAddParams.GiveIntegral)
	goodsModel.Cost = helper.JsonNumber2Float64(goodsAddParams.Cost)
	goodsModel.IsGood = helper.Switch2Int(goodsAddParams.IsGood)
	goodsModel.VirtualSales = helper.JsonNumber2Int(goodsAddParams.VirtualSales)
	goodsModel.Browse = helper.JsonNumber2Int(goodsAddParams.Browse)
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
