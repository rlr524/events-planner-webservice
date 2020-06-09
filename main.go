package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Event struct {
	ID string `json:"_id"`
	Title string `json:"title"`
	Date string `json:"date"`
	Time string `json:"time"`
	Location string `json:"location"`
	Description string `json:"description"`
	Organizer string `json:"organizer"`
	Category string `json:"category"`
}

type Attendee struct {
	ID string `json:"_id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
}

var events []Event

func getEvents(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	_ = json.NewEncoder(w).Encode(events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	for _, item := range events {
		if item.ID == params["_id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Event{})
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var event Event
	_ = json.NewDecoder(r.Body).Decode(&event)
	events = append(events, event)
	_ = json.NewEncoder(w).Encode(event)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range events {
		if item.ID == params["_id"] {
			events = append(events[:index], events[index+1:]...)
			var event Event
			_ = json.NewDecoder(r.Body).Decode(&event)
			event.ID = params["_id"] // use the existing id
			events = append(events, event)
			_ = json.NewEncoder(w).Encode(event)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(events)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range events {
		if item.ID == params["_id"] {
			events = append(events[:index], events[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(events)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/events", getEvents).Methods("GET")
	r.HandleFunc("/api/events/{_id}", getEvent).Methods("GET")
	r.HandleFunc("/api/events", createEvent).Methods("POST")
	r.HandleFunc("/api/events/{_id}", updateEvent).Methods("PUT")
	r.HandleFunc("/api/events/{_id}", deleteEvent).Methods("DELETE")

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
