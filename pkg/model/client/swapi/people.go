package swapi

import "time"

type FullPerson struct {
	Person        *Person
	HomeworldData *Planet     `json:"-"`
	SpeciesData   []*Species  `json:"-"`
	StarshipsData []*Starship `json:"-"`
}

type SearchPeopleResponse struct {
	Count    int       `json:"count,omitempty"`
	Next     string    `json:"next,omitempty"`
	Previous string    `json:"previous,omitempty"`
	Results  []*Person `json:"results,omitempty"`
}
type Person struct {
	Name      string    `json:"name,omitempty"`
	Height    string    `json:"height,omitempty"`
	Mass      string    `json:"mass,omitempty"`
	HairColor string    `json:"hair_color,omitempty"`
	SkinColor string    `json:"skin_color,omitempty"`
	EyeColor  string    `json:"eye_color,omitempty"`
	BirthYear string    `json:"birth_year,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Homeworld string    `json:"homeworld,omitempty"`
	Films     []string  `json:"films,omitempty"`
	Species   []string  `json:"species,omitempty"`
	Vehicles  []string  `json:"vehicles,omitempty"`
	Starships []string  `json:"starships,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	Edited    time.Time `json:"edited,omitempty"`
	URL       string    `json:"url,omitempty"`
}
