package spec_v2

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
	router.Post("/v2/echo/body", func(w http.ResponseWriter, r *http.Request) {
		var body Message
		json.NewDecoder(r.Body).Decode(&body)
		response := echoService.EchoBody(&body)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response.Ok)
	})
}

