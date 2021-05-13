package table


func (table *Table) Insert(info interface{}) (*Table, error) {
	return table.clone().Exec.Insert(info).Table, nil
}
