package swapi

import (
	"context"

	smodel "github.com/Oyal2/SwapiLookup/pkg/model/client/swapi"
)

// Interface for the SWAPI Api
type SwapiProvider interface {
	SearchPeople(ctx context.Context, query string) ([]*smodel.FullPerson, error)
}
