package handler

import (
	"encoding/json"
	"net/http"
)

// InfoResponse represents the information from a country,
// using the REST countries API
type InfoResponse struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float32           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    []string          `json:"capital"`
}

// InfoHandler
// Takes the country code as a path parameter, and makes a GET request to the API,
// and returns the information from the country
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	apiURL := "http://129.241.150.113:8080/v3.1/alpha/" + code

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch country information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var countryInfo []InfoResponse
	json.NewDecoder(resp.Body).Decode(&countryInfo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countryInfo[0])
}
