package web

import (
	"fmt"
	"net/http"
)

// HandleFunc is handler to return all events
func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
