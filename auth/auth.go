package auth

import (
	"encoding/json"
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

type WebSocketAuth struct {
	Key string `json:"key"`
}

func Websocket(conn *websocket.Conn) (string, int, error) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return connectionError("Error to auth websocket, error to read message")
		}

		client := WebSocketAuth{}
		err = json.Unmarshal(msg, &client)
		if err != nil {
			return connectionError("invalid message received")
		}

		if isValidConnection(client.Key) {
			conn.WriteMessage(msgType, []byte("connected"))
			return client.Key, msgType, nil
		} else {
			return connectionError("invalid key received")
		}
	}
}

func isValidConnection(keySent string) bool {
	return keySent == viper.GetString("develop.websocket.key")
}

func connectionError(err string) (string, int, error) {
	logrus.Error(err)
	return "", 0, errors.New(err)
}
