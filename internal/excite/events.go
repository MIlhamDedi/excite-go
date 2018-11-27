package excite

import (
	"fmt"
	"net/http"

	"github.com/milhamdedi/excite-go/model"
)

// Get all events
func getEvents(w http.ResponseWriter, r *http.Request) {

}

// Get specific event
func getEvent(w http.ResponseWriter, r *http.Request) {

}

// Create a New Event
func createEvents(w http.ResponseWriter, r *http.Request) {

}

// Update an Event
func updateEvents(w http.ResponseWriter, r *http.Request) {

}

// Delete an Event
func deleteEvents(w http.ResponseWriter, r *http.Request) {

}

// SaveAndReadCard test util
func SaveAndReadCard() {
	e1 := &model.Event{Name: "Testevent", Description: []byte("Event for new Go Developers.")}
	e1.SaveCard()
	e2, _ := model.ReadCard("Testevent")
	fmt.Println(string(e2.Description))
}
