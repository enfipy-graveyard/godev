package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index endpoint
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the base endpoint!\n")
}

// Hello endpoint
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
