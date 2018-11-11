package main

import (
	"godev/src/routes"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/:name", routes.Hello)

	var srv http.Server
	srv.Addr = ":8000"
	srv.Handler = router

	log.Fatal(srv.ListenAndServe())
}
