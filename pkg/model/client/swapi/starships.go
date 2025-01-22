package swapi

import "time"

type SearchStarshipsResponse struct {
	Count    int         `json:"count,omitempty"`
	Next     string      `json:"next,omitempty"`
	Previous string      `json:"previous,omitempty"`
	Results  []*Starship `json:"results,omitempty"`
}
type Starship struct {
	Name                 string    `json:"name,omitempty"`
	Model                string    `json:"model,omitempty"`
	Manufacturer         string    `json:"manufacturer,omitempty"`
	CostInCredits        string    `json:"cost_in_credits,omitempty"`
	Length               string    `json:"length,omitempty"`
	MaxAtmospheringSpeed string    `json:"max_atmosphering_speed,omitempty"`
	Crew                 string    `json:"crew,omitempty"`
	Passengers           string    `json:"passengers,omitempty"`
	CargoCapacity        string    `json:"cargo_capacity,omitempty"`
	Consumables          string    `json:"consumables,omitempty"`
	HyperdriveRating     string    `json:"hyperdrive_rating,omitempty"`
	Mglt                 string    `json:"MGLT,omitempty"`
	StarshipClass        string    `json:"starship_class,omitempty"`
	Pilots               []string  `json:"pilots,omitempty"`
	Films                []string  `json:"films,omitempty"`
	Created              time.Time `json:"created,omitempty"`
	Edited               time.Time `json:"edited,omitempty"`
	URL                  string    `json:"url,omitempty"`
}
