package builder

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/libraries/db"
	"github.com/google/uuid"
	"reflect"
	"strconv"
)

type Builder struct {
}

func (builder *Builder) GetById(entityId int, connection db.Connection, model reflect.Type) (map[string]interface{}, error) {
	entityCollection, error := builder.GetWhere(connection, model, connection.PrimaryKey+" = "+strconv.Itoa(entityId), "ASC", 1)

	if error != nil {
		return nil, error
	}

	return entityCollection[0], nil
}

func (builder *Builder) GetByUuid(entityUuid uuid.UUID, connection db.Connection, model reflect.Type) (map[string]interface{}, error) {
	entityCollection, error := builder.GetWhere(connection, model, connection.UuidKey+" = '"+entityUuid.String()+"'", "ASC", 1)

	if error != nil {
		return nil, error
	}

	return entityCollection[0], nil
}

func (builder *Builder) GetWhere(connection db.Connection, model reflect.Type, whereClause string, sort string, limit int) ([]map[string]interface{}, error) {
	switch connection.DbType {
	case "postgres":
		return db.PostgresGetWhere(connection, model, whereClause, sort, limit)
	default:
		return db.MysqlGetWhere(connection, model, whereClause, sort, limit)
	}
}
