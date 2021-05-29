package crorm

import (
	"fmt"
	"github.com/ccccrrrr/crorm/typeMap"
	"github.com/ccccrrrr/crorm/util"
	"log"
	"reflect"
)

type TableConfig struct {
	TableName    string
	ColumnType   []string
	ColumnName   []string
	ColumnGoType []string
}

func (db *DB) CreateTable(config interface{}) (*Table, error) {
	tableName := util.GetTableName(config)
	tableConfig := generateTableConfig(config)
	tableConfig.TableName = tableName

	query := generateCreateTableQuery(tableConfig)
	log.Println(query)
	_, err := db.DefaultDB.Exec(query)
	if err != nil {
		// log.Println(err)
		// if the table already exist, crorm will match the name and type of the table
		// if all of the details are matched, you can just load the table
		return db.Table(tableName), err
	}
	return db.Table(tableName), nil
}

func generateTableConfig(defaultConfig interface{}) *TableConfig {
	tableConfig := &TableConfig{}
	tableConfig.TableName = fmt.Sprintf("%v", defaultConfig)
	_t := reflect.TypeOf(defaultConfig)
	tableConfig.TableName = _t.Elem().String()
	if _t.Kind() != reflect.Ptr || _t.Elem().Kind() != reflect.Struct {
		log.Println("参数应该为结构体指针")
		log.Println(_t.Kind().String())
		log.Println(_t.Elem().Kind().String())
		return nil
	}
	t := reflect.TypeOf(defaultConfig)
	v := reflect.ValueOf(defaultConfig)
	name := ""
	_type := ""
	for i := 0; i < t.Elem().NumField(); i++ {
		tag := v.Elem().Type().Field(i).Tag
		if tag.Get("crorm") == "" {
			name =  v.Elem().Type().Field(i).Name
		}else {
			name = tag.Get("crorm")
		}
		tableConfig.ColumnName = append(tableConfig.ColumnName, name)
		tableConfig.ColumnGoType = append(tableConfig.ColumnGoType, v.Elem().Field(i).Kind().String())
		_type = typeMap.TypeMap[v.Elem().Field(i).Kind().String()]
		tableConfig.ColumnType = append(tableConfig.ColumnType, _type)
	}
	return tableConfig
}


func generateCreateTableQuery(tableConfig *TableConfig) string {
	query := "CREATE TABLE `" + tableConfig.TableName + "` (\n"
	ColumnType := tableConfig.ColumnType
	ColumnName := tableConfig.ColumnName
	log.Println(ColumnType)
	for i := 0; i < len(tableConfig.ColumnName); i++ {
		query += ColumnName[i] + " " + ColumnType[i] + " DEFAULT "
		if ColumnType[i] == "varchar(64)" {
			query += "' '"
		} else {
			query += "0 "
		}
		if i == len(tableConfig.ColumnName) - 1 {
			query += "\n)"
		} else {
			query += ",\n"
		}
	}
	return query
}
