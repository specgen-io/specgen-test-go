package main

import (
	"flag"
	"fmt"
	"github.com/husobee/vestigo"
	"log"
	"net/http"
	"specgen-test-service-go/spec"
)

func main() {
	port := flag.String("port", "8081", "port number")
	flag.Parse()

	router := vestigo.NewRouter()

	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*", "*"},
	})

	spec.AddEchoRoutes(router, &spec.EchoService{})
	spec.AddCheckRoutes(router, &spec.CheckService{})

	fmt.Println("Starting service on port: "+*port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}