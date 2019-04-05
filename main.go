package main

import (
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
