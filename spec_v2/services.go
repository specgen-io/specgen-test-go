package spec_v2

type EchoBodyResponse struct {
	Ok *Message
}

type IEchoService interface {
	EchoBody(body *Message) (*EchoBodyResponse, error)
}
