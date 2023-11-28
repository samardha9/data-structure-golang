package types

// Color represents the color data structure as defined in data.json
type Color struct {
	Color    string `json:"color"`
	Category string `json:"category"`
	Type     string `json:"type,omitempty"`
	Code     struct {
		RGBA []int  `json:"rgba"`
		Hex  string `json:"hex"`
	} `json:"code"`
}

// Colors is a slice to hold multiple Color objects
type Colors []Color

// ColorData is a wrapper struct to match the JSON structure
type ColorData struct {
	Colors Colors `json:"colors"`
}

type StringData struct {
	String string `json:"string"`
}
