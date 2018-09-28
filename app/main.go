package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
)

// Index endpoint
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the base endpoint!\n")
}

// Hello endpoint
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	var srv http.Server
	http2.VerboseLogs = true
	srv.Addr = ":8000"
	srv.Handler = router

	http2.ConfigureServer(&srv, nil)

	log.Fatal(srv.ListenAndServeTLS("certs/dev.crt", "certs/dev.key"))
}
