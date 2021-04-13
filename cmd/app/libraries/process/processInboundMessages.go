package process

import (
	"errors"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/errorManagement"
	"github.com/gorilla/websocket"
	"net/http"
)

func ProcessInboundMessage(socketConnection *websocket.Conn, responseWriter http.ResponseWriter) []byte {

	messageType, messageString, messageError := socketConnection.ReadMessage()

	if messageError != nil {
		errorManagement.HandleErr(responseWriter, messageError, http.StatusInternalServerError)
		return nil
	}

	if messageType != websocket.TextMessage {
		errorManagement.HandleErr(responseWriter, errors.New("only text message are supported"), http.StatusNotImplemented)
		return nil
	}

	if string(messageString) == "" {
		return nil
	}

	return processMessage(socketConnection, messageType, messageString, responseWriter)
}

func processMessage(socketConnection *websocket.Conn, messageType int, messageString []byte, responseWriter http.ResponseWriter) []byte {

	return messageString
}

type messageResponse struct {
	isServer         bool
	getNotifications func() string
}
