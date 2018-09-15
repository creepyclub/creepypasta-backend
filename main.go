package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creepypasta home!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", MainHandler)
	log.Fatal(http.ListenAndServe(":9000", r))
}
