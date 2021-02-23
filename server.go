package main

import (
	"github.com/husobee/vestigo"
	"log"
	"net/http"
	"specgen-test-go/spec"
)

func main() {
	router := vestigo.NewRouter()

	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*", "*"},
	})

	spec.AddEchoRoutes(router, &spec.EchoService{})
	spec.AddCheckRoutes(router, &spec.CheckService{})

	log.Fatal(http.ListenAndServe(":8081", router))
}