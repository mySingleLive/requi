package http

type URL struct {
	Schema string
	Host   string
	Port   int
	Query  Query
	Ref    string
}
