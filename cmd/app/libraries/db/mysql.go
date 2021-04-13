package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
)

func GetWhere(connection Connection, model reflect.Type, whereClause string, sort string, limit int) ([]map[string]interface{}, error) {

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

	var returnCollection []map[string]interface{}

	cols, _ := rows.ColumnTypes()

	pointers := make([]interface{}, len(cols))
	modelInstance := make(map[string]interface{}, len(cols))

	for index, column := range cols {
		var value interface{}

		modelInstance[column.Name()] = value
		pointers[index] = value
	}

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			log.Fatal(err)
		}
		returnCollection = append(returnCollection, modelInstance)
	}

	return returnCollection, nil
}
