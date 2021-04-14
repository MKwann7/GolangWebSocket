package process

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/dtos"
	"log"
	"os"
	"strconv"
)

func CheckForNewNotifications(user *dtos.User) string {

	notifications := dtos.Notifications{}
	collection, err := notifications.GetWhere("user_id = '"+strconv.Itoa(user.UserId)+"'", "ASC", 1)

	if err != nil {
		panic(err)
	}

	log.Println(collection)

	os.Exit(1)

	return user.FirstName
}
