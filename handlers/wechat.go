package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/etc"
	"github.com/silenceper/wechat/v2"
	wxCache "github.com/silenceper/wechat/v2/cache"
	wxConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	wxMessage "github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/sirupsen/logrus"
)

//微信接口
func WeChat(c *gin.Context) {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := wxCache.NewMemory()
	cfg := &wxConfig.Config{
		AppID:          etc.Config.Wx.AppID,
		AppSecret:      etc.Config.Wx.AppSecret,
		Token:          etc.Config.Wx.Token,
		EncodingAESKey: etc.Config.Wx.EncodingAESKey,
		Cache:          memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	// 传入request和responseWriter
	server := officialAccount.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg wxMessage.MixMessage) *wxMessage.Reply {
		//回复消息：演示回复用户发送的消息
		text := wxMessage.NewText(msg.Content)
		return &wxMessage.Reply{MsgType: wxMessage.MsgTypeText, MsgData: text}
	})
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		logrus.Error("server.Serve error:", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		logrus.Error("server.Send error:", err)
		return
	}
}
