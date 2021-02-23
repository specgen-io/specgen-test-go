package spec

import (
	"encoding/json"
	"github.com/husobee/vestigo"
	"net/http"
	"strconv"
)

func AddEchoRoutes(router *vestigo.Router, echoService IEchoService) {
	router.Post("/echo/body", func (w http.ResponseWriter, r *http.Request) {
		var body Message
		json.NewDecoder(r.Body).Decode(&body)
		response := echoService.EchoBody(&body)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response.Ok)
	})
	router.Get("/echo/query", func (w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		intQuery, _ := strconv.Atoi(query["int_query"][0])
		stringQuery := query["string_query"][0]
		response := echoService.EchoQuery(intQuery, stringQuery)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response.Ok)
	})
	router.Get("/echo/header", func (w http.ResponseWriter, r *http.Request) {
		intHeader, _ := strconv.Atoi(r.Header.Get("Int-Header"))
		stringHeader := r.Header.Get("String-Header")
		response := echoService.EchoHeader(intHeader, stringHeader)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response.Ok)
	})
	router.SetCors("/echo/header", &vestigo.CorsAccessControl{
		AllowHeaders: []string{"Int-Header", "String-Header"},
	})
	router.Get("/echo/url_params/:int_url/:string_url", func (w http.ResponseWriter, r *http.Request) {
		intUrl, _ := strconv.Atoi(vestigo.Param(r, "int_url"))
		stringUrl := vestigo.Param(r, "string_url")
		response := echoService.EchoUrlParams(intUrl, stringUrl)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response.Ok)
	})
}

func AddCheckRoutes(router *vestigo.Router, checkService ICheckService) {
	router.Get("/check/query", func (w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		pString := query["p_string"][0]
		pStringOpt := &query["p_string_opt"][0]
		pStringArray := query["p_string_array"]
		pDate := query["p_date"][0]
		pDateArray := query["p_date_array"]
		pTime := query["p_time"][0]
		pDatetime := query["p_datetime"][0]
		pByte, _ := strconv.Atoi(query["p_byte"][0])
		pInt, _ := strconv.Atoi(query["p_int"][0])
		pLong, _ := strconv.ParseInt(query["p_long"][0], 10, 64)
		pDecimal, _ := strconv.ParseFloat(query["p_decimal"][0], 64)
		pChar := query["p_char"][0]
		pEnum := Choice(query["p_enum"][0])
		pStringDefaulted := query["p_string_defaulted"][0]
		response := checkService.CheckQuery(pString, pStringOpt, pStringArray, pDate, pDateArray, pTime, pDatetime, pByte, pInt, pLong, pDecimal, pChar, pEnum, pStringDefaulted)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response.Ok)
	})
}