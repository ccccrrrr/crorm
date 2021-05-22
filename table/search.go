package table

type Search struct {
	Table       *Table
	WhereQuery  string
	WhereArgs   []interface{}
	UpdateQuery string
	UpdateArgs  []interface{}
	FindQuery   string
	FindArgs    []interface{}
}

func NewSearch() *Search {
	return &Search{}
}

func (search *Search) SetWhereQuery(query string, args ...interface{}) *Search {
	search.WhereQuery = query
	search.WhereArgs = args
	return search
}