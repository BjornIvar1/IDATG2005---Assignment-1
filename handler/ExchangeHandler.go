package handler

import (
	"IDATG2005---Assignment-1/structs"
	"IDATG2005---Assignment-1/utils"
	"encoding/json"
	"net/http"
)

// MapToExchangeOutput
// Maps the response from the APIs to the ExchangeOutput struct
func MapToExchangeOutput(c structs.ExchangeCountry, e structs.ExchangeResponse, b []string) structs.ExchangeOutput {
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

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	conCode := r.PathValue("code")

	//Check Country API
	apiURLCon := utils.RestCountriesApiUrlBase + conCode

	conResp, err := http.Get(apiURLCon)
	if err != nil {
		http.Error(w, "Failed to fetch country information", http.StatusInternalServerError)
		return
	}

	if conResp.StatusCode != http.StatusOK {
		http.Error(w, "Country code is not valid", http.StatusNotFound)
		return
	}

	defer conResp.Body.Close()

	var exchangeCountry []structs.ExchangeCountry
	conErrJsonDecoder := json.NewDecoder(conResp.Body).Decode(&exchangeCountry)

	if conErrJsonDecoder != nil {
		http.Error(w, "Failed to retrieve country information", http.StatusInternalServerError)
		return
	}

	// Get the currency code of the country
	var currCode string
	for key := range exchangeCountry[0].Currencies {
		currCode = key
		break
	}

	// Check Currency API
	apiURLCur := utils.RestCurrenciesApiUrlBase + currCode
	respCur, err := http.Get(apiURLCur)
	if err != nil {
		http.Error(w, "Failed to fetch currency information", http.StatusInternalServerError)
		return
	}
	defer respCur.Body.Close()

	if respCur.StatusCode != http.StatusOK {
		http.Error(w, "Exchange rate code is not valid", http.StatusNotFound)
		return
	}

	var exchangeResponse structs.ExchangeResponse
	curErrJsonDecoder := json.NewDecoder(respCur.Body).Decode(&exchangeResponse)

	if curErrJsonDecoder != nil {
		http.Error(w, "Failed to retrieve currency information", http.StatusInternalServerError)
		return
	}

	// Get Borders and currency code of the country
	var borders []string
	for _, borderCountry := range exchangeCountry[0].Borders {
		apiURLBor := utils.RestCountriesApiUrlBase + borderCountry
		respBor, err := http.Get(apiURLBor)

		if err != nil {
			http.Error(w, "Failed to fetch border country information", http.StatusInternalServerError)
			return
		}

		defer respBor.Body.Close()

		if respBor.StatusCode != http.StatusOK {
			http.Error(w, "Country code is not valid", http.StatusNotFound)
			return
		}

		var borderCountry []structs.ExchangeCountry
		borErrJsonDecoder := json.NewDecoder(respBor.Body).Decode(&borderCountry)

		if borErrJsonDecoder != nil {
			http.Error(w, "Failed to retrieve currency information", http.StatusInternalServerError)
			return
		}

		var borderCurrCode string
		for key := range borderCountry[0].Currencies {
			borderCurrCode = key
			break
		}

		borders = append(borders, borderCurrCode)
	}

	//My API
	output := MapToExchangeOutput(exchangeCountry[0], exchangeResponse, borders)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
