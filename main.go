package main

import (
	config "assignment_1/config"
	handler "assignment_1/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("$PORT has not been set, deafulting to: %s", config.DEFAULT_PORT)
		port = config.DEFAULT_PORT
	}
	http.HandleFunc(config.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(config.DIAGNOSTICS_PATH, handler.DiagHandler)
	http.HandleFunc(config.COUNTRIES_PATH, handler.CountryHandler)
	http.HandleFunc(config.UNI_INFO_PATH, handler.UniversityHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
