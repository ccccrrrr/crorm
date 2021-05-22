package util

import (
	"reflect"
	"strings"
)

func GetTableName(entity interface{}) string {
	t := reflect.TypeOf(entity)
	stringSplit := strings.Split(t.Elem().String(), ".")
	tableName := strings.ToLower(stringSplit[len(stringSplit) - 1]) + "s"
	return tableName
}
