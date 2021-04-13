package dtos

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/builder"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/db"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type VisitorBrowsers struct {
	builder builder.Builder
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) GetById(userId int) (*VisitorBrowser, error) {
	connection := vb.getConnection()
	model := VisitorBrowser{}
	interfaceModel, error := vb.builder.GetById(userId, connection, reflect.TypeOf(model))

	if error != nil {
		return nil, error
	}

	returnModel := vb.assignInterfaceModel(interfaceModel)

	return returnModel, nil
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) GetByUuid(userUuid uuid.UUID) (*VisitorBrowser, error) {
	connection := vb.getConnection()
	model := VisitorBrowser{}
	interfaceModel, error := vb.builder.GetByUuid(userUuid, connection, reflect.TypeOf(model))

	if error != nil {
		return nil, error
	}

	returnModel := vb.assignInterfaceModel(interfaceModel)

	return returnModel, nil
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) GetWhere(whereClause string, sort string, limit int) ([]*VisitorBrowser, error) {
	connection := vb.getConnection()
	model := VisitorBrowser{}
	interfaceCollection, error := vb.builder.GetWhere(connection, reflect.TypeOf(model), whereClause, sort, limit)

	if error != nil {
		return nil, error
	}

	collection := make([]*VisitorBrowser, len(interfaceCollection))

	for i := 0; i < len(interfaceCollection); i++ {
		interfaceEntity := interfaceCollection[i]
		collectionEntity := vb.assignInterfaceModel(interfaceEntity)
		collection[i] = collectionEntity
	}

	return collection, nil
}

// LocalAddr returns the local network address.
func (vb *VisitorBrowsers) getConnection() db.Connection {
	connection := db.Connection{}
	return connection.GetTraffic("visitor_browser", "visitor_browser_id", "browser_cookie")
}

func (vb *VisitorBrowsers) assignInterfaceModel(model map[string]interface{}) *VisitorBrowser {
	returnModel := &VisitorBrowser{}
	returnModel.VisitorBrowserId, _ = model["visitor_browser_id"].(int)
	returnModel.CompanyId, _ = model["company_id"].(int)
	returnModel.UserId, _ = model["user_id"].(int)
	returnModel.ContactId, _ = model["contact_id"].(int)
	returnModel.BrowserCookie, _ = model["browser_cookie"].(string)
	returnModel.BrowserIp, _ = model["browser_ip"].(string)
	returnModel.DeviceType, _ = model["device_type"].(string)
	returnModel.LoggedInAt, _ = time.Parse("2020-04-15 13:05:01", model["logged_in_at"].(string))
	returnModel.LastUpdated, _ = time.Parse("2020-04-15 13:05:01", model["last_updated"].(string))
	returnModel.CreatedOn, _ = time.Parse("2020-04-15 13:05:01", model["created_on"].(string))

	return returnModel
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
