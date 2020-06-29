package services

import (
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
)

/*
通过微信openid，查询或新增用户
@param string openid
@return
*/
func SaveOpenid(openid string) (userModel *models.User, err error) {
	userModel = new(models.User)
	userModel.Openid = openid
	err = databases.NewDB().FirstOrCreate(userModel, map[string]interface{}{
		"openid": openid,
	}).Error
	return userModel, err
}
