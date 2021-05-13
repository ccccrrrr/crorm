package table

func (table *Table) Update(updateSet string, updateArgs ...interface{}) *Table {

	return table.clone().Exec.Update(updateSet, updateArgs...).Table
	//
	//
	//table.updateSet = updateSet
	//table.updateArgs = updateArgs
	//
	//query, args := generateUpdateQuery(table)
	//
	//_, err := table.DBInfo.DefaultDB.Exec(query, args...)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//return table
}
