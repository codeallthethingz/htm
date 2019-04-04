package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var spatialPooler *SpatialPooler
var threshold = 5
var overlap = 8

func main() {
	go trainMNIST()
	// scoreMNIST(spatialPooler)

	router := mux.NewRouter()
	SetupRoutes(router)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	log.WithField("port", port).Info("http server listening")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:1234"}, //you service is available and allowed for this base url
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})
	log.Fatal(http.ListenAndServe(":"+port, corsOpts.Handler(router)))
}

// func scoreMNIST(spatialPooler *SpatialPooler) {
// 	correct := 0
// 	incorrect := 0
// 	dataSet, err := ReadTestSet("./mnist")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for imageIndex := 0; imageIndex < 500; imageIndex++ {
// 		data := dataSet.Data[imageIndex].Image
// 		digit := dataSet.Data[imageIndex].Digit
// 		image := createImage(data)
// 		encoded := encode(image, dataSet.W)
// 		spatialPooler.Activate(encoded, threshold, overlap, false, "#")
// 		guess := spatialPooler.WhatIsIt(encoded)
// 		if guess == fmt.Sprintf("%d", digit) {
// 			fmt.Print("--> ")
// 			correct++
// 		} else {
// 			incorrect++
// 		}
// 		fmt.Printf("Actual: %d, Guess: %s\n", digit, guess)
// 	}
// 	fmt.Printf("Results, correct: %d, incorrect: %d, accuracy: %f\n", correct, incorrect, (float32(correct)/float32(incorrect+correct))*100)
// }

func trainMNIST() {
	spatialPooler = NewSpatialPooler(4096, 40, 28, 28)
	dataSet, _ := ReadTrainSet("./mnist")
	for i := 0; i < 10; i++ {
		fmt.Printf("learning %d ", i)
		currentImageIndex = 0
		for imageIndex := 0; imageIndex < 1000; imageIndex++ {
			image := GetDigit(dataSet, i)
			encoded := encode(image, dataSet.W)
			spatialPooler.Activate(encoded, threshold, overlap, true, fmt.Sprintf("%d", i))
			fmt.Print(".")
			spatialPooler.Deactivate()
		}
		fmt.Println()
	}

}

func encode(obj string, lineLength int) string {
	onBits, offBits := countBits(obj)
	totalBits := offBits + onBits
	target := int(float64(totalBits) * 0.04)
	return turnOffBits(obj, onBits, target, lineLength)
}

func turnOffBits(obj string, currentlyOn int, targetOn int, lineLength int) string {
	// fmt.Printf("encoding object starting length: %d, on: %d, target: %d\n", len(obj), currentlyOn, targetOn)
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

	// onBits, _ := countBits(newObj)
	// fmt.Printf("encoding object finishing length: %d, currentlyOn: %d\n", len(newObj), onBits)
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

func printEncoding(encoding string, width int) {
	for c := range encoding {
		if c%width == 0 {
			fmt.Print("\n")
		}
		fmt.Print(string(encoding[c]))
	}
	fmt.Print("\n")
}
