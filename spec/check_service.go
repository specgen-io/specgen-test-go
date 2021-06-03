package spec

import (
	"cloud.google.com/go/civil"
	"github.com/shopspring/decimal"
)

type CheckService struct {}

func (service *CheckService) CheckQuery(
	pString string,
	pStringOpt *string,
	pStringArray []string,
	pDate civil.Date,
	pDateArray []civil.Date,
	pDatetime civil.DateTime,
	pInt int,
	pLong int64,
	pDecimal decimal.Decimal,
	pEnum Choice,
	pStringDefaulted string) *CheckQueryResponse {

	return &CheckQueryResponse{Ok: Empty()}
}

func (service *CheckService) CheckForbidden() *CheckForbiddenResponse {
	return &CheckForbiddenResponse{Forbidden: Empty()}
}