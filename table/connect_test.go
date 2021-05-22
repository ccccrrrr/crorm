package table

import (
	"crorm/config"
	"log"
	"reflect"
	"testing"
)

func TestOpen(t *testing.T) {
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
		log.Println("dbname: " + db.DBName)
		return
}

