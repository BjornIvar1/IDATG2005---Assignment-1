package handler

import (
	"encoding/json"
	"net/http"
)

type ExchangeResponse struct {
	BaseCurrency string             `json:"base_currency"`
	ExchangeRate map[string]float64 `json:"exchange-rate"`
}

type ExchangeOutput struct {
	Basecode string `json:"base_code"`
	Rates    struct {
		EUR float64 `json:"EUR"`
		SEK float64 `json:"SEK"`
		RUB float64 `json:"RUB"`
	}
}

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	apiURL := "http://129.241.150.113:9090/currency/" + code

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch country information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Country code is not valid", http.StatusNotFound)
		return
	}

	var exchangeOutput ExchangeOutput
	errJsonDecoder := json.NewDecoder(resp.Body).Decode(&exchangeOutput)

	if errJsonDecoder != nil {
		http.Error(w, "Failed to retrieve country information", http.StatusInternalServerError)
		return
	}

	output := exchangeOutput

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
