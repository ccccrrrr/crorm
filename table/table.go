package table

type Table struct {
	DBInfo       *DB
	TableName    string
	ColumnType   []string
	ColumnName   []string
	ColumnGoType []string
	Model        *interface{}
	Search       *Search
	Exec         *Exec
}

func (table *Table) clone() *Table {
	newTable := Table{
		table.DBInfo,
		table.TableName,
		table.ColumnType,
		table.ColumnName,
		table.ColumnGoType,
		table.Model,
		table.Search,
		table.Exec,
	}
	return &newTable
}

func (table *Table) table() *Table {
	return table
}

//func (table *Table) GenerateTableQueryAndArgs() (string, []interface{}) {
//	query := "SELECT * FROM `" + table.TableName + "` "
//	if table.query != "" {
//		query += "WHERE " + table.query + ";"
//	}
//	return query, table.args
//}
//
//func (table *Table) WhereQuery(query string, args ...interface{}) *Table {
//	table.query = query
//	table.args = args
//	return table
//}
