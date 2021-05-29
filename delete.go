package crorm

func (table *Table) Delete() *Table {

	return table.clone().Exec.Delete().Table

	//query, args := generateDeleteQuery(table)
	//if element == nil {
	//	err := table.First(element)
	//
	//	if err != nil {
	//		log.Printf("[delete error] %v\n", err)
	//		return table
	//	}
	//
	//}
	//
	//_, err := table.DBInfo.DefaultDB.Exec(query, args...)
	//if err != nil {
	//	log.Printf("[delete error] %v\n", err)
	//}
	//return table
}

