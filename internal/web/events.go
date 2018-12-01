package web

import (
	"fmt"
	"net/http"

	"github.com/milhamdedi/excite-go/model"
	"github.com/milhamdedi/excite-go/pkg/sess"
)

// HandleEvents is handler to return all events
func HandleEvents(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/events/"):]
	method := r.URL.Query()["edit"]

	c, err := model.ReadCard(id)
	if err != nil {
		fmt.Fprintf(w, "<h1>404 Page Not Found</h1>")
		return
	}
	if len(method) != 0 && method[0] == "true" {
		sess.RenderTemplate(w, "events/edit", c)
	} else {
		sess.RenderTemplate(w, "events/view", c)
	}

}
