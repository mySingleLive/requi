package http

// Header of HTTP Req
type Header struct {
	Name  string // Name of Header
	Value string // Value of Header
}

func NewHeader(name string, value string) *Header {
	return &Header{
		Name:  name,
		Value: value,
	}
}

type Headers struct {
	list []*Header
}

func NewHeaders() Headers {
	hs := Headers{
		list: []*Header{},
	}
	return hs
}

func (hs *Headers) Add(h *Header) *Headers {
	hs.list = append(hs.list, h)
	return hs
}

func (hs *Headers) AddHeader(name string, value string) *Headers {
	h := NewHeader(name, value)
	return hs.Add(h)
}

func (hs *Headers) ALL() []*Header {
	return hs.list
}

func (hs *Headers) GetHeaders(name string) []*Header {
	var results []*Header
	for i := range hs.list {
		h := hs.list[i]
		if h.Name == name {
			results = append(results, h)
		}
	}
	return results
}

func (hs *Headers) GetHeader(name string) *Header {
	for i := range hs.list {
		h := hs.list[i]
		if h.Name == name {
			return h
		}
	}
	return nil
}
