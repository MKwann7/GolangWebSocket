package main

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/errorManagement"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/process"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	upgradeConnection = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := router()
	middleware := negroni.Classic()
	middleware.UseHandler(router)

	http.ListenAndServe(":"+os.Getenv("PORT_NUM"), middleware)
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

	user, validationError := process.ValidateConnection(webRequest)

	if validationError != nil {
		errorManagement.HandleErr(responseWriter, validationError, http.StatusBadRequest)
		return
	}

	socketConnection, upgradeError := upgradeConnection.Upgrade(responseWriter, webRequest, nil)

	if upgradeError != nil {
		errorManagement.HandleErr(responseWriter, upgradeError, http.StatusInternalServerError)
		return
	}

	defer socketConnection.Close()

	for {

		notificationCheckResult := process.CheckForNewNotifications(user)

		if notificationCheckResult != "" {
			writeNotificationError := socketConnection.WriteMessage(1, []byte(notificationCheckResult))

			if writeNotificationError != nil {
				errorManagement.HandleErr(responseWriter, writeNotificationError, http.StatusInternalServerError)
				break
			}
		}

		inboundMessageResult := process.ProcessInboundMessage(socketConnection, responseWriter)

		if inboundMessageResult != nil {
			writeNotificationError := socketConnection.WriteMessage(1, []byte(inboundMessageResult))

			if writeNotificationError != nil {
				errorManagement.HandleErr(responseWriter, writeNotificationError, http.StatusInternalServerError)
				break
			}
		}

		time.Sleep(1 * time.Second)
	}
}
