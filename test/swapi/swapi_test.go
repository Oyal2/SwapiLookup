package swapi

import (
	"cmp"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"

	"github.com/Oyal2/SwapiLookup/pkg/client/swapi"
	smodel "github.com/Oyal2/SwapiLookup/pkg/model/client/swapi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SwapiClient", func() {
	var (
		client    *swapi.SwapiClient
		server    *httptest.Server
		serverMux *http.ServeMux
	)

	BeforeEach(func() {
		serverMux = http.NewServeMux()
		server = httptest.NewServer(serverMux)
		var err error
		cfg := swapi.SwapiClientConfig{
			BaseURL:       server.URL,
			MaxConcurrent: 10,
		}
		client, err = swapi.New(cfg)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("SearchPeople", func() {
		Context("when searching for Luke Skywalker", func() {
			BeforeEach(func() {
				serverMux.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.SearchPeopleResponse{
						Count: 1,
						Results: []*smodel.Person{
							{
								Name:      "Luke Skywalker",
								Homeworld: server.URL + "/planets/1",
								Species:   []string{server.URL + "/species/1"},
								Starships: []string{server.URL + "/starships/12", server.URL + "/starships/13"},
							},
						},
					})
				})

				serverMux.HandleFunc("/planets/1", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.Planet{
						Name:    "Tatooine",
						Climate: "arid",
						Terrain: "desert",
					})
				})

				serverMux.HandleFunc("/species/1", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.Species{
						Name:           "Human",
						Classification: "mammal",
						Language:       "Galactic Basic",
					})
				})

				serverMux.HandleFunc("/starships/", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.Starship{
						Name:          "X-wing",
						Model:         "T-65 X-wing",
						StarshipClass: "Starfighter",
					})
				})
			})

			It("should return Luke's data with all related info", func() {
				ctx := context.Background()
				results, err := client.SearchPeople(ctx, "luke")

				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(1))
				Expect(results[0].Person.Name).To(Equal("Luke Skywalker"))
				Expect(results[0].HomeworldData.Name).To(Equal("Tatooine"))
				Expect(results[0].StarshipsData).To(HaveLen(2))
				Expect(results[0].SpeciesData).To(HaveLen(1))
			})
		})

		Context("when searching with pagination", func() {
			BeforeEach(func() {
				serverMux.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Query().Get("page") == "2" {
						json.NewEncoder(w).Encode(&smodel.SearchPeopleResponse{
							Count: 3,
							Results: []*smodel.Person{
								{Name: "Pilot 2", Homeworld: server.URL + "/planets/1"},
								{Name: "Pilot 3", Homeworld: server.URL + "/planets/1"},
							},
						})
						return
					}

					json.NewEncoder(w).Encode(&smodel.SearchPeopleResponse{
						Count: 3,
						Next:  server.URL + "/people?search=pilot&page=2",
						Results: []*smodel.Person{
							{Name: "Pilot 1", Homeworld: server.URL + "/planets/1"},
						},
					})
				})

				serverMux.HandleFunc("/planets/1", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.Planet{Name: "Tatooine"})
				})
			})

			It("should handle pagination correctly", func() {
				ctx := context.Background()
				results, err := client.SearchPeople(ctx, "pilot")

				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(3))
				Expect(slices.IsSortedFunc(results, func(a, b *smodel.FullPerson) int {
					return cmp.Compare(a.Person.Name, b.Person.Name)
				})).To(BeTrue())
			})
		})

		Context("when dealing with errors", func() {
			It("should handle HTTP errors", func() {
				serverMux.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				})

				ctx := context.Background()
				results, err := client.SearchPeople(ctx, "error")

				Expect(err).To(HaveOccurred())
				Expect(results).To(BeNil())
			})

			It("should handle malformed responses", func() {
				serverMux.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(`invalid json`))
				})

				ctx := context.Background()
				results, err := client.SearchPeople(ctx, "malformed")

				Expect(err).To(HaveOccurred())
				Expect(results).To(BeNil())
			})

			It("should handle missing data gracefully", func() {
				serverMux.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.SearchPeopleResponse{
						Results: []*smodel.Person{{
							Name:      "Missing Data Person",
							Homeworld: "invalid_url",
							Species:   []string{"invalid_species_url"},
						}},
					})
				})

				ctx := context.Background()
				results, err := client.SearchPeople(ctx, "missing")

				Expect(err).To(HaveOccurred())
				Expect(results).To(BeNil())
			})
		})

		Context("when enriching person data", func() {
			It("should handle missing related data gracefully", func() {
				serverMux.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.SearchPeopleResponse{
						Results: []*smodel.Person{{
							Name:      "Missing Data Person",
							Homeworld: "invalid_url",
							Species:   []string{"invalid_species_url"},
						}},
					})
				})

				ctx := context.Background()
				results, err := client.SearchPeople(ctx, "missing")

				Expect(err).To(HaveOccurred())
				Expect(results).To(BeNil())
			})

			It("should handle empty related data arrays", func() {
				serverMux.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.SearchPeopleResponse{
						Results: []*smodel.Person{{
							Name:      "Empty Data Person",
							Homeworld: server.URL + "/planets/1",
						}},
					})
				})

				serverMux.HandleFunc("/planets/1", func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(&smodel.Planet{Name: "Tatooine"})
				})

				ctx := context.Background()
				results, err := client.SearchPeople(ctx, "empty")

				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(1))
				Expect(results[0].SpeciesData).To(BeEmpty())
				Expect(results[0].StarshipsData).To(BeEmpty())
			})
		})
	})
})
