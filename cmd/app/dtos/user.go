package dtos

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/builder"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/db"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Users struct {
	builder builder.Builder
}

// LocalAddr returns the local network address.
func (users *Users) GetById(userId int) (*User, error) {
	connection := users.getConnection()
	model := User{}
	interfaceModel, error := users.builder.GetById(userId, connection, reflect.TypeOf(model))

	if error != nil {
		return nil, error
	}

	returnModel := users.assignInterfaceModel(interfaceModel)

	return returnModel, nil
}

// LocalAddr returns the local network address.
func (users *Users) GetByUuid(userUuid uuid.UUID) (*User, error) {
	connection := users.getConnection()
	model := User{}
	interfaceModel, error := users.builder.GetByUuid(userUuid, connection, reflect.TypeOf(model))

	if error != nil {
		return nil, error
	}

	returnModel := users.assignInterfaceModel(interfaceModel)

	return returnModel, nil
}

// LocalAddr returns the local network address.
func (users *Users) getConnection() db.Connection {
	connection := db.Connection{}
	return connection.GetMain("user", "user_id", "sys_row_id")
}

func (users *Users) assignInterfaceModel(model map[string]interface{}) *User {
	returnModel := &User{}
	returnModel.UserId, _ = model["user_id"].(int)
	returnModel.CompanyId, _ = model["company_id"].(int)
	returnModel.OriginatorId, _ = model["sponsor_id"].(int)
	returnModel.FirstName, _ = model["first_name"].(string)
	returnModel.LastName, _ = model["last_name"].(string)
	returnModel.NamePrefix, _ = model["name_prefix"].(string)
	returnModel.MiddleName, _ = model["middle_name"].(string)
	returnModel.NameSuffix, _ = model["name_sufx"].(string)
	returnModel.Username, _ = model["username"].(string)
	returnModel.Password, _ = model["password"].(string)
	returnModel.PasswordResetToken, _ = model["password_reset_token"].(string)
	returnModel.Pin, _ = model["pin"].(int)
	returnModel.UserEmail, _ = model["user_email"].(string)
	returnModel.UserPhone, _ = model["user_phone"].(string)

	return returnModel
}

type User struct {
	UserId             int       `field:"user_id"`
	CompanyId          int       `field:"company_id"`
	OriginatorId       int       `field:"sponsor_id"`
	FirstName          string    `field:"first_name"`
	LastName           string    `field:"last_name"`
	NamePrefix         string    `field:"name_prefix"`
	MiddleName         string    `field:"middle_name"`
	NameSuffix         string    `field:"name_sufx"`
	Username           string    `field:"username"`
	Password           string    `field:"password"`
	PasswordResetToken string    `field:"password_reset_token"`
	Pin                int       `field:"pin"`
	UserEmailId        int       `field:"user_email"`
	UserEmail          string    `field:"user_email_value"`
	UserPhoneId        int       `field:"user_phone"`
	UserPhone          string    `field:"user_phone_value"`
	CreatedOn          time.Time `field:"created_on"`
	CreatedBy          int       `field:"created_by"`
	LastUpdated        time.Time `field:"last_updated"`
	UpdateBy           int       `field:"update_by"`
	Status             string    `field:"status"`
	PreferredName      string    `field:"preferred_name"`
	LastLogin          time.Time `field:"last_login"`
	SysRowId           uuid.UUID `field:"sys_row_id"`
}
