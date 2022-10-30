package request

type URL struct {
	Schema string
	Host   string
	Port   int
	Query  Query
	Ref    string
}
