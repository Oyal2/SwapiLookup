package swapi

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"sync"

	smodel "github.com/Oyal2/SwapiLookup/pkg/model/client/swapi"
)

// Search for the Starwars characters
func (sc *SwapiClient) SearchPeople(ctx context.Context, query string) ([]*smodel.FullPerson, error) {
	var people []*smodel.Person
	people, err := sc.searchPeopleWithPagination(ctx, fmt.Sprintf("%s/people/?search=%s", sc.baseURL, query))
	if err != nil {
		return nil, err
	}

	var fps []*smodel.FullPerson
	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(people))

	for _, person := range people {
		wg.Add(1)
		go func(person *smodel.Person) {
			defer wg.Done()
			fullPerson, err := sc.enrichPeople(ctx, person)
			if err != nil {
				errChan <- fmt.Errorf("failed to enrich character %s: %v", person.Name, err)
				return
			}
			mu.Lock()
			fps = append(fps, fullPerson)
			mu.Unlock()
		}(person)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	slices.SortFunc(fps, func(a, b *smodel.FullPerson) int {
		return cmp.Compare(a.Person.Name, b.Person.Name)
	})

	return fps, nil
}

// Handles the pagination when doing a search query for a character
func (sc *SwapiClient) searchPeopleWithPagination(ctx context.Context, uri string) ([]*smodel.Person, error) {
	var people []*smodel.Person
	for uri != "" {
		var searchResp smodel.SearchPeopleResponse

		err := sc.get(ctx, uri, &searchResp)
		if err != nil {
			return nil, fmt.Errorf("failed to search people: %v", err)
		}

		people = append(people, searchResp.Results...)

		uri = searchResp.Next
	}

	return people, nil
}

// Gets all the information about a character
func (sc *SwapiClient) enrichPeople(ctx context.Context, person *smodel.Person) (*smodel.FullPerson, error) {
	fp := &smodel.FullPerson{
		Person:        person,
		HomeworldData: &smodel.Planet{},
		SpeciesData:   []*smodel.Species{},
		StarshipsData: []*smodel.Starship{},
	}
	fp.Person = person
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Utilize goroutines and try to run as many as the "worker pool" allows in order to speed up the enrichment process
	// We make sure to use mutex since we dont want a race condition when modifying the fp variable

	chanLen := 1 + len(person.Species) + len(person.Starships)
	errChan := make(chan error, chanLen)

	if person.Homeworld != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := sc.sem.Acquire(ctx, 1); err != nil {
				errChan <- fmt.Errorf("failed to acquire semaphore for homeworld: %w", err)
				return
			}
			defer sc.sem.Release(1)

			var planet smodel.Planet
			err := sc.get(ctx, person.Homeworld, &planet)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch planet: %w", err)
				return
			}

			mu.Lock()
			fp.HomeworldData = &planet
			mu.Unlock()
		}()
	}

	for _, url := range person.Species {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			if err := sc.sem.Acquire(ctx, 1); err != nil {
				errChan <- fmt.Errorf("failed to acquire semaphore for species: %w", err)
				return
			}
			defer sc.sem.Release(1)

			var species smodel.Species
			err := sc.get(ctx, url, &species)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch species: %w", err)
				return
			}

			mu.Lock()
			fp.SpeciesData = append(fp.SpeciesData, &species)
			mu.Unlock()
		}(url)
	}

	for _, url := range person.Starships {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			if err := sc.sem.Acquire(ctx, 1); err != nil {
				errChan <- fmt.Errorf("failed to acquire semaphore for starship: %w", err)
				return
			}

			defer sc.sem.Release(1)

			var starship smodel.Starship
			err := sc.get(ctx, url, &starship)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch starship: %w", err)
				return
			}

			mu.Lock()
			fp.StarshipsData = append(fp.StarshipsData, &starship)
			mu.Unlock()
		}(url)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return fp, nil
}

func (sc *SwapiClient) get(ctx context.Context, uri string, target interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return fmt.Errorf("failed to create a request: %v", err)
	}

	resp, err := sc.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed sending a request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get a successful response, %d: %s", resp.StatusCode, resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
