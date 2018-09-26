package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mediocregopher/radix.v2/redis"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s endpoint!", r.URL.Path)
}

func main() {
	port := "8000"
	http.HandleFunc("/", handle)
	fmt.Println("Starting web server on port " + port)

	conn, dialErr := redis.Dial("tcp", "redis:6379")
	if dialErr != nil {
		log.Fatal(dialErr)
	}
	defer conn.Close()

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
