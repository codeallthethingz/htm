package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes defines all the routs that this server will handle.
func SetupRoutes(router *mux.Router) {

	router.HandleFunc("/", HomeHandler).
		Methods("GET")
	router.HandleFunc("/spatialPooler", SpatialPoolerHandler).
		Methods("GET")
	router.HandleFunc("/activeForInput/{image}", ActiveSpatialPoolerForInputHandler).
		Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(HomeHandler)
}

type transfer struct {
	SpatialPooler *SpatialPooler
	Image         string
	Encoded       string
	Threshold     int
	Overlap       int
	Active        bool
}

// ActiveSpatialPoolerForInputHandler returns a json rep of spatialpooler.
func ActiveSpatialPoolerForInputHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if paramsNotOk(w, params) {
		return
	}
	image := Images[params["image"]]
	encoded := encode(image)
	threshold := 3
	overlap := 4
	spatialPooler.Activate(encoded, threshold, overlap)
	json, err := json.Marshal(&transfer{
		SpatialPooler: spatialPooler,
		Image:         image,
		Encoded:       encoded,
		Overlap:       overlap,
		Threshold:     threshold,
	})
	if err != nil {
		panic(err)
	}
	w.Write([]byte(json))
}

func paramsNotOk(w http.ResponseWriter, params map[string]string) bool {
	keys := ""
	for k := range Images {
		if params["image"] == k {
			return false
		}
		keys += k + ", "
	}
	response(w, 404, fmt.Sprintf("Unsupported image: %s supported types are %s", params["image"], keys))
	return true
}

func response(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Write([]byte(message))
}

// SpatialPoolerHandler returns a json rep of spatialpooler.
func SpatialPoolerHandler(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(&transfer{SpatialPooler: spatialPooler})
	if err != nil {
		panic(err)
	}
	w.Write([]byte(json))
}

// HomeHandler creates an index page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("have a look at handlers.go for what you can do"))
}
