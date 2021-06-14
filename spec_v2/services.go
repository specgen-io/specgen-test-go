package spec_v2

type EmptyDef struct{}

var Empty = EmptyDef{}

type EchoBodyResponse struct {
	Ok *Message
}

type IEchoService interface {
	EchoBody(body *Message) (*EchoBodyResponse, error)
}
