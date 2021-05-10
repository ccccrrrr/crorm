package main

import (
	crorm "crorm/structure"
	"log"
)

type Info struct {
	Name string `crorm:"name"`
	Age int `crorm:"age"`
	Gender string `crorm:"gender"`
}
func main() {
	config := crorm.DBConfig{
		UserName: "root",
		UserPassword: "xxx",
		Port : "3306",
		Ip : "localhost",
		DBName: "db1",
	}
	db, err := crorm.Open(config)

	if err != nil {
		log.Println(err)
		return
	}

	var info Info

	db.CreateTable("Info1", &info)

	table := db.Table("info1")
	var res Info

	info1 := Info{
		"ccrrrr",
		111,
		"nnn",
	}
	info2 := Info {
		"ccrr",
		111,
		"nnn",
	}


	table.Insert(&info1)
	table.Insert(&info2)

	var res Info

	table.Where("name = ?", "ccrr").First(&res)

	log.Println(res)

	table.Where("name = ?", "ccrrrr").Update("name = ?", "ccccrrrr")

	table.Where("name = ?", "ccccrrrr").Delete(&res)

	table.First(&res)

	log.Println(res)
}
