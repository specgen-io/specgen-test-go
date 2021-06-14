package spec

type EchoService struct{}

func (service *EchoService) EchoBody(body *Message) (*EchoBodyResponse, error) {
	return &EchoBodyResponse{Ok: body}, nil
}

func (service *EchoService) EchoQuery(intQuery int, stringQuery string) (*EchoQueryResponse, error) {
	return &EchoQueryResponse{Ok: &Message{IntField: intQuery, StringField: stringQuery}}, nil
}

func (service *EchoService) EchoHeader(intHeader int, stringHeader string) (*EchoHeaderResponse, error) {
	return &EchoHeaderResponse{Ok: &Message{IntField: intHeader, StringField: stringHeader}}, nil
}

func (service *EchoService) EchoUrlParams(intUrl int, stringUrl string) (*EchoUrlParamsResponse, error) {
	return &EchoUrlParamsResponse{Ok: &Message{IntField: intUrl, StringField: stringUrl}}, nil
}