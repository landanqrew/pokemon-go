package pokemon

import (
	"encoding/json"
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/storage"
	"github.com/landanqrew/pokemon-go/internal/web"
)


type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (l LocationAreaResponse) GetCount() int {
	return l.Count
}

func (l LocationAreaResponse) GetResults() []PokeType {
	results := []PokeType{}
	for _, result := range l.Results {
		results = append(results, result)
	}
	return results
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (l LocationArea) GetName() string {
	return l.Name
}

func (l LocationArea) GetURL() string {
	return l.URL
}

func (l LocationArea) GetPokemonNames() ([]string, error) {
	pokemonNames := []string{}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", l.Name)
	res, err := web.GetResponseBytesBaseUrl(url)
	if err  != nil {
		return pokemonNames, fmt.Errorf("cannot get reponse details. %v", err) 
	}

	locationDetails := ExploredLocationArea{}
	err = json.Unmarshal(res, &locationDetails)
	if err != nil {
		return pokemonNames, fmt.Errorf("cannot unmarshal to ExpoloredLocationArea struct. %v", err)
	}

	for _, encounter := range locationDetails.PokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	return pokemonNames, nil
}


func GetLocationAreas() ([]LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	responses, err := GetAllResponses[LocationAreaResponse](url, 500)
	if err != nil {
		return nil, err
	}

	locationAreas := []LocationArea{}
	for _, response := range responses {
		results := response.GetResults()
		for _, result := range results {
			locationAreas = append(locationAreas, result.(LocationArea))
		}
	}
	return locationAreas, nil
}

func ReadLocationsFromCache() ([]LocationArea, error) {
	if !storage.StorageFileExists("locationAreas.json") {
		locationAreas, err := GetAndStoreLocationAreas()
		if err != nil {
			return nil, err
		}
		return locationAreas, nil
	}

	jsonBytes, err := storage.ReadBytes("locationAreas.json")
	if err != nil {
		return nil, err
	}

	locationAreas := []LocationArea{}
	err = json.Unmarshal(jsonBytes, &locationAreas)
	if err != nil {
		return nil, err
	}
	return locationAreas, nil
}

func GetLocationNames() ([]string, error) {
	locationAreas, err := ReadLocationsFromCache()
	if err != nil {
		return nil, err
	}

	locationNames := []string{}
	for _, locationArea := range locationAreas {
		locationNames = append(locationNames, locationArea.Name)
	}
	return locationNames, nil
}

func GetAndStoreLocationAreas() ([]LocationArea, error) {
	locationAreas, err := GetLocationAreas()
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.MarshalIndent(locationAreas, "", "  ")
	if err != nil {
		return nil, err
	}

	err = storage.WriteBytes(jsonBytes, "locationAreas.json")
	if err != nil {
		return nil, err
	}

	return locationAreas, nil
}


type ExploredLocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
