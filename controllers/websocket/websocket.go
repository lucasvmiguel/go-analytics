package websocket

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lucasvmiguel/go-analytics/auth"
	"github.com/lucasvmiguel/go-analytics/errors"
	"github.com/lucasvmiguel/go-analytics/model"
)

type WebSocketClient struct {
	ID      string
	MsgType int
	Conn    *websocket.Conn
}

var clients []WebSocketClient

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error(errors.ErrUpgradeWebsocket.Error())
		return
	}

	usr, msgType, err := auth.Websocket(conn)

	if err != nil {
		logrus.Error(errors.ErrAuthWebsocket.Error())
	}

	clients = append(clients, WebSocketClient{usr, msgType, conn})
}

func WebsocketController(c *gin.Context) {
	clients = make([]WebSocketClient, 0)
	wshandler(c.Writer, c.Request)
}

func Send(noti model.Notification) {
	for _, client := range clients {
		noti.Time = time.Now()
		notiJSON, err := json.Marshal(noti)
		if err != nil {
			logrus.Error(errors.ErrSerializeNotification.Error())
			continue
		}
		client.Conn.WriteMessage(client.MsgType, notiJSON)
	}
}
