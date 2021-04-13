package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
)

func GetWhere(connection Connection, model interface{}, whereClause string, sort string, limit int) ([]interface{}, error) {

	database, databaseError := sql.Open("mysql", connection.UserName+":"+connection.Password+"@tcp("+connection.IpAddress+":"+connection.Port+")/"+connection.Database)

	if databaseError != nil {
		return nil, databaseError
	}

	defer database.Close()

	rows, queryError := database.Query("SELECT * FROM " + connection.Table + " WHERE " + whereClause)

	// if there is an error inserting, handle it
	if queryError != nil {
		return nil, queryError
	}

	// be careful deferring Queries if you are using transactions
	defer rows.Close()

	var returnCollection []interface{}

	for rows.Next() {
		modelInstance := reflect.New(reflect.TypeOf(model))
		modelElements := reflect.ValueOf(&modelInstance).Elem()
		numCols := modelElements.NumField()
		columns := make([]interface{}, numCols)

		for i := 0; i < numCols; i++ {
			field := modelElements.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		returnCollection = append(returnCollection, modelInstance)
	}

	return returnCollection, nil
}
