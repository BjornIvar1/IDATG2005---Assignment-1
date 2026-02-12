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

// APIStatus
// Checks if an external API is up by making a GET request,
// and returns the HTTP status code.
// Returns 503 if the service is unavailable
func APIStatus(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		return 503 // Service Unavailable
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

// StatusHandler
// This handler checks the uptime of the service and the status of the external APIs,
// and returns this information in a JSON response.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	uptimeSeconds := int(time.Since(startTime).Seconds())
	currencyAPIStatus := APIStatus(utils.CurrenciesApiUrl)
	restCountriesAPIStatus := APIStatus(utils.RestCountriesApiUrl)

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
