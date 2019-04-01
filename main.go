package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

func main() {
	printEncoding(encode(Cup))
	everything(Cup)

	router := mux.NewRouter()
	SetupRoutes(router)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8000"
	}
	log.WithField("port", port).Info("http server listening")
	log.Fatal(http.ListenAndServe(":"+port, router))
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

// Map encoding space to input space.
// each of the pixels has a N% chance of being connected to every pixel in the spatial pooler
func everything(image string) {
	// sdr := encode(image)
	spatialPooler := NewSpatialPooler(10, 40, 220)
	printSpatialPooler(image, spatialPooler)
	// train(spatialPooler, sdr)
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
					fmt.Print(spatialPooler.Cells[i].Permenances[index])
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
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
	fmt.Println()
}
