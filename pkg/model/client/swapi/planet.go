package swapi

import "time"

type SearchPlanetsResponse struct {
	Count    int       `json:"count,omitempty"`
	Next     string    `json:"next,omitempty"`
	Previous string    `json:"previous,omitempty"`
	Results  []*Planet `json:"results,omitempty"`
}

type Planet struct {
	Name           string    `json:"name,omitempty"`
	RotationPeriod string    `json:"rotation_period,omitempty"`
	OrbitalPeriod  string    `json:"orbital_period,omitempty"`
	Diameter       string    `json:"diameter,omitempty"`
	Climate        string    `json:"climate,omitempty"`
	Gravity        string    `json:"gravity,omitempty"`
	Terrain        string    `json:"terrain,omitempty"`
	SurfaceWater   string    `json:"surface_water,omitempty"`
	Population     string    `json:"population,omitempty"`
	Residents      []string  `json:"residents,omitempty"`
	Films          []string  `json:"films,omitempty"`
	Created        time.Time `json:"created,omitempty"`
	Edited         time.Time `json:"edited,omitempty"`
	URL            string    `json:"url,omitempty"`
}
