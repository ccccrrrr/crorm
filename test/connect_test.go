package test

import (
	"github.com/ccccrrrr/crorm"
	"log"
	"reflect"
	"testing"
)

func TestOpen(t *testing.T) {
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
		log.Println("dbname: " + db.DBName)
		return
}

