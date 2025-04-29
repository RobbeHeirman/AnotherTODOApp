package persistence

import (
	testing2 "github.com/robbeheirman/todo/shared/testing"
	"reflect"
	"testing"
)

type TestObj struct {
	Test              string
	TestNum           int
	TaggedField       string `db:"tagged_field"`
	JsonTaggedField   string `json:"json_tagged_field"`
	DoubleTaggedField int    `json:"json_double_tagged_field" db:"double_tagged_field"`
}

func TestQueryInsert(t *testing.T) {
	obj := TestObj{
		Test:              "test",
		TestNum:           1,
		TaggedField:       "tagged_field",
		JsonTaggedField:   "json_tagged_field",
		DoubleTaggedField: 2,
	}
	typeOf := reflect.TypeOf(obj)

	expect := "INSERT INTO test (test, testnum, tagged_field, jsontaggedfield, double_tagged_field) VALUES ($1, $2, $3, $4, $5)"
	query := CreateInsertQuery("test", typeOf, 1)
	testing2.AssertEqual(t, expect, query)

	expect = "INSERT INTO test (test, testnum, tagged_field, jsontaggedfield, double_tagged_field) VALUES ($1, $2, $3, $4, $5),($6, $7, $8, $9, $10)"
	query = CreateInsertQuery("test", typeOf, 2)
	testing2.AssertEqual(t, expect, query)
}
