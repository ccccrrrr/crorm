package table

import (
	"crorm/typeMap"
	"database/sql"
	"log"
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
	for rows.Next() {
		var Field, Type, Default, Key, Extra, Null sql.NullString
		if _err := rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra); _err != nil {
			log.Println(_err)
			return nil
		}
		_table.ColumnType = append(_table.ColumnType, Type.String)
		_table.ColumnName = append(_table.ColumnName, Field.String)
		_table.ColumnGoType = append(_table.ColumnGoType, typeMap.MapSqlToGoType[Type.String])
	}
	return &_table
}

func generateGetTableInfoQuery(db *DB, tableName string) string {
	query := "show columns from `" + tableName + "`;"
	return query
}