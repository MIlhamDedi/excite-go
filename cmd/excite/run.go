package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/milhamdedi/excite-go/internal/excite"
)

func main() {
	excite.SaveAndReadCard()

	http.HandleFunc("/api/", excite.HandleFunc)

	fmt.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
