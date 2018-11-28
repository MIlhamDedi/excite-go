package excite

import (
	"fmt"
	"log"
	"net/http"
)

// Application for Excite
type Application struct{}

// Start Excite Server
func (app *Application) Start() {
	http.HandleFunc("/api/", HandleFunc)
	fmt.Println("Listening on port 8000...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
