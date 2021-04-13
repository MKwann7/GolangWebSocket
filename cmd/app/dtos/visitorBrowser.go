package dtos

import (
	"errors"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/builder"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/db"
	"github.com/google/uuid"
	"time"
)

type VisitorBrowsers struct {
	builder builder.Builder
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) GetById(userId int) (*VisitorBrowser, error) {
	connection := vb.getConnection()
	user := VisitorBrowser{}
	interfaceModel, error := vb.builder.GetById(userId, connection, &user)

	if error != nil {
		return nil, error
	}

	model, success := interfaceModel.(*VisitorBrowser)

	if success != true {
		return nil, errors.New("could not process VisitorBrowser model")
	}

	return model, nil
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) GetByUuid(userUuid uuid.UUID) (*VisitorBrowser, error) {
	connection := vb.getConnection()
	user := VisitorBrowser{}
	interfaceModel, error := vb.builder.GetByUuid(userUuid, connection, user)

	if error != nil {
		return nil, error
	}

	model, success := interfaceModel.(*VisitorBrowser)

	if success != true {
		return nil, errors.New("could not process VisitorBrowser model")
	}

	return model, nil
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) GetWhere(whereClause string, sort string, limit int) ([]*VisitorBrowser, error) {
	connection := vb.getConnection()
	user := VisitorBrowser{}
	interfaceCollection, error := vb.builder.GetWhere(connection, user, whereClause, sort, limit)

	if error != nil {
		return nil, error
	}

	collection := make([]*VisitorBrowser, len(interfaceCollection))

	for i := 0; i < len(interfaceCollection); i++ {
		interfaceEntity := interfaceCollection[i]
		collectionEntity, success := interfaceEntity.(*VisitorBrowser)

		if success != true {
			continue
		}

		collection[i] = collectionEntity
	}

	return collection, nil
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) getConnection() db.Connection {
	return db.Connection{}.GetTraffic("visitor_browser", "visitor_browser_id", "browser_cookie")
}

type VisitorBrowser struct {
	VisitorBrowserId int       `field:"visitor_browser_id"`
	CompanyId        int       `field:"company_id"`
	UserId           int       `field:"user_id"`
	ContactId        int       `field:"contact_id"`
	BrowserCookie    string    `field:"browser_cookie"`
	BrowserIp        string    `field:"browser_ip"`
	DeviceType       string    `field:"device_type"`
	LoggedInAt       time.Time `field:"logged_in_at"`
	LastUpdated      time.Time `field:"last_updated"`
	CreatedOn        time.Time `field:"created_on"`
}
