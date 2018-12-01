package web

import (
	"fmt"
	"net/http"
	"strings"
)

// HandleFunc is handler to return all events
func HandleFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/"):]

	if strings.Index(path, "/") < 0 {
		HandleIndex(w, r)
		return
	}
	resource := path[:strings.Index(path, "/")]

	switch resource {
	case "events":
		HandleEvents(w, r)
	case "organizations":
		HandleEvents(w, r)
	}
}

// HandleIndex is handler to index page
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\nIt's the main page!")
}
