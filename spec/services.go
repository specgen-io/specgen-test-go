package spec

import "cloud.google.com/go/civil"
import "github.com/google/uuid"
import "github.com/shopspring/decimal"

type EmptyDef struct{}

var Empty = EmptyDef{}

type EchoBodyResponse struct {
	Ok *Message
}

type EchoQueryResponse struct {
	Ok *Message
}

type EchoHeaderResponse struct {
	Ok *Message
}

type EchoUrlParamsResponse struct {
	Ok *Message
}

type IEchoService interface {
	EchoBody(body *Message) (*EchoBodyResponse, error)
	EchoQuery(intQuery int, stringQuery string) (*EchoQueryResponse, error)
	EchoHeader(intHeader int, stringHeader string) (*EchoHeaderResponse, error)
	EchoUrlParams(intUrl int, stringUrl string) (*EchoUrlParamsResponse, error)
}

type CheckQueryResponse struct {
	Ok *EmptyDef
}

type CheckUrlParamsResponse struct {
	Ok *EmptyDef
}

type CheckForbiddenResponse struct {
	Ok *Message
	Forbidden *EmptyDef
}

type ICheckService interface {
	CheckQuery(pString string, pStringOpt *string, pStringArray []string, pDate civil.Date, pDateArray []civil.Date, pDatetime civil.DateTime, pInt int, pLong int64, pDecimal decimal.Decimal, pEnum Choice, pStringDefaulted string) (*CheckQueryResponse, error)
	CheckUrlParams(intUrl int64, stringUrl string, floatUrl float32, boolUrl bool, uuidUrl uuid.UUID, decimalUrl decimal.Decimal, dateUrl civil.Date) (*CheckUrlParamsResponse, error)
	CheckForbidden() (*CheckForbiddenResponse, error)
}
