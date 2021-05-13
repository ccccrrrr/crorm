package table

func (table *Table) Where(query string, args ...interface{}) *Table {
	return table.clone().Search.SetWhereQuery(query, args...).Table
}