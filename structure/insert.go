package structure

import (
	"fmt"
	"log"
	"reflect"
)

type InsertInfo struct {
	ColumnName []string
	ColumnType []string
	InsertList map[string]interface{}
}

func (table *Table) Insert(info interface{}) (*Table, error) {
	log.Println(table.ColumnName)
	insertInfo := InsertInfo{
		table.ColumnName,
		table.ColumnType,
		make(map[string]interface{}),
	}
	t := reflect.TypeOf(info)
	v := reflect.ValueOf(info)
	name := ""
	_type := ""
	for i := 0; i < t.Elem().NumField(); i++ {
		if name = v.Elem().Type().Field(i).Tag.Get("crorm"); name == "" {
			name =  v.Elem().Type().Field(i).Name
		}
		_type = TypeMap[v.Elem().Field(i).Kind().String()]
		value := fmt.Sprintf("%v", v.Elem().Field(i))
		log.Println("name : " + name + "  type: " + _type + " value: " + value)
		insertInfo.InsertList[name] = value
	}

	query, _ := generateInsertElementQuery(&insertInfo, table)

	log.Println("query: \n" + query)
	table.DBInfo.DefaultDB.Exec(query)

	return table, nil

}

func generateInsertElementQuery(insertInfo *InsertInfo, table *Table) (string, error) {
	query := ""
	query += "insert into `" + table.TableName + "` "
	log.Printf("len: %v\n",len(table.ColumnName))
	for i := 0; i < len(table.ColumnName); i++ {
		if i == 0 {
			query += "( "
		}
		query += table.ColumnName[i]
		if i == len(table.ColumnName) - 1 {
			query += ") "
		}else {
			query += ", "
		}
	}
	for i := 0; i < len(table.ColumnName); i++ {
		if i == 0 {
			query += "VALUES("
		}
		if table.ColumnType[i] == "varchar(64)" {
			query += "'" + reflect.ValueOf(insertInfo.InsertList[table.ColumnName[i]]).String() + "'"
		} else if table.ColumnType[i] == "int" {
			query += reflect.ValueOf(insertInfo.InsertList[table.ColumnName[i]]).String()
		}
		if i == len(table.ColumnName) - 1 {
			query += ");"
		}else {
			query += ", "
		}
	}
	return query, nil
}
