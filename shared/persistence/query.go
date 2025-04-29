package persistence

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func CreateInsertQuery[T any](tableName string, object ...T) (string, error) {
	if len(object) == 0 {
		return "", errors.New("need atleast one object")
	}
	fields := reflect.TypeOf(object[0])
	fieldCount := fields.NumField()
	fieldNames := make([]string, fieldCount)
	for i := 0; i < fieldCount; i++ {
		field := fields.Field(i)
		tag, ok := field.Tag.Lookup("db")
		if !ok {
			tag = strings.ToLower(field.Name)
		}
		fieldNames[i] = tag

	}
	objectCount := len(object)
	placeholders := make([]string, objectCount)
	for i := 0; i < objectCount; i++ {
		placeholders[i] = generatePlaceholderString(i*fieldCount, fieldCount)
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", tableName, strings.Join(fieldNames, ", "), strings.Join(placeholders, ",")), nil
}

func generatePlaceholderString(lastNum int, placeholders int) string {
	nums := make([]string, placeholders)
	for i := lastNum + 1; i <= lastNum+placeholders; i++ {
		nums[i-lastNum-1] = fmt.Sprintf("$%d", i)
	}
	return fmt.Sprintf("(%s)", strings.Join(nums, ", "))
}
