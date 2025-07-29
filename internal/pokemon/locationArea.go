package pokemon

import (
	"encoding/json"

	"github.com/landanqrew/pokemon-go/internal/storage"
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


/*
func GetLocationAreas() ([]LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	locAreaResponse, statusCode, err := web.FetchAndSerializeStruct[LocationAreaResponse](url)
	if err != nil {
		return nil, fmt.Errorf("error fetching location area response with error: %v and status code: %d", err, statusCode)
	}
	locChan := make(chan LocationArea)
	errChan := make(chan error)
	go GetLocationAreaSubsets(locChan, errChan, url, 100, 0, locAreaResponse.Count)

	locationAreas := []LocationArea{}
	for {
		select {
		case locationArea, ok := <-locChan:
			if !ok {
				// locChan is closed and exhausted
				locChan = nil // Set to nil to make this case non-selectable
			} else {
				locationAreas = append(locationAreas, locationArea)
			}
		case err, ok := <-errChan:
			if !ok {
				// errChan is closed and exhausted
				errChan = nil // Set to nil to make this case non-selectable
			} else {
				fmt.Println(err.Error()) // Print error for visibility
			}
		}

		// Break condition: exit loop when both channels are nil (closed and exhausted)
		if locChan == nil && errChan == nil {
			break
		}
	}
	

	return locationAreas, nil
}
*/

/*
func GetLocationAreaSubsets(locChan chan LocationArea, errChan chan error, baseUrl string, limit int, offset int, max int) {
	defer close(locChan)
	for i := offset; i < max; i+= limit {
		if i + limit > max {
			limit = max - i
		}
		url := fmt.Sprintf("%s?limit=%d&offset=%d", baseUrl, limit, i)
		locationAreaResponse, statusCode, err := web.FetchAndSerializeStruct[LocationAreaResponse](url)
		if err != nil {
			errChan <- fmt.Errorf("error fetching location area subset with error: %v and status code: %d", err, statusCode)
			//channel <- LocationArea{} // maybe retry call if the error is of a certain type
		}
		for _, locationArea := range locationAreaResponse.Results {
			locChan <- locationArea
		}
	}
	
}*/