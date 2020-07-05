package lighthouse

import (
	"net/http"
	"encoding/json"
)

// HomeHandler ...
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["message"] = "Hello from lighthouse ..."

	json.NewEncoder(w).Encode(data)
}
