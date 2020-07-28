package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
)

func GetGoodsCategoryList(GoodsCategoryModel *models.GoodsCategory, pagination *helper.Pagination) error {
	err := GetList(GoodsCategoryModel, pagination)
	if err != nil {
		return err
	}
	//填充p_title
	for _, item := range *pagination.Data.(*[]*models.GoodsCategory) {
		if item.Pid > 0 {
			pGoodsCategoryModel, err := FindGoodsCategoryById(item.ID)
			if err != nil {
				return err
			}
			item.PCateName = pGoodsCategoryModel.CateName
		} else {
			item.PCateName = ""
		}
	}
	return nil
}

func FindGoodsCategoryById(id uint) (*models.GoodsCategory, error) {
	goodsCategoryModel := new(models.GoodsCategory)
	goodsCategoryModel.ID = id
	err := databases.NewDB().Find(goodsCategoryModel).Error
	if err != nil {
		return nil, err
	}
	return goodsCategoryModel, nil
}

func AddGoodsCategory(GoodsCategoryModel *models.GoodsCategory) error {
	//验证pid存在
	if GoodsCategoryModel.Pid > 0 {
		_, err := FindGoodsCategoryById(GoodsCategoryModel.Pid)
		if err != nil {
			return fmt.Errorf("上级分类不存在")
		}
	}
	//填充默认值
	//GoodsCategoryModel.IsShow = models.GoodsCategoryIsShowTrue
	return databases.NewDB().Save(GoodsCategoryModel).Error
}
