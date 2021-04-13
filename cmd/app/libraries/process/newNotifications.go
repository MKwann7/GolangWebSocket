package process

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/dtos"
)

func CheckForNewNotifications(user *dtos.User) string {

	return user.FirstName
}
