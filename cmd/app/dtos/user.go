package dtos

import (
	"errors"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/builder"
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/db"
	"github.com/google/uuid"
	"time"
)

type Users struct {
	builder builder.Builder
}

// LocalAddr returns the local network address.
func (users *Users) GetById(userId int) (*User, error) {
	connection := users.getConnection()
	user := User{}
	interfaceModel, error := users.builder.GetById(userId, connection, user)

	if error != nil {
		return nil, error
	}

	model, success := interfaceModel.(*User)

	if success != true {
		return nil, errors.New("could not process User model")
	}

	return model, nil
}

// LocalAddr returns the local network address.
func (users *Users) GetByUuid(userUuid uuid.UUID) (*User, error) {
	connection := users.getConnection()
	user := User{}
	interfaceModel, error := users.builder.GetByUuid(userUuid, connection, user)

	if error != nil {
		return nil, error
	}

	model, success := interfaceModel.(*User)

	if success != true {
		return nil, errors.New("could not process User model")
	}

	return model, nil
}

// LocalAddr returns the local network address.
func (users *Users) getConnection() db.Connection {
	return db.Connection{}.GetMain("user", "user_id", "sys_row_id")
}

type User struct {
	UserId             int       `field:"user_id"`
	CompanyId          int       `field:"company_id"`
	OriginatorId       int       `field:"sponsor_id"`
	FirstName          string    `field:"first_name"`
	LastName           string    `field:"last_name"`
	NamePrefix         string    `field:"name_prefix"`
	MiddleName         string    `field:"middle_name"`
	NameSuffix         string    `field:"name_suffix"`
	Username           string    `field:"username"`
	Password           string    `field:"password"`
	PasswordResetToken string    `field:"password_reset_token"`
	Pin                int       `field:"pin"`
	UserEmailId        int       `field:"user_email_id"`
	UserEmail          string    `field:"user_email"`
	UserPhoneId        int       `field:"user_phone_id"`
	UserPhone          string    `field:"user_phone"`
	CreatedOn          time.Time `field:"created_on"`
	CreatedBy          int       `field:"created_by"`
	LastUpdated        time.Time `field:"last_updated"`
	UpdateBy           int       `field:"update_by"`
	Status             string    `field:"status"`
	PreferredName      string    `field:"preferred_name"`
	LastLogin          time.Time `field:"last_login"`
	SysRowId           uuid.UUID `field:"sys_row_id"`
}
