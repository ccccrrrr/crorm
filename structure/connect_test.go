package structure

import (
	"log"
	"testing"
)

func TestOpen(t *testing.T) {
		config := DBConfig{
			UserName: "root",
			UserPassword: "lotus20001006",
			Port : "3306",
			Ip : "localhost",
			DBName: "db1",
		}
		db, err := Open(config)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("dbname: " + db.DBName)
		return
}

