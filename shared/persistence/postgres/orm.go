package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/robbeheirman/todo/shared/persistence"
	"reflect"
)

func InsertAndGetObjects[T any, U any](conn *pgx.Conn, table string, rowTransform pgx.RowToFunc[U], objects ...*T) ([]U, error) {
	if len(objects) == 0 {
		return []U{}, nil
	}
	reflectedType := reflect.TypeOf(*objects[0])
	queryTemplate := persistence.CreateInsertQuery(table, reflectedType, len(objects))

	fields := make([]interface{}, len(objects)*reflectedType.NumField())
	for i, object := range objects {
		reflectedValue := reflect.ValueOf(*object)
		NumFields := reflectedValue.NumField()
		for j := 0; j < NumFields; j++ {
			fields[i*NumFields+j] = reflectedValue.Field(j)
		}
	}

	result, err := conn.Query(context.Background(), queryTemplate, fields...)
	defer result.Close()
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(result, rowTransform)
}
