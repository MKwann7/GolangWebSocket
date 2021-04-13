package process

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/dtos"
	"net/http"
)

func ValidateConnection(webRequest *http.Request) (*dtos.User, error) {

	authUuidString := webRequest.URL.Query().Get("auth")

	visitors := dtos.VisitorBrowsers{}
	collection, err := visitors.GetWhere("browser_cookie = '"+authUuidString+"'", "ASC", 1)

	if err != nil {
		return nil, err
	}

	vistiorBrowser := collection[0]

	users := dtos.Users{}
	user, err := users.GetById(vistiorBrowser.UserId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
