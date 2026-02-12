package handler

import (
	"IDATG2005---Assignment-1/structs"
	"encoding/json"
	"net/http"
)

// InfoHandler
// Takes the country code as a path parameter, and makes a GET request to the API,
// and returns the information from the country
// If the country code is not valid, it returns an error message with status code 404
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	apiURL := "http://129.241.150.113:8080/v3.1/alpha/" + code

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch country information", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Country code is not valid", http.StatusNotFound)
		return
	}

	defer resp.Body.Close()

	var countryInfo []structs.InfoResponse
	errJsonDecoder := json.NewDecoder(resp.Body).Decode(&countryInfo)

	if errJsonDecoder != nil {
		http.Error(w, "Failed to retrieve country information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countryInfo[0])
}
