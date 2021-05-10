package structure

import (
	"log"
)

func (table *Table) Update(updateSet string, updateArgs ...interface{}) *Table {
	table.updateSet = updateSet
	table.updateArgs = updateArgs

	query, args := generateUpdateQuery(table)

	log.Printf("query: %v\n", query)

	log.Printf("args: %v\n", args)

	_, err := table.DBInfo.DefaultDB.Exec(query, args...)

	if err != nil {
		log.Println(err)
	}

	return table
}
func generateUpdateQuery(table *Table) (string, []interface{}) {
	query := "UPDATE `" + table.TableName + "` \n"
	query += "SET " + table.updateSet + " \n"
	query += "WHERE " + table.query + " ;\n"
	newArgs := append(table.updateArgs, table.args...)
	return query,newArgs
}