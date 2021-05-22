package table

func (table *Table) Find(receiver interface{}) *Table {

	return table.clone().Exec.Find(receiver).Table
	//
	//query, args := table.GenerateTableQueryAndArgs()
	//rows, err := table.DBInfo.DefaultDB.Query(query, args...)
	//
	//
	//if err != nil {
	//	log.Println(err)
	//	return table
	//}
	//
	//tableColumnNames := table.ColumnName
	//tableColumnTypes := table.ColumnType
	//
	//receiverName := getReceiverName(receiver)
	//receiverType := getReceiverType(receiver)
	//log.Println(tableColumnNames)
	//log.Println(receiverName)
	//columnLength := len(tableColumnTypes)
	//
	////ptr := 0
	//
	//log.Println(tableColumnNames)
	//log.Println(receiverName)
	//ptr := 0
	//for rows.Next() {
	//	res := make([]interface{}, columnLength)
	//
	//	ptr ++
	//	for i := 0; i < columnLength; i++ {
	//		for j := 0; j < columnLength; j++ {
	//			if tableColumnNames[i] == receiverName[j] {
	//				if tableColumnTypes[i] != receiverType[j] {
	//					log.Println("[ Find error] name and type mismatch")
	//					return table
	//				}
	//				if tableColumnTypes[i] == "varchar(64)" {
	//					var A string
	//					res[i] = &A
	//				}else if tableColumnTypes[i] == "int" {
	//					var A int
	//					res[i] = &A
	//				}
	//			}
	//		}
	//	}
	//
	//	err = rows.Scan(res...)
	//
	//	if err != nil {
	//		log.Printf("[Find error] %v", err)
	//		return table
	//	}
	//
	//	newItem := reflect.New(reflect.TypeOf(receiver)) // struct
	//
	//	for i := 0; i < columnLength; i++ {
	//		log.Println(reflect.ValueOf(res[i]).Elem())
	//		reflect.ValueOf(newItem).Field(i).Set(reflect.ValueOf(res[i]).Elem())
	//	}
	//
	//	reflect.ValueOf(receiver).Set(reflect.Append(newItem))
	//}
	//return table
}
