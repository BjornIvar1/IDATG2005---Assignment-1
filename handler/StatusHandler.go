package handler

import (
	"IDATG2005---Assignment-1/utils"
	"encoding/json"
	"net/http"
	"time"
)

/*
StatusResponse
This handler provides a simple status endpoint,
which returns the uptime of the service in seconds.
*/
type StatusResponse struct {
	RestCountriesAPI int    `json:"restcountriesapi"`
	CurrencyAPI      int    `json:"currenciesapi"`
	Version          string `json:"version"`
	UptimeSeconds    int    `json:"uptime"`
}

// Start the timer
var startTime = time.Now()

/*
IsAPIUp
Check if the APIs are up using http.Get function
and return the status code. If there is an error,
return 503 (Service Unavailable).
*/
func IsAPIUp(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		return 503 // Service Unavailable
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

/*
StatusHandler
This handler checks the uptime of the service and the status of the external APIs,
and returns this information in a JSON response.
*/
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	uptimeSeconds := int(time.Since(startTime).Seconds())
	currencyAPIStatus := IsAPIUp(utils.CurrenciesApiUrl)
	restCountriesAPIStatus := IsAPIUp(utils.RestCountriesApiUrl)

	resp := StatusResponse{
		RestCountriesAPI: restCountriesAPIStatus,
		CurrencyAPI:      currencyAPIStatus,
		Version:          "v1",
		UptimeSeconds:    uptimeSeconds,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
