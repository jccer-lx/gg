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

func GetList(model interface{}, pagination *helper.Pagination) error {
	db := databases.NewDB()
	err := db.Model(model).
		Offset(pagination.Limit * (pagination.Page - 1)).
		Limit(pagination.Limit).
		Find(pagination.Data).Error
	if !checkError("GetAdminList", err) {
		return err
	}
	err = db.Model(model).Count(&pagination.Count).Error
	if !checkError("GetAdminList count", err) {
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
