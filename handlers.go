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
	SpatialPooler *SpatialPooler
	Image         string
	Encoded       string
	Threshold     int
	Overlap       int
	Active        bool
}

var spatialPooler *SpatialPooler

var dataSet *DataSet
var currentImageIndex = 0

// LearningsHandler returns a json rep of spatialpooler.
func LearningsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if paramsNotOk(w, params) {
		return
	}
	if spatialPooler == nil {
		spatialPooler = NewSpatialPooler(1000, 40, 28, 28)
	}
	dataSet, err := ReadTrainSet("./mnist")
	if err != nil {
		fmt.Println(err)
		return
	}

	image := convertMNIST(dataSet)

	encoded := encode(image, 28)
	threshold := 4
	overlap := 15
	spatialPooler.Activate(encoded, threshold, overlap, true, "1")
	json, _ := json.Marshal(&transfer{
		SpatialPooler: spatialPooler,
		Image:         image,
		Encoded:       encoded,
		Overlap:       overlap,
		Threshold:     threshold,
	})
	w.Write([]byte(json))
	currentImageIndex++
}

func createImage(image [][]uint8) string {
	result := ""
	for _, row := range image {
		for _, pix := range row {
			if pix < 50 {
				result += " "
			} else {
				result += "X"
			}
		}
	}
	return result
}

func convertMNIST(dataSet *DataSet) string {
	var data [][]uint8
	digit := 0
	for digit != 1 {
		data = dataSet.Data[currentImageIndex].Image
		digit = dataSet.Data[currentImageIndex].Digit
		currentImageIndex++
	}
	return createImage(data)
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
