package table

import (
	"log"
	"testing"
)

func TestTable_Delete(t *testing.T) {
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
	type hello struct {
		Name string
		Say string
		Number int64
	}
	hhhhello := hello{}

	db.Table("Hello").Where("number = ?", 3).Delete(&hhhhello)
	log.Printf("info: %v\n", hhhhello)

}
