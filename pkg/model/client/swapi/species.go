package swapi

import "time"

type SearchSpeciesResponse struct {
	Count    int        `json:"count,omitempty"`
	Next     string     `json:"next,omitempty"`
	Previous string     `json:"previous,omitempty"`
	Results  []*Species `json:"results,omitempty"`
}

type Species struct {
	Name            string    `json:"name,omitempty"`
	Classification  string    `json:"classification,omitempty"`
	Designation     string    `json:"designation,omitempty"`
	AverageHeight   string    `json:"average_height,omitempty"`
	SkinColors      string    `json:"skin_colors,omitempty"`
	HairColors      string    `json:"hair_colors,omitempty"`
	EyeColors       string    `json:"eye_colors,omitempty"`
	AverageLifespan string    `json:"average_lifespan,omitempty"`
	Homeworld       string    `json:"homeworld,omitempty"`
	Language        string    `json:"language,omitempty"`
	People          []string  `json:"people,omitempty"`
	Films           []string  `json:"films,omitempty"`
	Created         time.Time `json:"created,omitempty"`
	Edited          time.Time `json:"edited,omitempty"`
	URL             string    `json:"url,omitempty"`
}
