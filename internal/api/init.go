package api

import (
	"fmt"
	"net/http"
	"strings"
)

// HandleFunc is API entrypoint
func HandleFunc(w http.ResponseWriter, r *http.Request) {
	// to do: design better routing
	path := r.URL.Path[len("/api/"):]

	fmt.Fprintf(w, "Hello, %s!\n", path)
	if strings.Index(path, "/") < 0 {
		return
	}

	resource := path[:strings.Index(path, "/")]

	switch resource {
	case "events":
		fmt.Fprintf(w, "one "+path[:strings.Index(path, "/")]+r.Method)
	case "organizations":
		fmt.Fprintf(w, "two "+path[:strings.Index(path, "/")])
	default:
		fmt.Fprintf(w, "three "+path[:strings.Index(path, "/")])
	}
}
