package structure

import (
	"log"
	"testing"
)

func TestTable_insert(t *testing.T) {
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
	type hello2 struct {
		Name string
		Say string
		Number int64
	}
	hhello := hello2{
		"China",
		"nihao",
		3,
	}
	db.Table("hello2").Insert(&hhello)
}
