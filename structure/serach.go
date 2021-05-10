package structure

type search struct {
	Db *DB
	TableName string
	WhereConditions []map[string]interface{}
	Limit int64
}

func (s *search) db() *DB {
	return s.Db
}

func (s *search) limit(limit int64) *search {
	s.Limit = limit
	return s
}

func (s *search) Do() *search {
	query, args :=  GenerateQueryAndArgs(s)
	s.Db.DefaultDB.Exec(query, args...)
	return s
}

func GenerateQueryAndArgs(s *search) (string, []interface{}) {

	return "", nil

}

func (s *search) Where(query interface{}, values ...interface{}) *search {
	s.WhereConditions = append(s.WhereConditions, map[string]interface{}{"query": query, "args": values})
	return s
}
func (s *search) clone() *search {
	newSearch := &search{
		s.Db,
		s.TableName,
		make([]map[string]interface{}, len(s.WhereConditions)),
		s.Limit,
	}
	for _, value := range s.WhereConditions {
		newSearch.WhereConditions = append(newSearch.WhereConditions, value)
	}
	return newSearch
}
