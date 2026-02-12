package structs

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
