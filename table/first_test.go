package table

import (
	"crorm/config"
	"log"
	"reflect"
	"testing"
)

func Test_list_First(t *testing.T) {
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
		Name string `crorm:"name"`
		Say string `crorm:"say"`
		Number int64 `crorm:"number"`
	}

	hhhhello := hello2{}

	_, _ = db.Table("hello2").Where("number = ?", 3).First(&hhhhello)
	log.Printf("info: %v\n", hhhhello)

}
