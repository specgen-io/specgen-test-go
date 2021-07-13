//go:generate specgen service-go --spec-file ./spec.yaml --generate-path .

package main

import (
	"flag"
	"fmt"
	"github.com/husobee/vestigo"
	"github.com/specgen-io/test-service-go/spec"
	"github.com/specgen-io/test-service-go/spec_v2"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8081", "port number")
	flag.Parse()

	router := vestigo.NewRouter()

	router.Get("/", func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*", "*"},
	})

	spec.AddEchoRoutes(router, &spec.EchoService{})
	spec.AddCheckRoutes(router, &spec.CheckService{})
	spec_v2.AddEchoRoutes(router, &spec_v2.EchoService{})

	fmt.Println("Starting service on port: "+*port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}