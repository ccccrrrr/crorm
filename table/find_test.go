package table

import (
	"crorm/config"
	"log"
	"reflect"
	"testing"
)

func TestTable_Find(t *testing.T) {
	_config := config.Config
	var standardConfig DBConfig
	for i := 0; i < 5; i++ {
		reflect.ValueOf(&standardConfig).Elem().Field(i).Set(reflect.ValueOf(&_config).Elem().Field(i))
	}
	db, err := Open(standardConfig)
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
