package pokemon

import (
	"encoding/json"
	"fmt"

	"github.com/landanqrew/pokemon-go/internal/web"
)

type PokeApiResponse interface {
	GetCount() int
	GetResults() []PokeType
}

type PokeType interface {
	GetName() string
	GetURL() string
}

// ADD this struct to represent a standard API resource
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// ADD methods to make NamedAPIResource implement PokeType
func (n NamedAPIResource) GetName() string { return n.Name }
func (n NamedAPIResource) GetURL() string  { return n.URL }

// not sure if this can be used for all responses yet
type PokeTypeResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"` // Use the concrete type here
}

// ADD methods to make PokeTypeResponse implement PokeApiResponse
func (p PokeTypeResponse) GetCount() int { return p.Count }
func (p PokeTypeResponse) GetResults() []PokeType {
	res := make([]PokeType, len(p.Results))
	for i, r := range p.Results {
		res[i] = r
	}
	return res
}

func GetSubsetSerialized[T PokeApiResponse](dataChan chan T, errChan chan error, baseUrl string, limit int, offset int, max int) {
	// REMOVE these defer statements. The caller should manage the channel lifecycle.
	// defer close(dataChan)
	// defer close(errChan)
	if offset+limit > max {
		limit = max - offset
	}
	url := fmt.Sprintf("%s?limit=%d&offset=%d", baseUrl, limit, offset)
	response, statusCode, err := web.FetchAndSerializeStruct[T](url)
	if err != nil {
		errChan <- fmt.Errorf("error fetching location area subset with error: %v and status code: %d", err, statusCode)
	} else {
		dataChan <- response
	}
}

func GetAllResponses[T PokeApiResponse](baseUrl string, limit int) ([]T, error) {
	responses := []T{}
	res, statusCode, err := web.FetchAndSerializeStruct[T](baseUrl)
	if err != nil {
		return nil, fmt.Errorf("error fetching object response with error: %v and status code: %d", err, statusCode)
	}
	responses = append(responses, res)
	dataChan := make(chan T)
	defer close(dataChan)
	errChan := make(chan error)
	defer close(errChan)
	go GetSubsetSerialized[T](dataChan, errChan, baseUrl, limit, 0, res.GetCount())
	// handle channel events
	for {
		select {
		case response, ok := <-dataChan:
			if !ok {
				dataChan = nil
			} else {
				responses = append(responses, response)
			}
		case err, ok := <-errChan:
			if !ok {
				errChan = nil
			} else {
				fmt.Printf("error fetching object subset with error: %v and status code: %d", err, statusCode)
			}
		}
		// break if both channels are exhausted
		if dataChan == nil && errChan == nil {
			break
		}
	}
	return responses, nil
}

func PrintAllResults[T PokeType](results []T) {
	bytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("error marshalling results to json", err)
	}
	fmt.Println(string(bytes))
}

func PrintResponses[T PokeApiResponse](responses []T) {
	bytes, err := json.MarshalIndent(responses, "", "  ")
	if err != nil {
		fmt.Println("error marshalling results to json", err)
	}
	fmt.Println(string(bytes))
}
