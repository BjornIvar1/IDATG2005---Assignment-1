package structs

type ExchangeResponse struct {
	Country      string             `json:"country"`
	BaseCode     string             `json:"base_code"`
	ExchangeRate map[string]float64 `json:"rates"`
}

type ExchangeCountry struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Borders []string `json:"borders"`
}

type ExchangeOutput struct {
	Country  string             `json:"country"`
	Basecode string             `json:"base_code"`
	Rates    map[string]float64 `json:"exchange-rate"`
}
