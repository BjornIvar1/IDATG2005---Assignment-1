package handler

import (
	"IDATG2005---Assignment-1/structs"
	"IDATG2005---Assignment-1/utils"
	"encoding/json"
	"net/http"
	"time"
)

// startTime Start the timer
var startTime = time.Now()

// statusAPI
// Checks if an external API is up by making a GET request,
// and returns the HTTP status code.
// Returns 503 if the service is unavailable
func statusAPI(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		return 503 // Service Unavailable
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

// statusHandler
// This handler checks the uptime of the service and the status of the external APIs,
// and returns this information in a JSON response.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	uptimeSeconds := int(time.Since(startTime).Seconds())
	currencyAPIStatus := statusAPI(utils.CurrenciesApiUrl)
	restCountriesAPIStatus := statusAPI(utils.RestCountriesApiUrl)

	resp := structs.StatusResponse{
		RestCountriesAPI: restCountriesAPIStatus,
		CurrencyAPI:      currencyAPIStatus,
		Version:          "v1",
		UptimeSeconds:    uptimeSeconds,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
