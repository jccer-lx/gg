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

var oa *officialaccount.OfficialAccount

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
	if oa == nil {
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
		oa = wc.GetOfficialAccount(cfg)
	}

	return oa
}

//消息处理
func wxMessageFunc(msg wxMessage.MixMessage) *wxMessage.Reply {

	switch msg.MsgType {
	case wxMessage.MsgTypeText: //文本类型

	case wxMessage.MsgTypeEvent: //事件类型

	}

	return nil
}

//初始化自定义菜单
func InitMenu() {
	officialAccount := officialAccount()
	//url处理
	ggOauth := officialAccount.GetOauth()
	hostName := etc.Config.Host
	userInfoUrl, err := ggOauth.GetRedirectURL(hostName+"/v/wx/user_info", "snsapi_userinfo", "")
	checkMenuError(err)
	//菜单设置
	ggMenu := officialAccount.GetMenu()
	//1.我的
	myBtn := new(menu.Button)
	//1-1.个人信息
	userInfoBtn := new(menu.Button)
	userInfoBtn.SetViewButton("个人信息", userInfoUrl)
	//1-2.交易信息
	transactionBtn := new(menu.Button)
	transactionBtn.SetViewButton("交易信息", "http://www.baidu.com")
	//1-3.余额查询
	moneyBtn := new(menu.Button)
	moneyBtn.SetClickButton("余额查询", "money")
	//1-4.充值
	rechargeBtn := new(menu.Button)
	rechargeBtn.SetViewButton("充值", "http://www.baidu.com")
	myBtn.SetSubButton("我的", []*menu.Button{
		userInfoBtn,
		transactionBtn,
		moneyBtn,
		rechargeBtn,
	})
	//2.付款
	payBtn := new(menu.Button)
	payBtn.SetViewButton("付款", "http://www.baidu.com")
	//3.系统
	systemBtn := new(menu.Button)
	//3-1.关于我们
	aboutBtn := new(menu.Button)
	aboutBtn.SetViewButton("关于我们", "http://www.baidu.com")
	//3-2.商家绑定
	storeBtn := new(menu.Button)
	storeBtn.SetScanCodeWaitMsgButton("商家绑定", "store")
	//3-3.商家收款
	collectBtn := new(menu.Button)
	collectBtn.SetScanCodeWaitMsgButton("商家收款", "collect")
	systemBtn.SetSubButton("系统", []*menu.Button{
		aboutBtn,
		storeBtn,
		collectBtn,
	})
	//保存到菜单
	err = ggMenu.SetMenu([]*menu.Button{
		myBtn,
		payBtn,
		systemBtn,
	})
	checkMenuError(err)
}

//文本消息处理
func msgTypeText(msg wxMessage.MixMessage, reply *wxMessage.Reply) {

}

//事件处理
func msgTypeEvent(msg wxMessage.MixMessage, reply *wxMessage.Reply) {
	switch msg.Event {
	case wxMessage.EventSubscribe: //关注
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = wxMessage.NewText("感谢关注")
	case wxMessage.EventUnsubscribe: //取消关注
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = wxMessage.NewText("后会有期")
	case wxMessage.EventClick: //点击
		eventClick(msg, reply)
	}
}

//菜单点击事件
func eventClick(msg wxMessage.MixMessage, reply *wxMessage.Reply) {
	switch msg.EventKey {

	}
}

func checkMenuError(err error) {
	if err != nil {
		logrus.Error("checkMenuError:", err)
		panic(err)
	}
}
