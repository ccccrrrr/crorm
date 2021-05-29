package test

import (
	"github.com/ccccrrrr/crorm"
	"log"
	"reflect"
	"testing"
)

type hello struct {
	Name string
	Say string
	Number int
}

func TestDB_CreateOrOverride(t *testing.T) {
	_config := Config
	var standardConfig crorm.DBConfig
	for i := 0; i < 5; i++ {
		reflect.ValueOf(&standardConfig).Elem().Field(i).Set(reflect.ValueOf(&_config).Elem().Field(i))
	}
	db, err := crorm.Open(standardConfig)
	if err != nil {
		log.Println(err)
		return
	}

	var entity hello

	_, _ = db.CreateOrOverrideTable(&entity)
	// search table in certain database with correct name, and every element name and type
	table, err := db.Sync(&entity)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(table)
}
