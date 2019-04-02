package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()
	SetupRoutes(router)

	generateCupLearnings(10)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8000"
	}
	log.WithField("port", port).Info("http server listening")
	//cors optionsGoes Below
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:1234"}, //you service is available and allowed for this base url
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})
	log.Fatal(http.ListenAndServe(":"+port, corsOpts.Handler(router)))

}

// CupLearnings a string array of json marshalled spatial pooler.
var CupLearnings string

func generateCupLearnings(iterate int) {
	CupLearnings += "["
	encoded := encode(cup)
	image := cup
	spatialPooler := NewSpatialPooler(100, 40, 19, 11)
	threshold := 5
	overlap := 4
	for i := 0; i < iterate; i++ {

		spatialPooler.Activate(encoded, threshold, overlap, false)
		if i > 0 {
			spatialPooler.Activate(encoded, threshold, overlap, true)
		}

		json, _ := json.Marshal(&transfer{
			SpatialPooler: spatialPooler,
			Image:         image,
			Encoded:       encoded,
			Overlap:       overlap,
			Threshold:     threshold,
		})
		CupLearnings += string(json)
		if i < iterate-1 {
			CupLearnings += ","
		}
	}
	CupLearnings += "]"
}

func encode(obj string) string {
	onBits, offBits := countBits(obj)
	totalBits := offBits + onBits
	target := int(float64(totalBits) * 0.04)
	return turnOffBits(obj, onBits, target)
}

func turnOffBits(obj string, currentlyOn int, targetOn int) string {
	lineLength := 19
	newObj := ""
	for c := range obj {
		if currentlyOn > targetOn && c > 0 && obj[c-1] == "X"[0] && obj[c] == "X"[0] && c < len(obj) && obj[c+1] == "X"[0] {
			newObj += " "
			currentlyOn--
		} else {
			newObj += string(obj[c])
		}
	}
	superNewObj := ""
	for c := range newObj {
		if currentlyOn > targetOn && newObj[c] == "X"[0] && c < len(newObj) && newObj[c+1] == "X"[0] {
			superNewObj += " "
			currentlyOn--
		} else {
			superNewObj += string(newObj[c])
		}

	}
	newObj = ""
	for c := range superNewObj {
		if currentlyOn > targetOn && c > lineLength && newObj[c-lineLength] == "X"[0] {
			newObj += " "
			currentlyOn--
		} else {
			newObj += string(superNewObj[c])
		}
	}
	return newObj
}

func countBits(obj string) (int, int) {
	onBits := 0
	offBits := 0
	for c := range obj {
		if obj[c] == "X"[0] {
			onBits++
		} else if obj[c] == " "[0] {
			offBits++
		}
	}
	return onBits, offBits
}

func printSpatialPooler(image string, spatialPooler *SpatialPooler) {
	for j := 0; j < 11; j++ {
		for i := 0; i < len(spatialPooler.Cells); i++ {
			fmt.Printf("cell: %d", i)
			for c := range image {
				if c%19 == 0 {
					fmt.Print("\n")
				}
				index, ok := spatialPooler.Cells[i].CoordLookup[c]
				if ok {
					fmt.Print(spatialPooler.Cells[i].Permanences[index])
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}

func printEncoding(encoding string) {

	for c := range encoding {
		if c%19 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(string(encoding[c]))
	}
	fmt.Print("\n")
}
