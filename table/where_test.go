package table

import (
	"log"
	"testing"
)

func TestTable_Where(t *testing.T) {

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
	db.Table("Hello").Where("number = ?", 3)
}
