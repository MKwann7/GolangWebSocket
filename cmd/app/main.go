package main

import (
	"fmt"
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

	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), middleware))

	fmt.Println("We are done!")
}

func router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path("/").
		HandlerFunc(socketListen)

	fmt.Println("Router created!")

	return router
}

func socketListen(responseWriter http.ResponseWriter, webRequest *http.Request) {

	upgradeConnection := websocket.Upgrader{}
	socketConnection, _ := upgradeConnection.Upgrade(responseWriter, webRequest, nil)

	fmt.Println("Socket connection upgraded!")

	defer socketConnection.Close()

	fmt.Println("Preparing to loop!")

	for {
		messageType, messageString, _ := socketConnection.ReadMessage()

		if messageType != websocket.TextMessage {
			continue
		}
		socketConnection.WriteMessage(messageType, []byte(messageString))
	}
}
