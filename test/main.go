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

type FakeInfo struct {
	Name   string `crorm:"name__"`
	Age    int    `crorm:"age__"`
	Gender string `crorm:"gender__"`
	Hello  string `crorm:"hello__"`
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

	fakeInfo := FakeInfo{}

	myTable.Find(&infoFind)

	_, err := myTable.First(&fakeInfo)

	log.Println(infoFind)

	log.Println(fakeInfo)

	log.Println(err)
}
