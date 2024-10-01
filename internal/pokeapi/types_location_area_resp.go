package pokeapi

type LocationAreResp struct {
	Count    int     `json: "count, omitempty"`
	Next     *string `json: "next,omitempty"`
	Previous *string `json: "previous,omitempty"`
	Results  []struct {
		Name string `json: "name,omitempty"`
		Url  string `json: "url,omitempty"`
	} `json: "results,omitempty"`
}
