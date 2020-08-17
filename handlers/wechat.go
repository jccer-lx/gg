package handlers

import (
	"fmt"
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

var oa *officialaccount.OfficialAccount

const (
	WxOptionA    = "A"
	WxOptionB    = "B"
	WxOptionC    = "C"
	WxOptionD    = "D"
	WxOptionE    = "E"
	WxJudgeTrue  = "对"
	WxJudgeFalse = "错"
	WxBegin      = "wx_begin"      //开始答题
	WxMyScore    = "wx_my_score"   //我的战绩
	WxCorrection = "wx_correction" //题目纠错
)

//返回消息格式
const (
	AnswerErrMessageTpl = `答案：%s
解析：%s`
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
	//记录openid
	_, _ = services.SaveOpenid(string(msg.FromUserName))
	reply := wxMessage.Reply{
		MsgType: wxMessage.MsgTypeText,
		MsgData: wxMessage.NewText("暂时不能处理"),
	}

	switch msg.MsgType {
	case wxMessage.MsgTypeText: //文本类型
		logrus.Debug("MsgTypeText")
		msgTypeText(msg, &reply)
	case wxMessage.MsgTypeEvent: //事件类型
		logrus.Debug("MsgTypeEvent")
		msgTypeEvent(msg, &reply)
	}
	logrus.Debug("reply.MsgType:", reply.MsgType)
	logrus.Debug("reply.MsgData:", reply.MsgData)
	return &reply
}

//初始化自定义菜单
func InitMenu() {
	officialAccount := officialAccount()
	m := officialAccount.GetMenu()
	//我的
	myBtn := new(menu.Button)
	//充值
	myBtn.SetSubButton("我的", []*menu.Button{
		//个人信息
		{
			Name: "个人信息",
			URL:  "http://www.baidu.com",
		},
		{
			Name: "交易信息",
			URL:  "http://www.baidu.com",
		},
		{
			Name: "余额查询",
			Key:  "money",
		},
		{
			Name: "充值",
			URL:  "http://www.baidu.com",
		},
	})

	//付款
	payBtn := new(menu.Button)
	payBtn.Name = "付款"
	payBtn.URL = "http://www.baidu.com"

	//系统
	systemBtn := new(menu.Button)
	//关于我们
	aboutBtn := new(menu.Button)
	aboutBtn.Name = "关于我们"
	aboutBtn.URL = "http://www.baidu.com"
	//商家绑定
	storeBtn := new(menu.Button)
	storeBtn.SetScanCodePushButton("商家绑定", "store")

	//保存到菜单
	err := m.SetMenu([]*menu.Button{
		myBtn,
		payBtn,
		systemBtn,
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
			//答题内容判断
			checkQuestion(msg, reply, item)
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
		checkQuestion(msg, reply, msg.EventKey)
	case WxBegin: //开始答题
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = wxMessage.NewText("开始答题")

	case WxMyScore: //我的战绩
		reply.MsgType = wxMessage.MsgTypeText
		reply.MsgData = wxMessage.NewText("我的战绩")

	case WxCorrection: //题目纠错
		correctionQuestion(msg, reply)
	}
}

//发送客服文本消息
func sendManagerTextMessage(openid string, text string) error {
	manager := wxMessage.NewMessageManager(officialAccount().GetContext())
	//构造文本消息
	customerMessage := wxMessage.NewCustomerTextMessage(openid, text)
	return manager.Send(customerMessage)
}

//发送下一题题目信息
func sendNextQuestionContent(reply *wxMessage.Reply, wxQuestionService *services.WxQuestionService) {
	nextQuestionStr, err := wxQuestionService.NextQuestion()
	if err != nil {
		logrus.Error("sendQuestionContent error:", err)
	}
	reply.MsgType = wxMessage.MsgTypeText
	reply.MsgData = wxMessage.NewText(nextQuestionStr)
}

/*
判断答题情况
@param string item 回答的答案
*/
func checkQuestion(msg wxMessage.MixMessage, reply *wxMessage.Reply, item string) {
	wxQuestionService := services.NewWxQuestionService(string(msg.FromUserName))
	res, questionBank, err := wxQuestionService.Answer(item)
	if !res {
		//第一种情况：报错。显示日志
		if err != nil && err != services.AnswerNotStartedErr {
			logrus.Error("services.Answer error:", err)
		}
		//第二种情况：未开始。什么都不干
		//if err == services.AnswerNotStartedErr {
		//
		//}
		//第三种情况：就是普通的回答错误。通过客服消息，给用户返回解析
		if err == nil {
			err = sendManagerTextMessage(string(msg.FromUserName),
				fmt.Sprintf(AnswerErrMessageTpl,
					questionBank.Question.GetAnswer(),
					questionBank.Question.GetAnalysis()))
			if err != nil {
				logrus.Error("sendManagerTextMessage", err)
			}
		}
	}
	//发送下一题
	sendNextQuestionContent(reply, wxQuestionService)
}

//题目纠错
func correctionQuestion(msg wxMessage.MixMessage, reply *wxMessage.Reply) {
	wxQuestionService := services.NewWxQuestionService(string(msg.FromUserName))
	err := wxQuestionService.Correction()
	if err != nil {
		logrus.Error("wxQuestionService.Correction error:", err)
	}
	//感谢反馈
	err = sendManagerTextMessage(string(msg.FromUserName), "感谢反馈")
	if err != nil {
		logrus.Error("sendManagerTextMessage error:", err)
	}
	//发送下一题
	sendNextQuestionContent(reply, wxQuestionService)
}
