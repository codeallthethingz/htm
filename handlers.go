package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes defines all the routs that this server will handle.
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/learnings/{image}", LearningsHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(HomeHandler)
}

type transfer struct {
	SpatialPooler    *SpatialPooler
	Image            string
	Encoded          string
	Threshold        int
	Overlap          int
	Active           bool
	InputSpaceWidth  int
	InputSpaceHeight int
}

var spatialPooler *SpatialPooler
var eyeInputNeurons []*Neuron

// LearningsHandler returns a json rep of spatialpooler.
func LearningsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if paramsNotOk(w, params) {
		return
	}
	image := Images[params["image"]]
	if eyeInputNeurons == nil {
		eyeInputNeurons = MakeInputNeurons(19, 11)
	}
	Encode(image, eyeInputNeurons, 0.04, 19)
	if spatialPooler == nil {
		spatialPooler = NewSpatialPooler(10, 40, eyeInputNeurons)
	}
	threshold := 5
	overlap := 4
	spatialPooler.Activate(threshold, overlap, true)
	json, _ := json.Marshal(&transfer{
		SpatialPooler:    spatialPooler,
		Image:            image,
		Encoded:          InputNeuronsToString(eyeInputNeurons),
		Overlap:          overlap,
		Threshold:        threshold,
		InputSpaceWidth:  19,
		InputSpaceHeight: 11,
	})
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

// HomeHandler creates an index page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("have a look at handlers.go for what you can do"))
}
