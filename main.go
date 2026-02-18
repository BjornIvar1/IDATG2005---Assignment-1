package main

import (
	"IDATG2005---Assignment-1/handler"
	"IDATG2005---Assignment-1/utils"
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
	router.HandleFunc(utils.DefaultPath, handler.emptyHandler)
	router.HandleFunc(utils.StatusPath, handler.statusHandler)
	router.HandleFunc(utils.InfoPath, handler.InfoHandler)
	router.HandleFunc(utils.ExchangePath, handler.exchangeHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
