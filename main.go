package main

import (
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
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8000"
	}
	log.WithField("port", port).Info("http server listening")
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:1234"}, //you service is available and allowed for this base url
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})
	log.Fatal(http.ListenAndServe(":"+port, corsOpts.Handler(router)))
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
