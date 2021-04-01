package spec_v2

type EchoService struct {}

func (service *EchoService) EchoBody(body *Message) *EchoBodyResponse {
	return &EchoBodyResponse{Ok: *body}
}