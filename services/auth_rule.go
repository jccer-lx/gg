package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
)

func AddAuthRule(authRuleModel *models.AuthRule) error {
	//验证name唯一
	oldAuthRuleModel, _ := FindAuthRuleByName(authRuleModel.Name)
	if oldAuthRuleModel.ID > 0 {
		return fmt.Errorf("「值」已经存在")
	}
	//验证pid存在
	if authRuleModel.Pid > 0 {
		_, err := FindAuthRuleById(authRuleModel.Pid)
		if err != nil {
			return fmt.Errorf("上级菜单不存在")
		}
	}
	//填充默认值
	authRuleModel.Status = models.AuthRuleStatusNormal
	authRuleModel.Type = models.AuthRuleTypeFile
	authRuleModel.IsMenu = models.AuthRuleIsMenuTrue
	return databases.NewDB().Save(authRuleModel).Error
}

func FindAuthRuleById(id uint) (authRule *models.AuthRule, err error) {
	authRule = new(models.AuthRule)
	authRule.ID = id
	err = databases.NewDB().Find(authRule).Error
	return
}

func FindAuthRuleByName(name string) (authRule *models.AuthRule, err error) {
	authRule = new(models.AuthRule)
	err = databases.NewDB().Find(authRule, map[string]interface{}{
		"name": name,
	}).Error
	return
}

func GetAuthRuleList(authRuleModel *models.AuthRule, pagination *helper.Pagination) error {
	err := GetList(authRuleModel, pagination)
	if err != nil {
		return err
	}
	//填充p_title
	for _, item := range *pagination.Data.(*[]*models.AuthRule) {
		if item.Pid > 0 {
			pAuthRuleModel, err := FindAuthRuleById(item.ID)
			if err != nil {
				return err
			}
			item.PTitle = pAuthRuleModel.Title
		} else {
			item.PTitle = ""
		}
	}
	return nil
}
