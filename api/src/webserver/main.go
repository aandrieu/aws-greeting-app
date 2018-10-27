package main

import (
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
	v1 "webserver/api/v1"
)

func main() {
	rand.Seed(time.Now().Unix())

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/greet", v1.HandleGreet).Methods("POST")

	log.Print("Starting api server on port 8000.")

	log.Fatal(http.ListenAndServe(":8000", r))
}
