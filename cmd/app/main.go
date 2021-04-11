package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
)

var (
	upgradeConnection = websocket.Upgrader{}
)

func main() {

	// configuration
	port := 8080

	router := router()
	middleware := negroni.Classic()
	middleware.UseHandler(router)

	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), middleware))
}

func router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path("/socket").
		HandlerFunc(socketListen)

	return router
}

func socketListen(responseWriter http.ResponseWriter, webRequest *http.Request) {

	socketConnection, error := upgradeConnection.Upgrade(responseWriter, webRequest, nil)

	if error != nil {
		handleErr(responseWriter, error, http.StatusInternalServerError)
		return
	}

	defer socketConnection.Close()

	for {
		messageType, messageString, messageError := socketConnection.ReadMessage()

		if messageError != nil {
			handleErr(responseWriter, messageError, http.StatusInternalServerError)
			continue
		}

		if messageType != websocket.TextMessage {
			handleErr(responseWriter, errors.New("only text message are supported"), http.StatusNotImplemented)
			continue
		}

		if string(messageString) == "" {
			continue
		}

		writeError := socketConnection.WriteMessage(messageType, []byte(messageString))

		if writeError != nil {
			handleErr(responseWriter, writeError, http.StatusInternalServerError)
			break
		}
	}
}

func handleErr(responseWriter http.ResponseWriter, err error, status int) {
	msg, err := json.Marshal(&httpErr{
		Msg:  err.Error(),
		Code: status,
	})
	if err != nil {
		msg = []byte(err.Error())
	}
	http.Error(responseWriter, string(msg), status)
}

type httpErr struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
