package main

import (
	"goodev/app/routes"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
)

func main() {
	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/hi/:name", routes.Hello)

	var srv http.Server
	// http2.VerboseLogs = true
	srv.Addr = ":8000"
	srv.Handler = router

	http2.ConfigureServer(&srv, nil)
	// log.Fatal(srv.ListenAndServeTLS("certs/dev.crt", "certs/dev.key"))
	log.Fatal(srv.ListenAndServe())
}
