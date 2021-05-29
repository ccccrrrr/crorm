package crorm

func (table *Table) Update(updateSet string, updateArgs ...interface{}) *Table {

	return table.clone().Exec.Update(updateSet, updateArgs...).Table
}
