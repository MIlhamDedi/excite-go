package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/milhamdedi/excite-go/features/events"
)

func main() {
	fmt.Println("Initializing server...")

	http.HandleFunc("/api/", events.HandleFunc)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
