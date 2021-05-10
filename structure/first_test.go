package structure

import (
	"log"
	"testing"
)

func Test_list_First(t *testing.T) {
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
		Name string `crorm:"name"`
		Say string `crorm:"say"`
		Number int64 `crorm:"number"`
	}

	hhhhello := hello2{}

	err = db.Table("hello2").Where("number = ?", 3).First(&hhhhello)
	log.Printf("info: %v\n", hhhhello)

}
