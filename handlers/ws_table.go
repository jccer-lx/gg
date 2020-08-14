package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
)

type wstMessage struct {
	Ctrl       int    `json:"ctrl"`
	EventKey   string `json:"event_key"`
	RowNum     int    `json:"row_num"`
	LatestData string `json:"latest_data"`
	DataId     uint   `json:"data_id"`
	MessageId  int    `json:"message_id"`
}

const (
	wstMessageNoCtrl         = 0
	wstMessageLockColCtrl    = 1
	wstMessageUnLockColCtrl  = 2
	wstMessageChangeDataCtrl = 3
)

var ggWS *helper.GGWebsocket
var messageId = 0

func init() {
	ggWS = helper.NewGGWebsocket()
}

//ws连接
func WsTable(c *gin.Context) {
	token := getGGToken(c)
	err := ggWS.NewClient(c.Writer, c.Request, token, readMessageCallback)
	if err != nil {
		fmt.Println("ggWS error: ", err)
		return
	}
}

//接受消息回调
func readMessageCallback(ggWS *helper.GGWebsocket, ggWSC *helper.GGWebsocketClient, messageType int, p []byte) {
	messageId++
	wstM := new(wstMessage)
	wstM.MessageId = messageId
	err := json.Unmarshal(p, wstM)
	if err != nil {
		logrus.Error("readMessageCallback error : ", err)
		return
	}
	wstMB, err := json.Marshal(wstM)
	if err != nil {
		logrus.Error("wstMB error : ", err)
		return
	}
	wsTableDataModel := new(models.WsTableData)
	switch wstM.Ctrl {
	case wstMessageLockColCtrl, wstMessageUnLockColCtrl:
		//开始编辑 //提交编辑
		go ggWS.BroadcastByteWithoutSender(ggWSC.GetClientId(), wstMB)
	case wstMessageChangeDataCtrl:
		//修改信息
		wsTableDataModel.ID = wstM.DataId
		err := services.SetField(wsTableDataModel.TableName(), wstM.DataId, map[string]interface{}{
			wstM.EventKey: wstM.LatestData,
		})
		if err != nil {
			logrus.Error("wstMessageChangeDataCtrl", err)
			return
		}
		//广播
		go ggWS.BroadcastByteWithoutSender(ggWSC.GetClientId(), wstMB)
	default:
	}
}
