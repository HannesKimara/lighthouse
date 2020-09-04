package lighthouse

import (
	"net/http"
	"log"
	"text/template"
)

func getIPAddress(r *http.Request) string {
	return r.RemoteAddr
}

func loadTemplates(path string) *template.Template {
	t, err := template.ParseFiles(path)

	if err != nil {
		log.Fatal(err)
	}
	return t
}