package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/services"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var clientConnList = make(map[string]*websocket.Conn)
var sendAllClientMessageChan = make(chan *sendAllClientMessage)

//消息序号
var messageNo = 0

type sendAllClientMessage struct {
	Sender    string
	WsMessage []byte
}

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

var wsUpGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	go func() {
		for {
			select {
			case st := <-sendAllClientMessageChan:
				for sr, conn := range clientConnList {
					if sr == st.Sender {
						continue
					}
					_ = conn.WriteMessage(websocket.TextMessage, st.WsMessage)
				}
			}
		}
	}()
}

func WsTable(c *gin.Context) {
	conn, err := wsUpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}
	u := fmt.Sprintf("%d", time.Now().UnixNano())
	clientConnList[u] = conn
	defer conn.Close()
	//监听read message
	for {
		t, reply, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			logrus.Info("message type != 1")
			continue
		}
		//fmt.Println(t)
		//fmt.Println(string(reply))
		//json -> struct
		wstM := new(wstMessage)
		err = json.Unmarshal(reply, wstM)
		if err != nil {
			logrus.Error("json.Unmarshal(reply, wstM) ", err)
			continue
		}
		go wsTableMessage(u, wstM)
	}
	delete(clientConnList, u)

}

//消息处理
func wsTableMessage(sender string, wstM *wstMessage) {
	wsTableDataModel := new(models.WsTableData)
	switch wstM.Ctrl {
	case wstMessageLockColCtrl, wstMessageUnLockColCtrl:
		//开始编辑 //提交编辑
		go sendAllClientWithoutSender(sender, wstM)
	case wstMessageChangeDataCtrl:
		//修改信息
		wsTableDataModel.ID = wstM.DataId
		err := services.SetField(wsTableDataModel.TableName(), wstM.DataId, map[string]interface{}{
			wstM.EventKey: wstM.LatestData,
		})
		if err != nil {
			logrus.Error("wstMessageChangeDataCtrl", err)
		}
		//广播
		go sendAllClientWithoutSender(sender, wstM)
	default:
	}

}

//广播内容(排除发送者)
func sendAllClientWithoutSender(sender string, wstM *wstMessage) {
	//消息序号
	messageNo++
	wstM.MessageId = messageNo
	wstMByte, err := json.Marshal(wstM)
	if err != nil {
		return
	}
	sendAllClientMessageChan <- &sendAllClientMessage{
		Sender:    sender,
		WsMessage: wstMByte,
	}
}
