package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/etc"
	"github.com/silenceper/wechat/v2"
	wxCache "github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	wxConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	wxMessage "github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/sirupsen/logrus"
)

var ot *officialaccount.OfficialAccount

//微信接口
func WeChat(c *gin.Context) {
	officialAccount := officialAccount()
	// 传入request和responseWriter
	wxServer := officialAccount.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	wxServer.SetMessageHandler(wxMessageFunc)
	//处理消息接收以及回复
	err := wxServer.Serve()
	if err != nil {
		logrus.Error("server.Serve error:", err)
		return
	}
	//发送回复的消息
	err = wxServer.Send()
	if err != nil {
		logrus.Error("server.Send error:", err)
		return
	}
}

func officialAccount() *officialaccount.OfficialAccount {
	if ot == nil {
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
		ot = wc.GetOfficialAccount(cfg)
	}
	return ot
}

//消息处理
func wxMessageFunc(msg wxMessage.MixMessage) *wxMessage.Reply {
	reply := new(wxMessage.Reply)

	switch msg.MsgType {
	case wxMessage.MsgTypeText: //文本类型
	case wxMessage.MsgTypeEvent: //事件类型
		switch msg.Event {
		case wxMessage.EventSubscribe: //关注
		case wxMessage.EventUnsubscribe: //取消关注
		case wxMessage.EventClick: //点击
		}
	default:
		//暂未处理类型
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = wxMessage.NewText("暂时不能处理")
	}
	return reply
}

//初始化自定义菜单
func InitMenu() {
	officialAccount := officialAccount()
	m := officialAccount.GetMenu()

	//测试点击
	btn1 := new(menu.Button)
	btn1.SetClickButton("测试1", "test1")
	btn2 := new(menu.Button)
	btn2.SetClickButton("测试2", "test2")
	//测试二级菜单
	btn3 := new(menu.Button)
	btn31 := new(menu.Button)
	btn31.SetClickButton("测试31", "test31")
	btn32 := new(menu.Button)
	btn32.SetClickButton("测试32", "test32")
	btn33 := new(menu.Button)
	btn33.SetClickButton("测试33", "test33")
	btn3.SetSubButton("测试3", []*menu.Button{
		btn31,
		btn32,
		btn33,
	})
	err := m.SetMenu([]*menu.Button{
		btn1,
		btn2,
		btn3,
	})
	if err != nil {
		logrus.Error("m.SetMenu:", err)
	}
}
