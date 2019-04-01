package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes defines all the routs that this server will handle.
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).
		Methods("GET")
	router.HandleFunc("/spatialPooler", SpatialPoolerHandler).
		Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(HomeHandler)
}

// SpatialPoolerSerializable for sending over http
type SpatialPoolerSerializable struct{}

// SpatialPoolerHandler returns a json rep of spatialpooler.
func SpatialPoolerHandler(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(&SpatialPoolerSerializable{})
	if err != nil {
		panic(err)
	}
	w.Write([]byte(json))
}

// HomeHandler creates an index page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("have a look at handlers.go for what you can do"))
}
