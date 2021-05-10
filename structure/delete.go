package structure

import (
	"log"
)

func (table *Table) Delete(element interface{}) *Table {

	query, args := generateDeleteQuery(table)
	if element == nil {
		err := table.First(element)

		if err != nil {
			log.Printf("[delete error] %v\n", err)
			return table
		}

	}

	_, err := table.DBInfo.DefaultDB.Exec(query, args...)
	if err != nil {
		log.Printf("[delete error] %v\n", err)
	}
	return table
}

func generateDeleteQuery(table *Table) (string, []interface{}) {
	query := "DELETE FROM `" + table.TableName + "` WHERE " + table.query
	return query, table.args
}