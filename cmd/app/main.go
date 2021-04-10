package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
)

func main() {

	// configuration
	port := 8080

	router := router()
	middleware := negroni.Classic()
	middleware.UseHandler(router)

	http.ListenAndServe(strconv.Itoa(port), middleware)
}

func router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path("/").
		HandlerFunc(socketListen)

	return router
}

func socketListen(responseWriter http.ResponseWriter, webRequest *http.Request) {

	upgradeConnection := websocket.Upgrader{}
	socketConnection, _ := upgradeConnection.Upgrade(responseWriter, webRequest, nil)

	defer socketConnection.Close()

	for {
		messageType, messageString, _ := socketConnection.ReadMessage()

		if messageType != websocket.TextMessage {
			continue
		}
		socketConnection.WriteMessage(messageType, []byte(messageString))
	}
}
