package api

import (
	"fmt"
	"net/http"
	"strings"
)

// HandleFunc is API entrypoint
func HandleFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/api/"):]
	fmt.Fprintf(w, "Hello, World. Hello, %s!", path)
	switch path[:strings.Index(path, "/")] {
	case "events":
		fmt.Println("one" + path[:strings.Index(path, "/")])
	case "organizations":
		fmt.Println("two" + path[:strings.Index(path, "/")])
	default:
		fmt.Println("three" + path[:strings.Index(path, "/")])
	}
}
