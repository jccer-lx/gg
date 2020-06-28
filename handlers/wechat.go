package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/etc"
	"github.com/lvxin0315/gg/services"
	"github.com/silenceper/wechat/v2"
	wxCache "github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	wxConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	wxMessage "github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/sirupsen/logrus"
)

var ot *officialaccount.OfficialAccount

const (
	WxOptionA    = "A"
	WxOptionB    = "B"
	WxOptionC    = "C"
	WxOptionD    = "D"
	WxOptionE    = "E"
	WxJudgeTrue  = "对"
	WxJudgeFalse = "错"
	WxBegin      = "wx_begin"    //开始答题
	WxMyScore    = "wx_my_score" //我的战绩
)

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
	reply.MsgType = wxMessage.MsgTypeText
	reply.MsgData = wxMessage.NewText("暂时不能处理")

	switch msg.MsgType {
	case wxMessage.MsgTypeText: //文本类型
		msgTypeText(msg, reply)
	case wxMessage.MsgTypeEvent: //事件类型
		msgTypeEvent(msg, reply)
	}
	logrus.Debug("reply.MsgType:", reply.MsgType)
	logrus.Debug("reply.MsgData:", reply.MsgData)
	return reply
}

//初始化自定义菜单
func InitMenu() {
	officialAccount := officialAccount()
	m := officialAccount.GetMenu()
	//选择题快捷键
	choiceBtn := new(menu.Button)
	choiceABtn := new(menu.Button)
	choiceBBtn := new(menu.Button)
	choiceCBtn := new(menu.Button)
	choiceDBtn := new(menu.Button)
	choiceEBtn := new(menu.Button)
	choiceBtn.SetSubButton("选择键", []*menu.Button{
		choiceABtn,
		choiceBBtn,
		choiceCBtn,
		choiceDBtn,
		choiceEBtn,
	})
	choiceABtn.SetClickButton("选项A", WxOptionA)
	choiceBBtn.SetClickButton("选项B", WxOptionB)
	choiceCBtn.SetClickButton("选项C", WxOptionC)
	choiceDBtn.SetClickButton("选项D", WxOptionD)
	choiceEBtn.SetClickButton("选项E", WxOptionE)
	//判断题快捷键
	judgeBtn := new(menu.Button)
	judgeTrueBtn := new(menu.Button)
	judgeFalseBtn := new(menu.Button)
	judgeBtn.SetSubButton("判断键", []*menu.Button{
		judgeTrueBtn,
		judgeFalseBtn,
	})
	judgeTrueBtn.SetClickButton("对", WxJudgeTrue)
	judgeFalseBtn.SetClickButton("错", WxJudgeFalse)
	//其他
	otherBtn := new(menu.Button)
	//开始答题
	beginBtn := new(menu.Button)
	beginBtn.SetClickButton("答题", WxBegin)
	//我的战绩
	myScore := new(menu.Button)
	myScore.SetClickButton("我的战绩", WxMyScore)
	otherBtn.SetSubButton("我的", []*menu.Button{
		beginBtn,
		myScore,
	})
	//保存到菜单
	err := m.SetMenu([]*menu.Button{
		choiceBtn,
		judgeBtn,
		otherBtn,
	})
	if err != nil {
		logrus.Error("m.SetMenu:", err)
	}
}

//文本消息处理
func msgTypeText(msg wxMessage.MixMessage, reply *wxMessage.Reply) {
	//判断是不是答题内容
	for _, item := range []string{WxOptionA, WxOptionB, WxOptionC, WxOptionD, WxOptionE, WxJudgeTrue, WxJudgeFalse} {
		if msg.Content == item {
			//答题内容
			resContent, err := services.Answer(msg.OpenID, item)
			if err != nil {
				logrus.Error("services.Answer error:", err)
				return
			}
			reply.MsgType = wxMessage.MsgTypeText
			reply.MsgData = resContent
			return
		}
	}
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
	case WxOptionA, WxOptionB, WxOptionC, WxOptionD, WxOptionE, WxJudgeTrue, WxJudgeFalse:
		//答题状态
		resContent, err := services.Answer(msg.OpenID, msg.EventKey)
		if err != nil {
			logrus.Error("services.Answer error:", err)
			return
		}
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = resContent
	case WxBegin: //开始答题
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = "开始答题"

	case WxMyScore: //我的战绩
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = "我的战绩"
	}
}
