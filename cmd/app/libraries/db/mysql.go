package db

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
)

func GetWhere(connection Connection, model reflect.Type, whereClause string, sort string, limit int) ([]interface{}, error) {

	database, databaseError := sqlx.Open("mysql", connection.UserName+":"+connection.Password+"@tcp("+connection.IpAddress+":"+connection.Port+")/"+connection.Database)

	if databaseError != nil {
		return nil, databaseError
	}

	defer database.Close()

	rows, queryError := database.Queryx("SELECT * FROM " + connection.Table + " WHERE " + whereClause)

	// if there is an error inserting, handle it
	if queryError != nil {
		return nil, queryError
	}

	// be careful deferring Queries if you are using transactions
	defer rows.Close()

	var returnCollection []interface{}

	for rows.Next() {
		modelInstance := helper.NewEntityFromReflection(model)
		err := rows.StructScan(modelInstance)
		if err != nil {
			log.Fatal(err)
		}
		returnCollection = append(returnCollection, modelInstance)
	}

	return returnCollection, nil
}
