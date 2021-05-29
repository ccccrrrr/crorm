package test

import (
	"github.com/ccccrrrr/crorm"
	"log"
	"reflect"
	"testing"
)

//func TestDB_CreateTable(t *testing.T) {
//	config := DBConfig{
//		UserName: "root",
//		UserPassword: "lotus20001006",
//		Port : "3306",
//		Ip : "localhost",
//		DBName: "db1",
//	}
//	db, err := connection.Open(config)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	result, err := db.DefaultDB.Exec("CREATE TABLE `userinfo` (\n " +
//		"`uid` int(10) NOT NULL AUTO_INCREMENT,\n  " +
//		"`username` varchar(64) DEFAULT NULL,\n " +
//		"`departname` varchar(64) DEFAULT NULL,\n " +
//		"`created` date DEFAULT NULL,\n " +
//		"PRIMARY KEY (`uid`)\n" +
//		")")
//	log.Println(result)
//
//}


func TestDB_CreateTable(t *testing.T) {
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
		hhhello := hello2{}
		db.CreateTable(&hhhello)

}

