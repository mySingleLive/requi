package request

// QueryParam of HTTP request
type QueryParam struct {
	Name  string // Name of query parameter
	Value string // Value of query parameter
}

func NewQuery(name string, value string) *QueryParam {
	return &QueryParam{
		Name:  name,
		Value: value,
	}
}

type Query struct {
	params []*QueryParam
}

func NewQueries() Query {
	return Query{
		params: []*QueryParam{},
	}
}

func (qs *Query) Add(q *QueryParam) *Query {
	qs.params = append(qs.params, q)
	return qs
}

func (qs *Query) AddQuery(name string, value string) *Query {
	qs.Add(NewQuery(name, value))
	return qs
}

func (qs *Query) GetQueries(name string) []*QueryParam {
	var rets []*QueryParam
	for i := range qs.params {
		q := qs.params[i]
		if q.Name == name {
			rets = append(rets, q)
		}
	}
	return rets
}

func (qs *Query) GetQuery(name string) *QueryParam {
	for i := range qs.params {
		q := qs.params[i]
		if q.Name == name {
			return q
		}
	}
	return nil
}
