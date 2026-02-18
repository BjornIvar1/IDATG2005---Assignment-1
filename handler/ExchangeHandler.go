package handler

import (
	"IDATG2005---Assignment-1/structs"
	"IDATG2005---Assignment-1/utils"
	"encoding/json"
	"net/http"
)

// MapToExchangeOutput
// Maps the response from the APIs to the ExchangeOutput struct
func mapToExchangeOutput(c structs.ExchangeCountry, e structs.ExchangeResponse, b []string) structs.ExchangeOutput {
	out := structs.ExchangeOutput{
		Country:  c.Name.Common,
		Basecode: e.BaseCode,
		Rates:    nil,
	}

	if len(b) > 0 {
		out.Rates = make(map[string]float64)
		for _, borderCurrCode := range b {
			if rate, exists := e.ExchangeRate[borderCurrCode]; exists {
				out.Rates[borderCurrCode] = rate
			}
		}
	}
	return out
}

// decodeJSON
// Decodes the JSON response from the API into the target struct
// If there is an error during decoding, it returns an error message with status code 500
func decodeJSON(w http.ResponseWriter, r *http.Response, target any, errMsg string) bool {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		http.Error(w, errMsg, http.StatusInternalServerError)
		return false
	}

	return true
}

// fetchJSON
// Fetches the JSON response from the API and checks for errors
// If there is an error during fetching, it returns an error message with status code 500
// If the status code of the response is not 200, it returns an error message with status code 404
func fetchJSON(w http.ResponseWriter, url, fetchErrMsg, statusErrMsg string) (*http.Response, bool) {
	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, fetchErrMsg, http.StatusInternalServerError)
		return nil, false
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		http.Error(w, statusErrMsg, http.StatusNotFound)
		return nil, false
	}

	return resp, true
}

// ExchangeHandler
// Takes the country code as a path parameter, and makes GET requests to the APIs,
// and returns the exchange rate information of the country and its borders.
// If there is an error it will return an error message with the appropriate status code
func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	conCode := r.PathValue("code")

	// Fetch the Country API
	conResp, ok := fetchJSON(w, utils.RestCountriesApiUrlBase+conCode, "Failed to fetch country information", "Country code is not valid")
	if !ok {
		return
	}

	var exchangeCountry []structs.ExchangeCountry
	if !decodeJSON(w, conResp, &exchangeCountry, "Failed to retrieve country information") {
		return
	}

	// Get the currency code of the country
	var currCode string
	for key := range exchangeCountry[0].Currencies {
		currCode = key
		break
	}

	// Fetch the Currency API
	curResp, ok := fetchJSON(w, utils.RestCurrenciesApiUrlBase+currCode, "Failed to fetch currencies information", "Currencie code is not valid")
	if !ok {
		return
	}

	var exchangeResponse structs.ExchangeResponse
	if !decodeJSON(w, curResp, &exchangeResponse, "Failed to retrieve currency information") {
		return
	}

	// Get Borders and currency code of the country and stores it in a list
	var borders []string
	for _, borderCode := range exchangeCountry[0].Borders {
		borResp, ok := fetchJSON(w, utils.RestCountriesApiUrlBase+borderCode,
			"Failed to fetch border country information", "Country code is not valid")
		if !ok {
			return
		}

		var borderCountry []structs.ExchangeCountry
		if ok = decodeJSON(w, borResp, &borderCountry, "Failed to retrieve currency information"); !ok {
			return
		}

		for borderCurrCode := range borderCountry[0].Currencies {
			borders = append(borders, borderCurrCode)
			break
		}
	}

	//My API-Response
	output := mapToExchangeOutput(exchangeCountry[0], exchangeResponse, borders)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
