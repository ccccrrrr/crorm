package test

type DBConfig struct {
UserName     string
UserPassword string
Port         string
Ip           string
DBName       string
}

var Config = DBConfig {
	UserName:     "root",
	UserPassword: "lotus20001006",
	Port:         "3306",
	Ip:           "localhost",
	DBName:       "db1",
}
