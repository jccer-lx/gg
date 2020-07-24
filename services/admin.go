package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
)

//添加管理员
func AddAdmin(adminModel *models.Admin) error {
	//明文密码加密
	p, s := makePassword(adminModel.Password)
	adminModel.Password = p
	adminModel.Salt = s
	//状态
	adminModel.Status = models.AdminStatusNormal
	//username唯一
	hasAdminModel, err := FindAdminByUsername(adminModel.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if hasAdminModel != nil && hasAdminModel.ID > 0 {
		return fmt.Errorf("username 不能重复")
	}
	err = databases.NewDB().Save(adminModel).Error
	return err
}

//生成6位盐对密码加密
func makePassword(str string) (password string, salt string) {
	salt = helper.RandString(6)
	password = helper.Md5V(salt + str)
	return
}

//根据username获取admin
func FindAdminByUsername(username string) (*models.Admin, error) {
	adminModel := new(models.Admin)
	err := databases.NewDB().Model(adminModel).Find(adminModel, map[string]interface{}{
		"username": username,
	}).Error
	if err != nil {
		return adminModel, err
	}
	return adminModel, nil
}

//登录
func Login(username string, password string) (*models.Admin, error) {
	adminModel, err := FindAdminByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("账号不存在")
		}
		return nil, err
	}
	//判断密码
	if adminModel.Password != helper.Md5V(adminModel.Salt+password) {
		return nil, fmt.Errorf("密码错误")
	}
	//状态
	if adminModel.Status != models.AdminStatusNormal {
		return nil, fmt.Errorf("账号被禁用")
	}
	//生成token
	adminModel.Token = helper.SaveToken(adminModel.ID)
	err = databases.NewDB().Model(adminModel).Save(adminModel).Error
	if err != nil {
		return nil, fmt.Errorf("token异常")
	}
	return adminModel, nil
}
