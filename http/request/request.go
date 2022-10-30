package request

type RequestType uint8

// Request Types
const (
	GET RequestType = iota
	POST
	PUT
	HEAD
	DELETE
	OPTIONS
	TRACE
	PATCH
)

func (t RequestType) Name() string {
	switch t {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case HEAD:
		return "HEAD"
	case DELETE:
		return "DELETE"
	case OPTIONS:
		return "OPTIONS"
	case TRACE:
		return "TRACE"
	case PATCH:
		return "PATCH"
	}
	return "unspecified"
}

type Request struct {
	Type    RequestType
	URL     string
	Headers []Header
	Body    Body
}

func New(typ RequestType) *Request {
	return &Request{
		Type: typ,
	}
}
