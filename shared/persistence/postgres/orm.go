package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/robbeheirman/todo/shared/persistence"
	"reflect"
)

func InsertObject[T any](conn *pgx.Conn, table string, objects ...T) (int, error) {
	if len(objects) == 0 {
		return 0, nil
	}
	reflectedType := reflect.TypeOf(objects[0])
	queryTemplate := persistence.CreateInsertQuery(table, reflectedType, len(objects))

	fields := make([]interface{}, len(objects)*reflectedType.NumField())
	for i, object := range objects {
		if i != 0 {
			reflectedType = reflect.TypeOf(object)
		}
		for j := 0; j < reflectedType.NumField(); j++ {
			fields[i*j] = reflectedType.Field(j)
		}
	}

	result, err := conn.Exec(context.Background(), queryTemplate, fields...)
	if err != nil {
		return 0, err
	}
	return int(result.RowsAffected()), err
}
