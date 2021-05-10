package structure

type list struct {
	context []interface{}
}

func (table *Table) Where(query string, args ...interface{}) *Table {
	return table.clone().addQuery(query, args...).table()

	//mysqlQuery := generateWhereQuery(table.TableName, query)
	//rows, err := table.DBInfo.DefaultDB.Query(mysqlQuery, args...)
	//if err != nil {
	//	log.Println(err)
	//	return nil
	//}
	//columns, _ := rows.Columns()
	//columnLength := len(columns)
	//log.Println(columns)
	//res := make([]interface{}, columnLength)
	//for index, _ := range res {
	//	if table.ColumnType[index] == "varchar(64)" {
	//		var a string
	//		res[index] = &a
	//	} else {
	//		var a int
	//		res[index] = &a
	//	}
	//}
	//var _list []interface{}
	//for rows.Next() {
	//	err = rows.Scan(res...)
	//	if err != nil {
	//		log.Println(err)
	//		return nil
	//	}
	//	for i, _ := range res {
	//		if table.ColumnType[i] == "varchar(64)" {
	//			fmt.Print(*res[i].(*string))
	//		} else {
	//			fmt.Print(*res[i].(*int))
	//		}
	//		//fmt.Print(reflect.TypeOf(res[i]).Kind())
	//	}
	//	fmt.Printf("\n")
	//	//fmt.Println(res)
	//	//var item []string
	//	//for i, data := range res {
	//	//	if table.ColumnType[i] == "varchar(64)" {
	//	//		item = append(item, *data.(*string))
	//	//	}else {
	//	//		item = append(item, strconv.Itoa(*data.(*int)))
	//	//	}
	//	//}
	//	//_list = append(_list, item)
	//}
	////log.Printf("where: %v", _list)
	//return &list{context: _list}
}

func generateWhereQuery(tableName string, query string) string {
	mysqlQuery := "SELECT * FROM `" + tableName + "` WHERE " + query + ";"
	return mysqlQuery
}
