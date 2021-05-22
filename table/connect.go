package table

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Open(config DBConfig) (*DB, error) {
	dataSourceName := config.UserName + ":" + config.UserPassword + "@/" + config.DBName
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("database connection fail")
		return nil, err
	}
	crormDB := &DB{
		DefaultDB: db,
		DBName:    config.DBName,
	}
	return crormDB, nil
}
