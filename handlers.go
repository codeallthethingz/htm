package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SetupRoutes defines all the routs that this server will handle.
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/learnings/reset", LearningsResetHandler).Methods("GET")
	router.HandleFunc("/learnings/{image}", LearningsHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(HomeHandler)
}

type transfer struct {
	SpatialPooler *SpatialPooler
	Image         string
	Encoded       string
	Threshold     int
	Overlap       int
	Active        bool
	Guess         string
}

var dataSet *DataSet

// LearningsResetHandler returns a json rep of spatialpooler.
func LearningsResetHandler(w http.ResponseWriter, r *http.Request) {
	currentImageIndex = 0
}

// LearningsHandler returns a json rep of spatialpooler.
func LearningsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if dataSet == nil {
		dataSet, _ = ReadTestSet("./mnist")
	}
	digit, _ := strconv.Atoi(params["image"])
	image := GetDigit(dataSet, digit)

	encoded := encode(image, 28)
	spatialPooler.Activate(encoded, threshold, overlap, false, "#")
	guess := spatialPooler.WhatIsIt(encoded)

	json, _ := json.Marshal(&transfer{
		SpatialPooler: spatialPooler,
		Image:         image,
		Encoded:       encoded,
		Overlap:       overlap,
		Threshold:     threshold,
		Guess:         guess,
	})
	spatialPooler.Deactivate()
	w.Write([]byte(json))
	currentImageIndex++
}

func createImage(image [][]uint8) string {
	result := ""
	for _, row := range image {
		for _, pix := range row {
			if pix == 0 {
				result += " "
			} else {
				result += "X"
			}
		}
	}
	return result
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
