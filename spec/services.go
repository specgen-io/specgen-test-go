package spec

import (
	"cloud.google.com/go/civil"
	"github.com/shopspring/decimal"
)

type EchoBodyResponse struct {
	Ok Message
}

type EchoQueryResponse struct {
	Ok Message
}

type EchoHeaderResponse struct {
	Ok Message
}

type EchoUrlParamsResponse struct {
	Ok Message
}

type IEchoService interface {
	EchoBody(body *Message) *EchoBodyResponse
	EchoQuery(intQuery int, stringQuery string) *EchoQueryResponse
	EchoHeader(intHeader int, stringHeader string) *EchoHeaderResponse
	EchoUrlParams(intUrl int, stringUrl string) *EchoUrlParamsResponse
}

type EmptyDef struct {}

func Empty() EmptyDef {
	return EmptyDef{}
}

type CheckQueryResponse struct {
	Ok interface{}
}

type CheckResponseForbiddenResponse struct {
	Ok interface{}
	Forbidden interface{}
}

type ICheckService interface {
	CheckQuery(pString string, pStringOpt *string, pStringArray []string, pDate civil.Date, pDateArray []civil.Date, pDatetime civil.DateTime, pInt int, pLong int64, pDecimal decimal.Decimal, pEnum Choice, pStringDefaulted string) *CheckQueryResponse
	CheckResponseForbidden() *CheckResponseForbiddenResponse
}