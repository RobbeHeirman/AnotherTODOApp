package persistence

import (
	"fmt"
	"reflect"
	"strings"
)

func CreateInsertQuery(tableName string, reflectedType reflect.Type, quantity int) string {

	fieldCount := reflectedType.NumField()
	fieldNames := make([]string, fieldCount)
	for i := 0; i < fieldCount; i++ {
		field := reflectedType.Field(i)
		tag, ok := field.Tag.Lookup("db")
		if !ok {
			tag = strings.ToLower(field.Name)
		}
		fieldNames[i] = tag
	}
	placeholders := make([]string, quantity)
	for i := 0; i < quantity; i++ {
		placeholders[i] = generatePlaceholderString(i*fieldCount, fieldCount)
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", tableName, strings.Join(fieldNames, ", "), strings.Join(placeholders, ","))
}

func generatePlaceholderString(lastNum int, placeholders int) string {
	nums := make([]string, placeholders)
	for i := lastNum + 1; i <= lastNum+placeholders; i++ {
		nums[i-lastNum-1] = fmt.Sprintf("$%d", i)
	}
	return fmt.Sprintf("(%s)", strings.Join(nums, ", "))
}
