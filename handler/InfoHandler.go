package handler

import (
	"IDATG2005---Assignment-1/structs"
	"encoding/json"
	"net/http"
)

// MapToCountryInfo
// Maps the response from the API to the CountryInfoOut struct
func MapToCountryInfo(c structs.RestCountryInfo) structs.CountryInfoOut {
	out := structs.CountryInfoOut{
		Name:       c.Name.Common,
		Continents: c.Continents,
		Population: c.Population,
		Area:       float64(c.Area),
		Languages:  c.Languages,
		Borders:    c.Borders,
		Flag:       c.Flag,
		Capital:    "",
	}
	if len(c.Capital) > 0 {
		out.Capital = c.Capital[0]
	}
	return out
}

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

	var restCountry []structs.RestCountryInfo
	errJsonDecoder := json.NewDecoder(resp.Body).Decode(&restCountry)

	if errJsonDecoder != nil {
		http.Error(w, "Failed to retrieve country information", http.StatusInternalServerError)
		return
	}

	output := MapToCountryInfo(restCountry[0])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
