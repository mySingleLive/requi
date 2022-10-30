package request

type URL struct {
	Schema string
	Host   string
	Port   int
	Path   string
	Query  Query
	Ref    string
}
