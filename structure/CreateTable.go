package structure

import (
	"fmt"
	"log"
	"reflect"
)

func (db *DB) CreateTable(tableName string, config interface{}) (*DB, error) {
	tableConfig := generateTableConfig(config)
	tableConfig.TableName = tableName
	query := generateCreateTableQuery(tableConfig)
//	log.Println("query: " + query)
	_, err := db.DefaultDB.Exec(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//tableLink := &Table{
	//	DBInfo: db,
	//	TableName: tableConfig.TableName,
	//	ColumnName: tableConfig.ColumnName,
	//	ColumnType: tableConfig.ColumnType,
	//}
	return db, nil
}

func generateTableConfig(defaultConfig interface{}) *TableConfig {
	tableConfig := &TableConfig{}
	tableConfig.TableName = fmt.Sprintf("%v", defaultConfig)
//	log.Println("tablename: " + tableConfig.TableName)
	_t := reflect.TypeOf(defaultConfig)
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
		_type = TypeMap[v.Elem().Field(i).Kind().String()]
		tableConfig.ColumnType = append(tableConfig.ColumnType, _type)
	}
	return tableConfig
}


func generateCreateTableQuery(tableConfig *TableConfig) string {
	query := "CREATE TABLE `" + tableConfig.TableName + "` (\n"
	//result, err := db.DefaultDB.Exec("CREATE TABLE `userinfo` (\n " +
	//	"`uid` int(10) NOT NULL AUTO_INCREMENT,\n  " +
	//	"`username` varchar(64) DEFAULT NULL,\n " +
	//	"`departname` varchar(64) DEFAULT NULL,\n " +
	//	"`created` date DEFAULT NULL,\n " +
	//	"PRIMARY KEY (`uid`)\n" +
	//	")")
	ColumnType := tableConfig.ColumnType
	ColumnName := tableConfig.ColumnName
	for i := 0; i < len(tableConfig.ColumnName); i++ {
		query += ColumnName[i] + " " + ColumnType[i] + " DEFAULT "
		if ColumnType[i] == "varchar(64)" {
			query += "' '"
		}else {
			query += "0 "
		}
		if i == len(tableConfig.ColumnName) - 1 {
			query += "\n)"
		}else {
			query += ",\n"
		}
	}
	return query
}
