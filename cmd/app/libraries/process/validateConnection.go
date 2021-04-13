package process

import (
	"errors"
	"github.com/MKwann7/GolangWebSocket/cmd/app/dtos"
	"net/http"
)

func ValidateConnection(webRequest *http.Request) (*dtos.User, error) {

	authUuidString := webRequest.URL.Query().Get("auth")

	collection, err := dtos.VisitorBrowsers{}.GetWhere("browser_cookie = '"+authUuidString+"'", "ASC", 1)

	if err != nil {
		return nil, errors.New("we were unable to find an active session")
	}

	vistiorBrowser := collection[0]

	user, err := dtos.Users{}.GetById(vistiorBrowser.UserId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
