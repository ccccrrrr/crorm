package structure

import "database/sql"

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
	Search    *search
	Value     interface{}
}

type Table struct {
	DBInfo       *DB
	TableName    string
	ColumnType   []string
	ColumnName   []string
	ColumnGoType []string
	Model        interface{}
	query        string
	args         []interface{}
	updateSet    string
	updateArgs   []interface{}
}

func (table *Table) clone() *Table {
	newTable := Table{
		table.DBInfo,
		table.TableName,
		table.ColumnType,
		table.ColumnName,
		table.ColumnGoType,
		table.Model,
		table.query,
		table.args,
		table.updateSet,
		table.updateArgs,
	}
	return &newTable
}

type TableConfig struct {
	TableName    string
	ColumnType   []string
	ColumnName   []string
	ColumnGoType []string
}

type DataBaseOperation interface {
	CreateTable(tableName string, tableConfig *TableConfig) (*DB, error)
}

func (db *DB) clone() *DB {

	newDB := &DB{
		DefaultDB: db.DefaultDB,
		UserName:  db.UserName,
		DBName:    db.DBName,
		Port:      db.Port,
		Ip:        db.Ip,
		Search:    db.Search,
		Value:     db.Value,
	}

	return newDB
}
