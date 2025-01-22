# Swapi Lookup

A cli to search for Star Wars characters and view their information including homeworld, species, and starships.

## Example

Search for `Luke`

```bash
make run

Star Wars Character Information Searcher
Enter a starwars character name to search, or 'C+CRTL' to quit
----------------------------------------

==================================================

Character: Luke Skywalker

Starship 1:
  Name: Imperial shuttle
  Cargo Capacity: 80000
  Class: Armed government transport

Starship 2:
  Name: X-wing
  Cargo Capacity: 110
  Class: Starfighter

Home Planet:
  Name: Tatooine
Home Planet:
Home Planet:
  Name: Tatooine
  Population: 200000
  Climate: arid

Species:
  No species information available
```

Search for partial names `sky`

```bash
make run

Enter a starwars character name to search, or 'C+CRTL' to quit
----------------------------------------

Enter search term: sky
==================================================

Character: Anakin Skywalker

Starship 1:
  Name: Naboo fighter
  Cargo Capacity: 65
  Class: Starfighter

Starship 2:
  Name: Trade Federation cruiser
  Cargo Capacity: 50000000
  Class: capital ship

Starship 3:
  Name: Jedi Interceptor
  Cargo Capacity: 60
  Class: starfighter

Home Planet:
  Name: Tatooine
  Population: 200000
  Climate: arid

Species:
  No species information available


==================================================

Character: Luke Skywalker

Starship 1:
  Name: Imperial shuttle
  Cargo Capacity: 80000
  Class: Armed government transport

Starship 2:
  Name: X-wing
  Cargo Capacity: 110
  Class: Starfighter

Home Planet:
  Name: Tatooine
  Population: 200000
  Climate: arid

Species:
  No species information available


==================================================

Character: Shmi Skywalker

No starship information available

Home Planet:
  Name: Tatooine
  Population: 200000
  Climate: arid

Species:
  No species information available
```

## Installation

### Prerequisites
- Go 1.22.0 or higher
- Make (optional, for using Makefile commands)

#### Environment Configuration

The application uses the following environment variables that can be set in a `.env` file in the `cmd/finder` folder:

```bash
SWAPI_BASE_URL="https://swapi.dev/api"     # Base URL for SWAPI
SWAPI_MAX_CONCURRENT=10                     # Maximum number of concurrent API requests
```

### Building
1. Clone the repository
2. Run make commands at root directory:
```bash
# Run tests
make test

# Build and run
make run
```

