package spec

import (
	"encoding/json"
	"fmt"
	"github.com/husobee/vestigo"
	"net/http"
)

func checkErrors(params *ParamsParser, w http.ResponseWriter) bool {
	if len(params.Errors) > 0 {
		w.WriteHeader(400)
		fmt.Println(params.Errors)
		return false
		}
	return true
}

func AddEchoRoutes(router *vestigo.Router, echoService IEchoService) {
	router.Post("/echo/body", func(w http.ResponseWriter, r *http.Request) {
		var body Message
		json.NewDecoder(r.Body).Decode(&body)
		response := echoService.EchoBody(&body)
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			return
		}
	})
	router.Get("/echo/query", func(w http.ResponseWriter, r *http.Request) {
		query := NewParamsParser(r.URL.Query())
		intQuery := query.Int("int_query")
		stringQuery := query.String("string_query")
		if !checkErrors(query, w) { return }
		response := echoService.EchoQuery(intQuery, stringQuery)
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			return
		}
	})
	router.Get("/echo/header", func(w http.ResponseWriter, r *http.Request) {
		header := NewParamsParser(r.Header)
		intHeader := header.Int("Int-Header")
		stringHeader := header.String("String-Header")
		if !checkErrors(header, w) { return }
		response := echoService.EchoHeader(intHeader, stringHeader)
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			return
		}
	})
	router.SetCors("/echo/header", &vestigo.CorsAccessControl{
		AllowHeaders: []string{"Int-Header, String-Header"},
	})
	router.Get("/echo/url_params/:int_url/:string_url", func(w http.ResponseWriter, r *http.Request) {
		query := NewParamsParser(r.URL.Query())
		intUrl := query.Int(":int_url")
		stringUrl := query.String(":string_url")
		if !checkErrors(query, w) { return }
		response := echoService.EchoUrlParams(intUrl, stringUrl)
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			return
		}
	})
}

func AddCheckRoutes(router *vestigo.Router, checkService ICheckService) {
	router.Get("/check/query", func(w http.ResponseWriter, r *http.Request) {
		query := NewParamsParser(r.URL.Query())
		pString := query.String("p_string")
		pStringOpt := query.StringNullable("p_string_opt")
		pStringArray := query.StringArray("p_string_array")
		pDate := query.Date("p_date")
		pDateArray := query.DateArray("p_date_array")
		pDatetime := query.DateTime("p_datetime")
		pInt := query.Int("p_int")
		pLong := query.Int64("p_long")
		pDecimal := query.Decimal("p_decimal")
		pEnum := Choice(query.StringEnum("p_enum", ChoiceValuesStrings))
		pStringDefaulted := query.StringDefaulted("p_string_defaulted", "the default value")
		if !checkErrors(query, w) { return }
		response := checkService.CheckQuery(pString, pStringOpt, pStringArray, pDate, pDateArray, pDatetime, pInt, pLong, pDecimal, pEnum, pStringDefaulted)
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			return
		}
	})
	router.Get("/check/forbidden", func(w http.ResponseWriter, r *http.Request) {
		response := checkService.CheckForbidden()
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			return
		}
		if response.Forbidden != nil {
			w.WriteHeader(403)
			return
		}
	})
}

