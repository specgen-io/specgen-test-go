package spec_v2

type EchoService struct{}

func (service *EchoService) EchoBody(body *Message) (*EchoBodyResponse, error) {
	return &EchoBodyResponse{Ok: body}, nil
}