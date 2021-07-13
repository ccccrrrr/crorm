# crorm

an imitation orm from gorm
crorm supports basic CRUD operations based on mysql
while not like gorm you operate mainly by gorm.DB, crorm do most operations in crorm.Table

## Open a database

you can configure your mysql connection info using DBConfig struct

```go

type DBConfig struct {
   UserName     string
   UserPassword string
   Port         string
   IP           string
   DBName       string
}

// after configuration, connect

db, err := Open(standardConfig)

```

## Select a table

In crorm, you can choose a existed table, create a new table or override an old one. all of the operation returns the table links, or returns the error if there is type mismatch or something wrong to get table information

```go
// pure create table operation followed by the type of standard struct
table, err := db.CreateTable(&standardStruct)

// if the table already exists(have the same name), crorm will just override it
table, err := db.CreateOrOverrideTable(&standardStruct)

// crorm will search for the existed table in the database, and returns table information if all the name and type matches(following a very strict rule)
table, err := db.Sync(&standardStruct)

```

## Insert


## Delete


## Find and First


## Update








