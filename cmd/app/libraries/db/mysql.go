package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"reflect"
)

func PostgresGetWhere(connection Connection, model reflect.Type, whereClause string, sort string, limit int) ([]map[string]interface{}, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connection.IpAddress, connection.Port, connection.UserName, connection.Password, connection.Database)
	database, databaseError := sql.Open("postgres", psqlInfo)

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

	var returnCollection []map[string]interface{}

	cols, _ := rows.ColumnTypes()

	pointers := make([]interface{}, len(cols))
	modelInstance := make(map[string]interface{}, len(cols))

	for index, column := range cols {
		var value interface{}

		switch column.DatabaseTypeName() {
		case "INT":
			value = new(sql.NullInt32)
		case "VARCHAR", "STRING":
			value = new(sql.NullString)
		case "TIMESTAMP", "DATETIME":
			value = new(sql.NullString)
		case "UUID":
			value = new(sql.NullString)
		default:
			value = new(interface{}) // destination must be a pointer

			// use to figure out types for columns
			// log.Println(column.Name(), column.DatabaseTypeName())
		}

		modelInstance[column.Name()] = value
		pointers[index] = value
	}

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			log.Println(err)
		}
		returnCollection = append(returnCollection, modelInstance)
	}

	return returnCollection, nil
}

func MysqlGetWhere(connection Connection, model reflect.Type, whereClause string, sort string, limit int) ([]map[string]interface{}, error) {

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

	var returnCollection []map[string]interface{}

	cols, _ := rows.ColumnTypes()

	pointers := make([]interface{}, len(cols))
	modelInstance := make(map[string]interface{}, len(cols))

	for index, column := range cols {
		var value interface{}

		switch column.DatabaseTypeName() {
		case "INT":
			value = new(sql.NullInt32)
		case "VARCHAR", "STRING":
			value = new(sql.NullString)
		case "TIMESTAMP", "DATETIME":
			value = new(sql.NullString)
		case "UUID":
			value = new(sql.NullString)
		default:
			value = new(interface{}) // destination must be a pointer

			// use to figure out types for columns
			// log.Println(column.Name(), column.DatabaseTypeName())
		}

		modelInstance[column.Name()] = value
		pointers[index] = value
	}

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			log.Println(err)
		}
		returnCollection = append(returnCollection, modelInstance)
	}

	return returnCollection, nil
}
