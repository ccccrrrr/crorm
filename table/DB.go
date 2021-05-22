package table

import (
	"crorm/typeMap"
	"crorm/util"
	"database/sql"
	"errors"
	"log"
	"reflect"
)

// DBConfig is configuration to connect to certain database
type DBConfig struct {
	UserName     string
	UserPassword string
	Port         string
	Ip           string
	DBName       string
}

type DB struct {
	DefaultDB *sql.DB
	UserName  string
	DBName    string
	Port      string
	Ip        string
	Value     interface{}
}

func (db *DB) clone() *DB {

	newDB := &DB{
		DefaultDB: db.DefaultDB,
		UserName:  db.UserName,
		DBName:    db.DBName,
		Port:      db.Port,
		Ip:        db.Ip,
		Value:     db.Value,
	}

	return newDB
}



// Sync compares the existed table to the interface's config
// this name is just used in zorm...
func (db *DB) Sync(tableConfig interface{}) (*Table, error) {
	syncName := util.GetTableName(tableConfig)
	columnName, columnType := db.getTableConfig(syncName)

	if len(columnName) == 0 || len(columnType) == 0 || len(columnName) != len(columnType) {
		return nil, errors.New("column number mismatch")
	}

	syncTableConfig := generateTableConfig(tableConfig)

	if reflect.DeepEqual(columnName, syncTableConfig.ColumnName) && reflect.DeepEqual(columnType, syncTableConfig.ColumnType) {
		//table := db.Table(syncName)
		table := Table{}
		table.TableName = syncName
		table.DBInfo = db
		table.Exec = &Exec{}
		table.Search = NewSearch()
		table.Search.Table = &table
		table.Exec.Table = &table
		table.ColumnType = columnType
		table.ColumnName = columnName
		return &table, nil
	} else {
		err := errors.New("no match type to this struct")
		return nil, err
	}
}

func (db *DB) getTableConfig(tableName string) (ColumnName []string, ColumnType []string) {
	query := "desc `" + tableName + "`;"
	rows, err := db.DefaultDB.Query(query)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	var columnName []string
	var columnType []string
	for rows.Next() {
		var Field, Type, Null, Key, Default, Extra sql.NullString
		err = rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)
		if err != nil {
			log.Println(err)
			return nil, nil
		}
		columnName = append(columnName, Field.String)
		columnType = append(columnType, Type.String)
	}
	return columnName, columnType
}

// CreateOrOverrideTable function create the table or just delete the table with the same name
func (db *DB) CreateOrOverrideTable(tableConfig interface{}) (*Table, error) {
	// generate table name
	tableName := util.GetTableName(tableConfig)
	log.Println("create table name " + tableName)

	table, err := db.CreateTable(tableConfig)

	if err != nil {
		// this means that the database already has the same table name
		query := "drop tables `" + tableName + "`;"
		_, err := db.DefaultDB.Exec(query)
		if err != nil {
			return nil, err
		}
		table, err = db.CreateTable(tableConfig)
		if err != nil {
			log.Println(errors.New("create table and override"))
			return nil, err
		}
		return table, err
	}
	return table, nil
}

// this Table function is not useful
// will be removed
func (db *DB) Table(tableName string) *Table {
	_table := Table{}
	_table.TableName = tableName
	_table.DBInfo = db
	_table.Exec = &Exec{}
	_table.Search = NewSearch()
	_table.Search.Table = &_table
	_table.Exec.Table = &_table
	query := generateGetTableInfoQuery(db, tableName)
	rows, err := db.DefaultDB.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	flag := 0
	for rows.Next() {
		flag = 1
		var Field, Type, Default, Key, Extra, Null sql.NullString
		if _err := rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra); _err != nil {
			log.Println(_err)
			return nil
		}
		_table.ColumnType = append(_table.ColumnType, Type.String)
		_table.ColumnName = append(_table.ColumnName, Field.String)
		_table.ColumnGoType = append(_table.ColumnGoType, typeMap.MapSqlToGoType[Type.String])
	}
	if flag == 0 {
		log.Println(errors.New("no table found"))
		return nil
	}
	return &_table
}

func generateGetTableInfoQuery(db *DB, tableName string) string {
	query := "show columns from `" + tableName + "`;"
	return query
}