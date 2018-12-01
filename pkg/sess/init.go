package sess

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/milhamdedi/excite-go/model"
)

// RenderTemplate to render html page - tempoarily only for event card
func RenderTemplate(w http.ResponseWriter, tmpl string, c *model.Event) {
	t, err := template.ParseFiles("../../web/templates/" + tmpl + ".html")
	if err != nil {
		fmt.Fprintf(w, "<h1>404 Page Not Found</h1>")
		return
	}
	t.Execute(w, c)
}
