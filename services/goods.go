package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
)

func AddGoods(goods *models.Goods) error {
	//验证分类id存在
	_, err := FindGoodsCategoryById(goods.CategoryId)
	if err != nil {
		return fmt.Errorf("无效分类")
	}
	//验证BarCode，必须唯一，若无则忽略
	if goods.BarCode != "" {
		barCodeGoodsModel := new(models.Goods)
		databases.NewDB().Where(map[string]interface{}{
			"bar_code": goods.BarCode,
		}).First(barCodeGoodsModel)
		if barCodeGoodsModel.ID > 0 {
			return fmt.Errorf("bar_code已经存在")
		}
	}
	return databases.NewDB().Save(goods).Error
}

//更新指定字段
func UpdateGoodsForField(id uint, field string, data interface{}) error {
	goodsModel := new(models.Goods)
	goodsModel.ID = id
	err := databases.NewDB().Model(goodsModel).Where("id = ?", id).Update(field, data).Error
	return err
}
