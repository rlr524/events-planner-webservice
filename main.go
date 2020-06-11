package main

import (
	. "./endpoints"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

var db *mongo.Database

const (
	COLLECTON_EV = "events"
	COLLECTION_ATT = "attendees"
)

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

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
