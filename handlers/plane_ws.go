package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
)

var planeWS *helper.GGWebsocket

func init() {
	//监听准备就绪

}

func PlaneWs(c *gin.Context) {
	token := getGGToken(c)
	planeWS.NewClient(c.Writer, c.Request, token, func(ggWS *helper.GGWebsocket, ggWSC *helper.GGWebsocketClient, messageType int, p []byte) {

	})

}

//监听准备情况
func OkListener() {

}
