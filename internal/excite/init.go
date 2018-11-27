package excite

import (
	"fmt"
	"net/http"
)

// Get all events
func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
