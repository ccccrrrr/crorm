package crorm

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
