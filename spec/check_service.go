package spec

type CheckService struct {}

func (service *CheckService) CheckQuery(
	pString string,
	pStringOpt *string,
	pStringArray []string,
	pDate string,
	pDateArray []string,
	pTime string,
	pDatetime string,
	pByte int,
	pInt int,
	pLong int64,
	pDecimal float64,
	pChar string,
	pEnum Choice,
	pStringDefaulted string) *CheckQueryResponse {

	return &CheckQueryResponse{Ok: Empty()}
}
