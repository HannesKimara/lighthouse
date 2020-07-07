package lighthouse

import (
	"context"
	"log"
	"time"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	svgTemplate = loadTemplates("static/badge.svg")
)

// HomeHandler ...
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["message"] = "Hello from lighthouse ..."

	json.NewEncoder(w).Encode(data)
}

// AnalyticHandler ...
func (app *Application) AnalyticHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := r.URL.Query()
	collection := app.Db.Collection(vars["coll"])

	record := analyticRecord{getIPAddress(r), time.Now()}
	_, err := collection.InsertOne(context.TODO(), record)

	if err != nil {
		log.Fatal(err)
	}

	if _, ok := params["badge"]; ok{
		w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
		svgTemplate.Execute(w, BadgeSvgData{Count: "3.5M"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Application) StatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := app.Db.Collection(vars["coll"])

	opts := options.Count().SetMaxTime(200 * time.Millisecond)
	result, err := collection.CountDocuments(context.TODO(), bson.D{{}}, opts)
	
	if err !=nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(statModel{result})
}
