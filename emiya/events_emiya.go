package emiya

import (
	. "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetEvents(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	_ = json.NewEncoder(w).Encode(Events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	for _, item := range Events {
		if item.ID == params["_id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Event{})
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var event Event
	_ = json.NewDecoder(r.Body).Decode(&event)
	Events = append(Events, event)
	_ = json.NewEncoder(w).Encode(event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Events {
		if item.ID == params["_id"] {
			Events = append(Events[:index], Events[index+1:]...)
			var event Event
			_ = json.NewDecoder(r.Body).Decode(&event)
			event.ID = params["_id"] // use the existing id
			Events = append(Events, event)
			_ = json.NewEncoder(w).Encode(event)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(Events)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Events {
		if item.ID == params["_id"] {
			Events = append(Events[:index], Events[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(Events)
}
