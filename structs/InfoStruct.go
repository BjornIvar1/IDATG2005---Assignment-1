package structs

// RestCountryInfo represents the information using the REST Countries API
type RestCountryInfo struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float32           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		PNG string `json:"png"`
	}
	Capital []string `json:"capital"`
}

// CountryInfoOut represents the information for my API response
type CountryInfoOut struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      string            `json:"flag"`
	Capital    string            `json:"capital"`
}
