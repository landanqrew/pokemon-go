package pokemon

import (
	"testing"
	"time"

	"github.com/landanqrew/pokemon-go/internal/api"
	"github.com/landanqrew/pokemon-go/internal/pokecache"
)

func TestCatch(t *testing.T) {
	iterations := 10000
	cachePtr := pokecache.NewCache(30 * time.Second)
	client := api.NewClient(cachePtr, "https://pokeapi.co/api/v2/")
	cases := []struct {
		pokemonName string
		minCountExpected int
		maxCountExpected int
	}{
		{"mew", 1000, 3000},
		{"mewtwo", 1000, 3000},
		{"ho-oh", 1000, 3000},
		{"lugia", 1000, 3000},
		{"celebi", 1000, 6000},
		{"articuno", 1000, 6000},
	}
	for _, c := range cases {
		response, err := client.Cache.Get(client.BaseURL + "pokemon/" + c.pokemonName)
		if err != nil {
			t.Errorf("error getting pokemon: %v", err)
		}
		pokemon, err := ParsePokemonFromResponse(response)
		if err != nil {
			t.Errorf("error parsing pokemon: %v", err)
		}
		count := 0
		for i := 0; i < iterations; i++ {
			if pokemon.Catch() {
				count++
			}
		}
		if count < c.minCountExpected || count > c.maxCountExpected {
			t.Errorf("pokemon %s caught %d times, expected between %d and %d", c.pokemonName, count, c.minCountExpected, c.maxCountExpected)
		}
	}
}