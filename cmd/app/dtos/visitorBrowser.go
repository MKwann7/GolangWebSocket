package dtos

import (
	"errors"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/builder"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/db"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/helper"
	"github.com/google/uuid"
	"reflect"
)

type VisitorBrowsers struct {
	builder builder.Builder
}

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

		if collectionEntity.VisitorBrowserId == -1 {
			continue
		}

		collection[i] = collectionEntity
	}

	if len(collection) == 0 {
		return nil, errors.New("no rows returned")
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

	returnModel.VisitorBrowserId = helper.CastAsNullableInt(model["visitor_browser_id"])
	returnModel.CompanyId = helper.CastAsNullableInt(model["company_id"])
	returnModel.UserId = helper.CastAsNullableInt(model["user_id"])
	returnModel.ContactId = helper.CastAsNullableInt(model["contact_id"])
	returnModel.BrowserCookie = helper.CastAsNullableString(model["browser_cookie"])
	returnModel.BrowserIp = helper.CastAsNullableString(model["browser_ip"])
	returnModel.DeviceType = helper.CastAsNullableString(model["device_type"])
	returnModel.LoggedInAt = helper.CastToNullableTime(model["logged_in_at"])
	returnModel.LastUpdated = helper.CastToNullableTime(model["last_updated"])
	returnModel.CreatedOn = helper.CastToNullableTime(model["created_on"])

	return returnModel
}

type VisitorBrowser struct {
	VisitorBrowserId int             `field:"visitor_browser_id"`
	CompanyId        int             `field:"company_id"`
	UserId           int             `field:"user_id"`
	ContactId        int             `field:"contact_id"`
	BrowserCookie    string          `field:"browser_cookie"`
	BrowserIp        string          `field:"browser_ip"`
	DeviceType       string          `field:"device_type"`
	LoggedInAt       helper.NullTime `field:"logged_in_at"`
	LastUpdated      helper.NullTime `field:"last_updated"`
	CreatedOn        helper.NullTime `field:"created_on"`
}
