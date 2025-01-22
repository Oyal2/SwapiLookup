package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Oyal2/SwapiLookup/pkg/client/swapi"
	smodel "github.com/Oyal2/SwapiLookup/pkg/model/client/swapi"

	"golang.org/x/net/context"
)

type App struct {
	reader      *bufio.Reader
	swapiClient swapi.SwapiProvider
}

func New() (*App, error) {
	cfg := swapi.DefaultSwapiConfig()
	sclient, err := swapi.New(cfg)
	if err != nil {
		return nil, err
	}
	return &App{
		reader:      bufio.NewReader(os.Stdin),
		swapiClient: sclient,
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	fmt.Println("Enter a starwars character name to search, or 'C+CRTL' to quit")
	fmt.Println(strings.Repeat("-", 40))
	for ctx.Err() == nil {
		fmt.Print("\nEnter search term: ")
		query, err := a.reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		query = strings.TrimSpace(query)
		people, err := a.swapiClient.SearchPeople(ctx, query)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(people) == 0 {
			fmt.Printf("could not find the term: %v\n", query)
		}

		for _, person := range people {
			fmt.Println(formatOutput(person))
		}
	}
	return nil
}

func formatOutput(fp *smodel.FullPerson) string {
	var sb strings.Builder

	sb.WriteString(strings.Repeat("=", 50) + "\n")
	sb.WriteString(fmt.Sprintf("\nCharacter: %s\n\n", fp.Person.Name))

	if len(fp.StarshipsData) > 0 {
		for i, ship := range fp.StarshipsData {
			sb.WriteString(fmt.Sprintf("Starship %d:\n", i+1))
			sb.WriteString(fmt.Sprintf("  Name: %s\n", ship.Name))
			sb.WriteString(fmt.Sprintf("  Cargo Capacity: %s\n", ship.CargoCapacity))
			sb.WriteString(fmt.Sprintf("  Class: %s\n", ship.StarshipClass))
			sb.WriteString("\n")
		}
	} else {
		sb.WriteString("No starship information available\n\n")
	}

	sb.WriteString("Home Planet:\n")
	if fp.HomeworldData != nil {
		sb.WriteString(fmt.Sprintf("  Name: %s\n", fp.HomeworldData.Name))
		sb.WriteString(fmt.Sprintf("  Population: %s\n", fp.HomeworldData.Population))
		sb.WriteString(fmt.Sprintf("  Climate: %s\n", fp.HomeworldData.Climate))
	} else {
		sb.WriteString("  No planet information available\n")
	}
	sb.WriteString("\n")

	sb.WriteString("Species:\n")
	if len(fp.SpeciesData) > 0 {
		for _, species := range fp.SpeciesData {
			sb.WriteString(fmt.Sprintf("  Name: %s\n", species.Name))
			sb.WriteString(fmt.Sprintf("  Language: %s\n", species.Language))
			sb.WriteString(fmt.Sprintf("  Average Lifespan: %s\n\n", species.AverageLifespan))
		}
	} else {
		sb.WriteString("  No species information available\n\n")
	}

	return sb.String()

}
