package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type transfer struct {
	SpatialPooler    *SpatialPooler `json:"spatialPooler"`
	Image            string         `json:"image"`
	Encoded          string         `json:"encoded"`
	Threshold        int            `json:"threshold"`
	ProximalOverlap  int            `json:"proximalOverlap"`
	DistalOverlap    int            `json:"distalOverlap"`
	Active           bool           `json:"active"`
	InputSpaceWidth  int            `json:"inputSpaceWidth"`
	InputSpaceHeight int            `json:"inputSpaceHeight"`
}

// SetupRoutes defines all the routs that this server will handle.
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/learnings/{image}", LearningsHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(HomeHandler)
}

var spatialPooler *SpatialPooler
var noteImages []string
var notes []*Neuron
var noteIndex = 0

// LearningsHandler returns a json rep of spatialpooler.
func LearningsHandler(w http.ResponseWriter, r *http.Request) {
	if notes == nil {
		setupNotes()
	}
	spSize := 100
	if spatialPooler == nil {
		spatialPooler = NewSpatialPooler(4, spSize, 0.4, notes)
	}
	threshold := 5
	proximalOverlap := int(float64(len(notes)) * 0.07)
	distalOverlap := int(float64(len(spatialPooler.getAllNeurons())) * 0.009)
	makeInputNote(noteIndex % len(noteImages))
	start := time.Now()
	spatialPooler.Activate(threshold, proximalOverlap, distalOverlap, true)
	fmt.Printf("activate took: %d\n", time.Since(start)/1000000)
	start = time.Now()
	json, _ := json.Marshal(&transfer{
		SpatialPooler:    spatialPooler,
		Image:            noteImages[noteIndex%len(noteImages)],
		Encoded:          InputNeuronsToString(notes),
		ProximalOverlap:  proximalOverlap,
		DistalOverlap:    distalOverlap,
		Threshold:        threshold,
		InputSpaceWidth:  10,
		InputSpaceHeight: 10,
	})

	fmt.Printf("marshal took: %d\n", time.Since(start)/1000000)
	noteIndex++
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

func setupNotes() {
	noteImages = make([]string, 5)
	notes = MakeInputNeurons(10, 10)
	noteImages[0] = "" +
		"     X    " +
		"     X    " +
		"    X X   " +
		"    X X   " +
		"   X   X  " +
		"   X   X  " +
		"  XXXXXXX " +
		"  X     X " +
		" X       X" +
		" X       X"

	noteImages[1] = "" +
		" XXXXXX   " +
		" X     X  " +
		" X     X  " +
		" X     X  " +
		" XXXXXX   " +
		" X     X  " +
		" X     X  " +
		" X     X  " +
		" X     X  " +
		" XXXXXX   "
	noteImages[2] = "" +
		"   XXXXX  " +
		"  X     X " +
		" X        " +
		" X        " +
		" X        " +
		" X        " +
		" X        " +
		" X        " +
		"  X     X " +
		"   XXXXX  "

	noteImages[3] = "" +
		" XXXXXX   " +
		" X     X  " +
		" X     X  " +
		" X      X " +
		" X      X " +
		" X      X " +
		" X      X " +
		" X     X  " +
		" X     X  " +
		" XXXXXX   "

	noteImages[4] = "" +
		" XXXXXXXX " +
		" X        " +
		" X        " +
		" X        " +
		" XXXXXX   " +
		" X        " +
		" X        " +
		" X        " +
		" X        " +
		" XXXXXXXX "
}

func makeInputNote(index int) {
	for _, k := range notes {
		k.Active = false
	}
	for i := index * 20; i < index*20+20; i++ {
		notes[i].Active = true
	}
}
