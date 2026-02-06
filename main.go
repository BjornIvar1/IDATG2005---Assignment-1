package main

import (
	"IDATG2005---Assignment-1/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Instantiate new router
	router := http.NewServeMux()

	// Set up and attach handler endpoints to router
	router.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	//router.HandleFunc(handler.STATUS_PATH, handler.LocationHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
