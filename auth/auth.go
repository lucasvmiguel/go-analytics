package auth

import (
	"encoding/json"
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

const KEY_STUB = "123456"

type WebSocketAuth struct {
	Key string
}

func Websocket(conn *websocket.Conn) (string, error) {
	client := WebSocketAuth{}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.Error("Error to auth websocket, error to read message")
			return "", errors.New("Error to auth websocket, error to read message")
		}
		json.Unmarshal(msg, client)

		if client.Key == KEY_STUB {
			return client.Key, nil
		} else {
			return "", errors.New("invalid key" + client.Key)
		}
	}
}
