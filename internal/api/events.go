package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/milhamdedi/excite-go/model"
)

// getEvents all events
func getEvents(id string) any.Any {
	events := []model.Event{}
	toRead := []string{}
	if id != "" {
		toRead = append(toRead, id)
	} else {
		toRead = append(toRead, "codejam", "soc")
	}

	for _, id := range toRead {
		event, err := model.ReadCard(id)
		if err != nil {
			return events
		}
		events = append(events, *event)
	}
	return events
}

// Create a New Event
func createEvent(id string, r map[string]any.Any) {

}

// Update an Event
func updateEvents(w http.ResponseWriter, r *http.Request) {

}

// Delete an Event
func deleteEvents(w http.ResponseWriter, r *http.Request) {

}

// SaveAndReadCard test util
func SaveAndReadCard() {
	e1 := model.Event{Name: "Testevent", Description: []byte("Event for new Go Developers.")}
	e1.SaveCard()
	e2, _ := model.ReadCard("Testevent")
	fmt.Println(string(e2.Description))
}

// HandleEvents is API entrypoint
func HandleEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[len("/api/events"):]

	var response map[string]*any.Any

	switch r.Method {
	case "GET":
		response["events"] = getEvents(id)
	case "POST":
		createEvent(id, map[string]*any.Any)
	case "PUT":
		updateEvent(id, map[string]*any.Any)
	case "DELETE":
		deleteEvent(id)
	default:
		response["message"] = "Action not recognized"
	}

	json.NewEncoder(w).Encode(response)
}
