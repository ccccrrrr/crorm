package crorm

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type InsertInfo struct {
	ColumnName []string
	ColumnType []string
	InsertList map[string]interface{}
}

type Exec struct {
	Table *Table
}

//WhereQuery(query string, args ...interface{}) *Table
func (exec *Exec) Update(query string, args ...interface{}) *Exec {
	_query, _args := exec.generateUpdateQuery()
	exec.Table.DBInfo.DefaultDB.Exec(_query, _args...)
	return exec
}

func (exec *Exec) Insert(info interface{}) *Exec {

	insertInfo := InsertInfo{
		ColumnName: exec.Table.ColumnName,
		ColumnType: exec.Table.ColumnType,
		InsertList: make(map[string]interface{}),
	}

	t := reflect.TypeOf(info)
	v := reflect.ValueOf(info)

	name := ""
	//_type := ""

	for i := 0; i < t.Elem().NumField(); i++ {
		if name = v.Elem().Type().Field(i).Tag.Get("crorm"); name == "" {
			name = v.Elem().Type().Field(i).Name
		}
		//_type := typeMap.TypeMap[v.Elem().Field(i).Kind().String()]
		value := fmt.Sprintf("%v", v.Elem().Field(i))
		//log.Println("name : " + name + "  type: " + _type + " value: " + value)
		insertInfo.InsertList[name] = value
	}

	_query, _args := exec.generateInsertQuery(&insertInfo)

	_, err := exec.Table.DBInfo.DefaultDB.Exec(_query, _args...)

	if err != nil {
		log.Println(err)
	}

	return exec

}
func (exec *Exec) Delete() *Exec {
	query, args := exec.generateDeleteQuery()

	_, err := exec.Table.DBInfo.DefaultDB.Exec(query, args...)

	if err != nil {
		log.Println(err)
	}

	return exec

}
func (exec *Exec) Find(receiver interface{}) *Exec {

	query, args := exec.generateFirstQuery()
	//log.Println(query)
	//log.Println(args)
	rows, err := exec.Table.DBInfo.DefaultDB.Query(query, args...)

	if err != nil {
		log.Println(err)
		return exec
	}

	tableColumnNames := exec.Table.ColumnName
	tableColumnTypes := exec.Table.ColumnType

	receiverName := getReceiverName(receiver)
	receiverType := getReceiverType(receiver)
	//log.Println(tableColumnNames)
	//log.Println(receiverName)
	columnLength := len(tableColumnTypes)

	maxLen := reflect.ValueOf(receiver).Elem().Len()
	//log.Println(tableColumnNames)
	//log.Println(receiverName)

	ptr := -1
	for rows.Next() && ptr < maxLen - 1 {
		res := make([]interface{}, columnLength)
		ptr ++
		for i := 0; i < columnLength; i++ {
			for j := 0; j < columnLength; j++ {
				if tableColumnNames[i] == receiverName[j] {
					if tableColumnTypes[i] != receiverType[j] {
						log.Println("[ Find error] name and type mismatch")
						return exec
					}
					if tableColumnTypes[i] == "varchar(64)" {
						var A string
						res[i] = &A
					}else if tableColumnTypes[i] == "int" {
						var A int
						res[i] = &A
					}
				}
			}
		}

		err = rows.Scan(res...)

		if err != nil {
			log.Printf("[Find error] %v", err)
			return exec
		}

		// I cannot append new element to the slice
		// because reflect.ValueOf().Set method cannot assign new value to unaddressable value
		//log.Println(reflect.TypeOf(receiver).Kind())
		//log.Println(receiver)
		//
		//reflect.ValueOf(receiver).Set(reflect.Append(reflect.ValueOf(receiver).Elem(), reflect.New(reflect.TypeOf(receiver).Elem().Elem()).Elem()))

		for i := 0; i < columnLength; i++ {
			reflect.ValueOf(receiver).Elem().Index(ptr).Field(i).Set(reflect.ValueOf(res[i]).Elem())
		}
	}
	return exec
}

func (exec *Exec) First(receiver interface{}) *Exec {
	query, args := exec.generateFirstQuery()
	rows, err := exec.Table.DBInfo.DefaultDB.Query(query, args...)

	if err != nil {
		log.Println(err)
		return exec
	}

	tableColumnNames := exec.Table.ColumnName
	tableColumnTypes := exec.Table.ColumnType

	receiverName := getReceiverName(receiver)
	receiverType := getReceiverType(receiver)

	log.Println(receiverName)
	log.Println(receiverType)


	columnLength := len(tableColumnTypes)
	res := make([]interface{}, columnLength)

	if rows.Next() {
		for i := 0; i < columnLength; i++ {
			for j := 0; j < columnLength; j++ {
				if tableColumnNames[i] == receiverName[j] {
					if tableColumnTypes[i] != receiverType[j] {
						log.Println(errors.New("name and type mismatch"))
						return exec
					}
					if tableColumnTypes[i] == "varchar(64)" {
						var a string
						res[i] = &a
						continue
					}else if tableColumnTypes[i] == "int" {
						var a int
						res[i] = &a
						continue
					}
				}
				log.Println(errors.New("no correct name found"))
				return exec
			}
		}
		err = rows.Scan(res...)
		if err != nil {
			log.Println(err)
			return exec
		}

		for i := 0; i < columnLength; i++ {
			//log.Println(reflect.TypeOf(receiver).Kind().String())
			if reflect.ValueOf(receiver).Elem().Field(i).Type() == reflect.ValueOf(res[i]).Elem().Type() {
				reflect.ValueOf(receiver).Elem().Field(i).Set(reflect.ValueOf(res[i]).Elem())
			} else {
				log.Println("no match type")
			}
		}
	} else {
		log.Println(errors.New("no matching in this condition"))
		return exec
	}
	return exec
}
