package crorm

func (table *Table) Where(query string, args ...interface{}) *Table {
	if table == nil {
		return nil
	}
	return table.clone().Search.SetWhereQuery(query, args...).Table
}