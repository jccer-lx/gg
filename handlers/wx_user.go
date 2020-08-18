package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
	"net/http"
)

//个人信息页面
func UserInfoView(c *gin.Context) {
	//c.HTML(http.StatusOK, "user/user_info.tpl", nil)
	//return

	userModel, err := saveWechatUserInfo(c)
	if err != nil {
		logrus.Error("saveWechatUserInfo：", err)
		ErrorPage(c, err)
		return
	}
	//处理用户信息页面
	userPage(c, userModel)
}

//处理微信用户信息
func saveWechatUserInfo(c *gin.Context) (*models.User, error) {
	code := c.Query("code")
	if code == "" {
		return nil, fmt.Errorf("code error")
	}
	logrus.Info("微信的code：", code)
	officialAccount := officialAccount()
	resAt, err := officialAccount.GetOauth().GetUserAccessToken(code)
	if err != nil {
		return nil, err
	}
	logrus.Info("微信的openid：", resAt.OpenID)
	//获取用户全信息
	userInfo, err := officialAccount.GetOauth().GetUserInfo(resAt.AccessToken, resAt.OpenID)
	if err != nil {
		return nil, err
	}
	logrus.Info(userInfo)
	//保存&更新
	userModel, err := services.SaveOrUpdateWechatUserInfo(&userInfo)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

//重定向到我的信息
func userPage(c *gin.Context, userModel *models.User) {
	c.Redirect(http.StatusPermanentRedirect, "/v/user/openid/"+userModel.Openid)
}
