package structs

// ExchangeResponse represents the response for the Currencies API
type ExchangeResponse struct {
	BaseCode     string             `json:"base_code"`
	ExchangeRate map[string]float64 `json:"rates"`
}

// ExchangeCountry represents the response for the REST Countries API
type ExchangeCountry struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Borders []string `json:"borders"`
}

// ExchangeOutput represents the output for the exchange endpoint
type ExchangeOutput struct {
	Country  string             `json:"country"`
	Basecode string             `json:"base_code"`
	Rates    map[string]float64 `json:"exchange-rate"`
}
