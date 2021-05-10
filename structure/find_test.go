package structure

import (
	"log"
	"testing"
)

func TestTable_Find(t *testing.T) {
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

	var hhhhello []hello2
	//hhhhello := make([]Hello, 0)
	table := db.Table("hello2")
	log.Println(table.TableName)
	table.Where("number = ?", 3).Find(&hhhhello)

	log.Printf("info: %v\n", hhhhello)
}
