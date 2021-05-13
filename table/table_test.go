package table

import (
	"log"
	"testing"
)

func TestTable(t *testing.T) {
		config := DBConfig{
			UserName: "root",
			UserPassword: "lotus20001006",
			Port : "3306",
			Ip : "localhost",
			DBName: "db1",
		}
		db, err := Open(config)
		if err != nil {
			log.Println(err)
			return
		}
		db.Table("hello2")
}
