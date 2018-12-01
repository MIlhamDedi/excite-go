package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/milhamdedi/excite-go/internal/api"
	"github.com/milhamdedi/excite-go/internal/web"
)

// ApplicationI Interface
type ApplicationI interface {
	Start() error
}

// Application for Excite
type Application struct{}

// Start Excite Server
func (app *Application) Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../web/static"))))
	http.HandleFunc("/api/", api.HandleFunc)
	http.HandleFunc("/", web.HandleFunc)

	fmt.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
