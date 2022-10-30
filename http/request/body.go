package request

type BodyType uint8

const (
	Form BodyType = iota
	Text
	JSON
	XML
	Multipart
	Binary
)

type Body interface {
	Type() BodyType
}
