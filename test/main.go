package main

import (
	"crorm/config"
	"crorm/table"
	"log"
)

type Info struct {
	Name   string `crorm:"name"`
	Age    int    `crorm:"age"`
	Gender string `crorm:"gender"`
	Hello  string `crorm:"hello"`
}

func main() {

	db, _ := table.Open(config.Config)

	myTable := db.Table("Info11")

	infoFind := make([]Info, 20)

	//var infoFind []Info

	//info := Info{
	//	Name: "=_=",
	//	Age: 23,
	//	Gender: "male",
	//	Hello: "hello",
	//}
	//
	//myTable.Insert(&info)

	myTable.Find(&infoFind)

	log.Println(infoFind)

}
