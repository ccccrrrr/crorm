package test

import (
	"github.com/ccccrrrr/crorm"
	"log"
	"reflect"
	"testing"
)

func Test_list_First(t *testing.T) {
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
		Name string `crorm:"name"`
		Say string `crorm:"say"`
		Number int64 `crorm:"number"`
	}

	hhhhello := hello2{}

	_, _ = db.Table("hello2").Where("number = ?", 3).First(&hhhhello)
	log.Printf("info: %v\n", hhhhello)

}
