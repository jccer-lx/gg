package services

import (
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/sirupsen/logrus"
)

func checkError(tag string, err error) bool {
	if err != nil {
		logrus.Error(tag, err)
		return false
	}
	return true
}

//带分页列表
func GetList(model interface{}, pagination *helper.Pagination) error {
	db := databases.NewDB()
	err := db.Model(model).
		Offset(pagination.Limit * (pagination.Page - 1)).
		Limit(pagination.Limit).
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
		Find(data).Error
	if !checkError("GetAllList", err) {
		return err
	}
	return nil
}
