package structure

var (
	TypeMap = map[string]string{
		"string": "varchar(64)",
		"int": "int",

	}

	MapSqlToGoType = map[string]string{
		"varchar(64)": "string",
		"int": "int",
	}
)
