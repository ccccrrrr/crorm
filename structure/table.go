package structure

import (
	"database/sql"
	"log"
)

func (db *DB) Table(tableName string) *Table {
	table := Table{}
	table.TableName = tableName
	table.DBInfo = db
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
		table.ColumnType = append(table.ColumnType, Type.String)
		table.ColumnName = append(table.ColumnName, Field.String)
		table.ColumnGoType = append(table.ColumnGoType, MapSqlToGoType[Type.String])
	}
	//log.Println(table.ColumnType)
	//log.Println(table.ColumnName)
	//log.Println(table.ColumnGoType)
	return &table
}

func (table *Table) table() *Table {
	return table
}

func (table *Table) GenerateTableQueryAndArgs() (string, []interface{}) {
	query := "SELECT * FROM `" + table.TableName + "` "
	if table.query != "" {
		query += "WHERE " + table.query + ";"
	}
	return query, table.args
}

func (table *Table) addQuery(query string, args ...interface{}) *Table {
	table.query = query
	table.args = args
	return table
}

func generateGetTableInfoQuery(db *DB, tableName string) string {
	query := "show columns from `" + tableName + "`;"
	return query
}
