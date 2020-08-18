package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"time"
)

//根据openId查询用户信息
func GetUserInfoByOpenid(openid string) (*models.User, error) {
	userInfo := new(models.User)
	err := databases.NewDB().First(userInfo, map[string]interface{}{
		"openid": openid,
	}).Error
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

//保存或更新用户信息
func SaveOrUpdateWechatUserInfo(oauthUserInfo *oauth.UserInfo) (*models.User, error) {
	if oauthUserInfo.OpenID == "" {
		return nil, fmt.Errorf("openid is error")
	}
	oldUserInfo, _ := GetUserInfoByOpenid(oauthUserInfo.OpenID)
	if oldUserInfo == nil || oldUserInfo.Id == 0 {
		//创建用户
		return addUserInfoByWechatInfo(oauthUserInfo)
	}
	//更新用户信息
	oldUserInfo.Nickname = oauthUserInfo.Nickname
	oldUserInfo.Logintime = int(time.Now().Unix())
	err := databases.NewDB().Save(oldUserInfo).Error
	if err != nil {
		return nil, err
	}
	return oldUserInfo, nil
}

//根据微信信息添加会员
func addUserInfoByWechatInfo(oauthUserInfo *oauth.UserInfo) (*models.User, error) {
	userModel := new(models.User)
	userModel.Nickname = oauthUserInfo.Nickname
	userModel.Openid = oauthUserInfo.OpenID
	userModel.Unionid = oauthUserInfo.Unionid
	userModel.Gender = int(oauthUserInfo.Sex)
	//openid -> 用户名
	userModel.Username = oauthUserInfo.OpenID
	//随机密码和salt
	password, salt := randPasswordAndSalt()
	userModel.Salt = salt
	userModel.Password = password
	//邮箱 openid+@gg.com
	userModel.Email = fmt.Sprintf("%s@gg.com", oauthUserInfo.OpenID)
	userModel.Avatar = oauthUserInfo.HeadImgURL
	//fa的默认值
	userModel.Level = 1
	userModel.Score = 0
	userModel.Status = "normal"
	//时间戳
	t := int(time.Now().Unix())
	userModel.Prevtime = t
	userModel.Logintime = t
	err := databases.NewDB().Save(userModel).Error
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

//随机的密码&salt
func randPasswordAndSalt() (password string, salt string) {
	salt = helper.RandString(6)
	password = fmt.Sprintf("%d", time.Now().Unix())
	//参考fa的代码
	password = helper.Md5V(helper.Md5V(password) + salt)
	return password, salt
}
