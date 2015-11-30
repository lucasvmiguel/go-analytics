package websocket

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lucasvmiguel/go-analytics/auth"
	"github.com/lucasvmiguel/go-analytics/model"
)

type WebSocketClient struct {
	ID   string
	Conn *websocket.Conn
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
		logrus.Error("Failed to set websocket upgrade: " + err.Error())
		return
	}

	usr, err := auth.Websocket(conn)

	if err != nil {
		logrus.Error("Error to auth websocket")
	}

	clients = append(clients, WebSocketClient{usr, conn})
}

func WebsocketController(c *gin.Context) {
	clients = make([]WebSocketClient, 0)
	wshandler(c.Writer, c.Request)
}

func Send(noti model.Notification) {
	for _, client := range clients {
		client.Conn.WriteMessage(0, []byte("oi"))
	}
}
