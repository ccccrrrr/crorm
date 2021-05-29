package test

import (
	"github.com/ccccrrrr/crorm"
	"log"
	"reflect"
	"testing"
)

func TestTable_Find(t *testing.T) {
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
	type hello2 struct {
		Name string
		Say string
		Number int64
	}

	var hhhhello []hello2
	//hhhhello := make([]Hello, 0)
	table := db.Table("hello2")
	log.Println(table.TableName)
	table.Where("number = ?", 3).Find(&hhhhello)

	log.Printf("info: %v\n", hhhhello)
}
