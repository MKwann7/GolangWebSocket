package db

import (
	"os"
)

type Connection struct {
	Table      string
	PrimaryKey string
	UuidKey    string
	IpAddress  string
	Port       string
	Database   string
	UserName   string
	Password   string
	DbType     string
}

const MySQL = "mysql"
const Postgres = "postgres"

func (connection *Connection) GetMain(tableName string, userKey string, uuidKey string) Connection {
	return Connection{
		os.Getenv("MAIN_DB_NAME") + "." + tableName,
		userKey,
		uuidKey,
		os.Getenv("MAIN_DB_HOST"),
		os.Getenv("MAIN_DB_PORT"),
		os.Getenv("MAIN_DB_NAME"),
		os.Getenv("MAIN_DB_USER"),
		os.Getenv("MAIN_DB_PASS"),
		MySQL}
}

func (connection *Connection) GetTraffic(tableName string, userKey string, uuidKey string) Connection {
	return Connection{
		os.Getenv("TRAFFIC_DB_NAME") + "." + tableName,
		userKey,
		uuidKey,
		os.Getenv("TRAFFIC_DB_HOST"),
		os.Getenv("TRAFFIC_DB_PORT"),
		os.Getenv("TRAFFIC_DB_NAME"),
		os.Getenv("TRAFFIC_DB_USER"),
		os.Getenv("TRAFFIC_DB_PASS"),
		MySQL}
}
