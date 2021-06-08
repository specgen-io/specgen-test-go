package spec

type EchoService struct {}

func (service *EchoService) EchoBody(body *Message) *EchoBodyResponse {
	return &EchoBodyResponse{Ok: body}
}

func (service *EchoService) EchoQuery(intQuery int, stringQuery string) *EchoQueryResponse {
	return &EchoQueryResponse{Ok: &Message{IntField: intQuery, StringField: stringQuery}}
}

func (service *EchoService) EchoHeader(intHeader int, stringHeader string) *EchoHeaderResponse {
	return &EchoHeaderResponse{Ok: &Message{IntField: intHeader, StringField: stringHeader}}
}

func (service *EchoService) EchoUrlParams(intUrl int, stringUrl string) *EchoUrlParamsResponse {
	return &EchoUrlParamsResponse{Ok: &Message{IntField: intUrl, StringField: stringUrl}}
}