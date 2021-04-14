package dtos

import (
	"errors"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/builder"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/db"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/helper"
	"github.com/google/uuid"
	"reflect"
)

type Notifications struct {
	builder builder.Builder
}

func (notify *Notifications) GetById(userId int) (*Notification, error) {
	connection := notify.getConnection()
	model := Notification{}
	interfaceModel, error := notify.builder.GetById(userId, connection, reflect.TypeOf(model))

	if error != nil {
		return nil, error
	}

	returnModel := notify.assignInterfaceModel(interfaceModel)

	return returnModel, nil
}

func (notify *Notifications) GetByUuid(userUuid uuid.UUID) (*Notification, error) {
	connection := notify.getConnection()
	model := Notification{}
	interfaceModel, error := notify.builder.GetByUuid(userUuid, connection, reflect.TypeOf(model))

	if error != nil {
		return nil, error
	}

	returnModel := notify.assignInterfaceModel(interfaceModel)

	return returnModel, nil
}

func (notify *Notifications) GetWhere(whereClause string, sort string, limit int) ([]*Notification, error) {
	connection := notify.getConnection()
	model := Notification{}
	interfaceCollection, error := notify.builder.GetWhere(connection, reflect.TypeOf(model), whereClause, sort, limit)

	if error != nil {
		return nil, error
	}

	collection := make([]*Notification, len(interfaceCollection))

	for i := 0; i < len(interfaceCollection); i++ {
		interfaceEntity := interfaceCollection[i]
		collectionEntity := notify.assignInterfaceModel(interfaceEntity)

		if collectionEntity.NotifyId == -1 {
			continue
		}

		collection[i] = collectionEntity
	}

	if len(collection) == 0 {
		return nil, errors.New("no rows returned")
	}

	return collection, nil
}

func (notify *Notifications) getNotifications() string {

	return ""
}

func (notify *Notifications) getConnection() db.Connection {
	connection := db.Connection{}
	return connection.GetNotification("notification", "notify_id", "sys_row_id")
}

func (notify *Notifications) assignInterfaceModel(model map[string]interface{}) *Notification {
	returnModel := &Notification{}
	returnModel.NotifyId = helper.CastAsNullableInt(model["notify_id"])
	returnModel.CompanyId = helper.CastAsNullableInt(model["company_id"])
	returnModel.UserId = helper.CastAsNullableInt(model["user_id"])
	returnModel.CardId = helper.CastAsNullableInt(model["card_id"])
	returnModel.NotifyType = helper.CastAsNullableInt(model["notify_type"])
	returnModel.EntityId = helper.CastAsNullableInt(model["entity_id"])
	returnModel.EntityType = helper.CastAsNullableString(model["entity_type"])
	returnModel.Title = helper.CastAsNullableString(model["title"])
	returnModel.Description = helper.CastAsNullableString(model["description"])
	returnModel.Status = helper.CastAsNullableString(model["status"])
	returnModel.CreatedOn = helper.CastToNullableTime(model["created_on"])
	returnModel.SysRowId = helper.CastAsNullableUuid(model["sys_row_id"])

	return returnModel
}

type Notification struct {
	NotifyId    int             `field:"notify_id"`
	CompanyId   int             `field:"company_id"`
	UserId      int             `field:"user_id"`
	CardId      int             `field:"card_id"`
	NotifyType  int             `field:"notify_type"`
	EntityId    int             `field:"entity_id"`
	EntityType  string          `field:"entity_type"`
	Title       string          `field:"title"`
	Description string          `field:"description"`
	Status      string          `field:"status"`
	CreatedOn   helper.NullTime `field:"created_on"`
	SysRowId    helper.NullUuid `field:"sys_row_id"`
}
