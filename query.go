package crorm

import (
	"log"
	"reflect"
)

func (exec *Exec) generateDeleteQuery() (string, []interface{}) {
	query := "DELETE FROM `" + exec.Table.TableName + "` "
	query1, args := exec.generateWhereQuery()
	query += query1
	return query, args
}

func (exec *Exec) generateWhereQuery() (string, []interface{}) {
	if len(exec.Table.Search.WhereQuery) == 0 {
		return "", nil
	}
 	return "WHERE " + exec.Table.Search.WhereQuery + " ", exec.Table.Search.WhereArgs
}

func (exec *Exec) generateInsertQuery(insertInfo *InsertInfo) (string, []interface{}) {
	query := "insert into `" + exec.Table.TableName + "` "
	for i := 0; i < len(exec.Table.ColumnName); i++ {
		if i == 0 {
			query += "( "
		}
		query += exec.Table.ColumnName[i]
		if i == len(exec.Table.ColumnName)-1 {
			query += ") "
		} else {
			query += ", "
		}
	}
	for i := 0; i < len(exec.Table.ColumnName); i++ {
		if i == 0 {
			query += "VALUES("
		}
		if exec.Table.ColumnType[i] == "varchar(64)" {
			query += "'" + reflect.ValueOf(insertInfo.InsertList[exec.Table.ColumnName[i]]).String() + "'"
		} else if exec.Table.ColumnType[i] == "int" {
			query += reflect.ValueOf(insertInfo.InsertList[exec.Table.ColumnName[i]]).String()
		}
		if i == len(exec.Table.ColumnName) - 1 {
			query += ");"
		} else {
			query += ", "
		}
	}
	return query, nil
}

func (exec *Exec) generateUpdateQuery() (string, []interface{}) {
	//query := "UPDATE `" + exec.Table.TableName + "` \n"
	//query += "SET " + exec.Table.updateSet + " \n"
	//query += "WHERE " + exec.Table.query + " ;\n"
	//newArgs := append(exec.Table.updateArgs, exec.Table.args...)
	//return query, newArgs
	return "", nil
}

func (exec *Exec) generateFirstQuery() (string, []interface{}) {
	query1 := "SELECT * FROM `" + exec.Table.TableName + "` "
	query, args := exec.generateWhereQuery()
	query1 += query
	log.Println(query1)
	return query1, args
}