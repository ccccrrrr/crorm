package test

import (
	"github.com/ccccrrrr/crorm"
	"log"
	"reflect"
	"testing"
)

func TestTable_Delete(t *testing.T) {
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
	type hello struct {
		Name string
		Say string
		Number int64
	}
	hhhhello := hello{}

	db.Table("Hello").Where("number = ?", 3).Delete()
	log.Printf("info: %v\n", hhhhello)

}
