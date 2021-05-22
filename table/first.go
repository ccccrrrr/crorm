package table

import (
	"crorm/typeMap"
	"errors"
	"log"
	"reflect"
)

func getReceiverName(receiver interface{}) []string {

	var receiverName []string
	_t := reflect.TypeOf(receiver)
	if _t.Kind() == reflect.Ptr && _t.Elem().Kind() == reflect.Slice {
		t := reflect.TypeOf(receiver).Elem()
		name := ""
		//_type := ""
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := t.Elem().Field(i).Tag
			if tag.Get("crorm") == "" {
				name = t.Elem().Field(i).Name
				log.Println(name)
			}else {
				name = tag.Get("crorm")
				log.Println(name)
			}
			receiverName = append(receiverName, name)
		}
		return receiverName
	}
	if _t.Elem().Kind() != reflect.Struct {
		log.Println("参数应该为结构体指针")
		return nil
	}
	if _t.Kind() == reflect.Ptr && _t.Elem().Kind() == reflect.Struct {
		t := reflect.TypeOf(receiver)
		name := ""
		//_type := ""
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := t.Elem().Field(i).Tag
			if tag.Get("crorm") == "" {
				name = t.Elem().Field(i).Name
			} else {
				name = tag.Get("crorm")
			}
			receiverName = append(receiverName, name)
		}
		return receiverName
	}
	return nil
}


func getReceiverType(receiver interface{}) []string {
	var receiverType []string
	_t := reflect.TypeOf(receiver)
	if _t.Kind() == reflect.Ptr && _t.Elem().Kind() == reflect.Slice {
		t := reflect.TypeOf(receiver).Elem()
		//name := ""
		_type := ""
		for i := 0; i < t.Elem().NumField(); i++ {
			_type = typeMap.TypeMap[t.Elem().Field(i).Type.String()]
			receiverType = append(receiverType, _type)
		}
		return receiverType
	}
	if _t.Kind() != reflect.Ptr || _t.Elem().Kind() != reflect.Struct {
		log.Println("参数应该为结构体指针")
		return nil
	}
	if _t.Kind() == reflect.Ptr && _t.Elem().Kind() == reflect.Struct {
		t := reflect.TypeOf(receiver)
		v := reflect.ValueOf(receiver)
		//name := ""
		_type := ""
		for i := 0; i < t.Elem().NumField(); i++ {
			_type = typeMap.TypeMap[v.Elem().Field(i).Kind().String()]
			receiverType = append(receiverType, _type)
		}
		return receiverType
	}
	return nil
}

func (table *Table) First(receiver interface{}) (*Table, error) {
	if table ==  nil {
		return nil, errors.New("no table")
	}
	return table.clone().Exec.First(receiver).Table, nil
	//
	//query, args := table.GenerateTableQueryAndArgs()
	//rows, err := table.DBInfo.DefaultDB.Query(query, args...)
	//
	//if err != nil {
	//	return err
	//}
	//
	//tableColumnNames := table.ColumnName
	//tableColumnTypes := table.ColumnType
	//
	//receiverName := getReceiverName(receiver)
	//receiverType := getReceiverType(receiver)
	//
	//columnLength := len(tableColumnTypes)
	//res := make([]interface{}, columnLength)
	//if rows.Next() {
	//	for i := 0; i < columnLength; i++ {
	//		for j := 0; j < columnLength; j++ {
	//			if tableColumnNames[i] == receiverName[j] {
	//				if tableColumnTypes[i] != receiverType[j] {
	//					return errors.New("name and type mismatch")
	//				}
	//				if tableColumnTypes[i] == "varchar(64)" {
	//					var a string
	//					res[i] = &a
	//				}else if tableColumnTypes[i] == "int" {
	//					var a int
	//					res[i] = &a
	//				}
	//			}
	//		}
	//	}
	//	err = rows.Scan(res...)
	//	if err != nil {
	//		return err
	//	}
	//
	//	for i := 0; i < columnLength; i++ {
	//		log.Println(reflect.TypeOf(receiver).Kind().String())
	//		if reflect.ValueOf(receiver).Elem().Field(i).Type() == reflect.ValueOf(res[i]).Elem().Type() {
	//			reflect.ValueOf(receiver).Elem().Field(i).Set(reflect.ValueOf(res[i]).Elem())
	//		}else {
	//			log.Println("no match type")
	//		}
	//	}
	//} else {
	//	return errors.New("no matching in this condition")
	//}
	//
	//return nil
}