package structs

// StatusResponse represents the status
type StatusResponse struct {
	RestCountriesAPI int    `json:"restcountriesapi"`
	CurrencyAPI      int    `json:"currenciesapi"`
	Version          string `json:"version"`
	UptimeSeconds    int    `json:"uptime"`
}
