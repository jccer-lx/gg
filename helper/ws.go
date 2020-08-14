package helper

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//连接容器
type GGWebsocket struct {
	clientList     []*GGWebsocketClient
	clientConnList []*websocket.Conn
	messageId      int64
	wsUpGrader     *websocket.Upgrader
}

//ws客户端
type GGWebsocketClient struct {
	//客户端标识
	clientId string
	//ws conn
	conn *websocket.Conn
	//接收消息的callback
	readMessageCallback func(ggWS *GGWebsocket, ggWSC *GGWebsocketClient, messageType int, p []byte)
	//发消息channel
	writeMessageChan chan *ggSysMessage
	//ws容器信息
	ggWS *GGWebsocket
}

type ReadMessageCallbackFunc = func(ggWS *GGWebsocket, ggWSC *GGWebsocketClient, messageType int, p []byte)

//群发消息体
type ggMessage struct {
	MessageType int         `json:"-"`
	MessageId   int64       `json:"message_id"`
	Data        interface{} `json:"data"`
	ClientId    string      `json:"client_id"`
	IsSystem    bool        `json:"is_system"`
	SendTime    time.Time   `json:"send_time"`
}

//内部消息体
type ggSysMessage struct {
	MessageType int
	Data        []byte
}

//客户端连接
func (ggWSC *GGWebsocketClient) run() error {
	//写
	go ggWSC.writeListener()
	//读
	for {
		t, reply, err := ggWSC.conn.ReadMessage()
		if err != nil {
			return err
		}
		ggWSC.readMessageCallback(ggWSC.ggWS, ggWSC, t, reply)
	}
}

//发送消息信道
func (ggWSC *GGWebsocketClient) writeListener() {
	ggWSC.writeMessageChan = make(chan *ggSysMessage)
	for {
		select {
		case bm := <-ggWSC.writeMessageChan:
			ggWSC.conn.WriteMessage(bm.MessageType, bm.Data)
		}
	}
}

//发送消息
func (ggWSC *GGWebsocketClient) writeGGMessage(m *ggMessage) error {
	sysM := new(ggSysMessage)
	sysM.MessageType = m.MessageType
	//json
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	sysM.Data = b
	ggWSC.writeMessageChan <- sysM
	return nil
}

func (ggWSC *GGWebsocketClient) writeGGSysMessage(m *ggSysMessage) error {
	ggWSC.writeMessageChan <- m
	return nil
}

//获取clientId
func (ggWSC *GGWebsocketClient) GetClientId() string {
	return ggWSC.clientId
}

//添加客户端(连接)
func (ggWS *GGWebsocket) NewClient(w http.ResponseWriter, r *http.Request, clientId string, readMessageCallback ReadMessageCallbackFunc) error {
	conn, err := ggWS.wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	ggWSC := new(GGWebsocketClient)
	ggWSC.clientId = clientId
	ggWSC.conn = conn
	ggWSC.readMessageCallback = readMessageCallback
	ggWSC.ggWS = ggWS
	//添加到集合中
	ggWS.clientConnList = append(ggWS.clientConnList, conn)
	ggWS.clientList = append(ggWS.clientList, ggWSC)
	defer ggWS.DeleteClient(ggWSC)
	//连接监听
	return ggWSC.run()
}

//删除客户端
func (ggWS *GGWebsocket) DeleteClient(ggWSC *GGWebsocketClient) {
	//关闭客户端conn
	defer ggWSC.conn.Close()
	//删除客户端结构体
	for i, c := range ggWS.clientList {
		if c.clientId == ggWSC.clientId {
			ggWS.clientList = append(ggWS.clientList[:i], ggWS.clientList[:i]...)
			break
		}
	}
	//删除客户端conn
	for i, conn := range ggWS.clientConnList {
		if conn == ggWSC.conn {
			ggWS.clientConnList = append(ggWS.clientConnList[:i], ggWS.clientConnList[:i]...)
			break
		}
	}
	return
}

//广播
func (ggWS *GGWebsocket) BroadcastByte(message []byte) {
	//消息序号
	ggWS.messageId++
	m := new(ggSysMessage)
	m.Data = message
	m.MessageType = websocket.TextMessage
	for _, client := range ggWS.clientList {
		err := client.writeGGSysMessage(m)
		if err != nil {
			logrus.Errorf("BroadcastByte error, clientId:%s, messageId: %d, error: %v",
				client.clientId, ggWS.messageId, err)
		}
	}
}

//广播interface{}
func (ggWS *GGWebsocket) BroadcastInterface(messageData interface{}) error {
	//消息序号
	ggWS.messageId++
	m := new(ggMessage)
	m.MessageType = websocket.TextMessage
	m.MessageId = ggWS.messageId
	m.IsSystem = false
	m.Data = messageData
	m.SendTime = time.Now()
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	ggWS.BroadcastByte(b)
	return nil
}

//广播（排除发送者）
func (ggWS *GGWebsocket) BroadcastByteWithoutSender(clientId string, message []byte) {
	//消息序号
	ggWS.messageId++
	m := new(ggSysMessage)
	m.Data = message
	m.MessageType = websocket.TextMessage
	for _, client := range ggWS.clientList {
		if clientId == client.clientId {
			continue
		}
		err := client.writeGGSysMessage(m)
		if err != nil {
			logrus.Errorf("BroadcastByte error, clientId:%s, messageId: %d, error: %v",
				client.clientId, ggWS.messageId, err)
		}
	}
}

//广播interface{}（排除发送者）
func (ggWS *GGWebsocket) BroadcastInterfaceWithoutSender(clientId string, messageData interface{}) error {
	//消息序号
	ggWS.messageId++
	m := new(ggMessage)
	m.MessageType = websocket.TextMessage
	m.MessageId = ggWS.messageId
	m.IsSystem = false
	m.Data = messageData
	m.SendTime = time.Now()
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	ggWS.BroadcastByteWithoutSender(clientId, b)
	return nil
}

//初始化ws
func NewGGWebsocket() *GGWebsocket {
	ggWS := new(GGWebsocket)
	ggWS.wsUpGrader = &websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		// 取消ws跨域校验
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return ggWS
}

//启动ws监听

//初始化ws并启动ws监听
