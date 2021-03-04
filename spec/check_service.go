package spec

type CheckService struct {}

func (service *CheckService) CheckQuery(
	pString string,
	pStringOpt *string,
	pStringArray []string,
	pDate string,
	pDateArray []string,
	pDatetime string,
	pInt int,
	pLong int64,
	pDecimal float64,
	pEnum Choice,
	pStringDefaulted string) *CheckQueryResponse {

	return &CheckQueryResponse{Ok: Empty()}
}
