package services

import (
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/sirupsen/logrus"
	"strings"
)

func checkError(tag string, err error) bool {
	if err != nil {
		logrus.Error(tag, err)
		return false
	}
	return true
}

//带分页列表
func GetList(model interface{}, pagination *helper.Pagination, orderBy ...string) error {
	db := databases.NewDB()
	orderBySql := "id DESC"
	if len(orderBy) > 0 {
		orderBySql = strings.Join(orderBy, ",")
	}
	err := db.Model(model).
		Offset(pagination.Limit * (pagination.Page - 1)).
		Limit(pagination.Limit).
		Order(orderBySql).
		Find(pagination.Data).Error
	if !checkError("GetList", err) {
		return err
	}
	err = db.Model(model).Count(&pagination.Count).Error
	if !checkError("GetList count", err) {
		return err
	}
	return nil
}

func GetOne(model interface{}, where map[string]interface{}) error {
	err := databases.NewDB().Model(model).Find(model, where).Error
	return err
}

func SaveOne(model interface{}) error {
	err := databases.NewDB().Model(model).Save(model).Error
	return err
}

func UpdateOne(model interface{}) error {
	err := databases.NewDB().Model(model).Update(model).Error
	return err
}

//全部列表
func GetAllList(model interface{}, data interface{}) error {
	db := databases.NewDB()
	err := db.Model(model).
		Find(data, map[string]interface{}{
			"pid": 0,
		}).Error
	if !checkError("GetAllList", err) {
		return err
	}
	return nil
}

//删除
func DeleteByIds(model interface{}, ids []uint) error {
	err := databases.NewDB().Where("id IN (?)", ids).Delete(model).Error
	return err
}

//token -> adminId
func GetAdminIdByToken(token string) (uint, error) {
	adminModel := new(models.Admin)
	err := databases.NewDB().Model(adminModel).Find(adminModel, map[string]interface{}{
		"token": token,
	}).Error

	if err != nil {
		return 0, err
	}
	return adminModel.ID, nil
}

//修改某个字段
func SetField(tableName string, id uint, data map[string]interface{}) error {
	db := databases.NewDB()
	err := db.Table(tableName).Where(map[string]interface{}{
		"id": id,
	}).Update(data).Error
	if !checkError("SetField", err) {
		return err
	}
	return nil
}