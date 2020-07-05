package main

import (
	"fmt"
	"net/http"

	"github.com/HannesKimara/lighthouse/src"
	"github.com/gorilla/mux"
)

func main() {
	app := lighthouse.CreateApp()

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", app.HomeHandler).Methods("GET")

	addr := fmt.Sprint(app.Config.Host) + ":" + fmt.Sprint(app.Config.Port)

	fmt.Println("Server starting at", addr)
	http.ListenAndServe(addr, r)
}
