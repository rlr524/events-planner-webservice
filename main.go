package main

import (
	. "./emiya"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/events", GetEvents).Methods("GET")
	r.HandleFunc("/api/events/{_id}", GetEvent).Methods("GET")
	r.HandleFunc("/api/events", CreateEvent).Methods("POST")
	r.HandleFunc("/api/events/{_id}", UpdateEvent).Methods("PUT")
	r.HandleFunc("/api/events/{_id}", DeleteEvent).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
