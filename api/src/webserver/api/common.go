package api

import (
	"log"
	"net/http"
)

func BadRequest(err error, w http.ResponseWriter) {
	log.Fatal(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request"))
}

func InternalServerError(err error, w http.ResponseWriter) {
	log.Fatal(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
