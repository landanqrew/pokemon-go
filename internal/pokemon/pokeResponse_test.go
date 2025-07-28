package pokemon

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestGetSubsetSerialized(t *testing.T) {
	cases := []struct {
		name        string
		url         string
		expectError bool
	}{
		{
			name:        "valid url",
			url:         "https://pokeapi.co/api/v2/location-area/",
			expectError: false,
		},
		{
			name:        "invalid url",
			url:         "https://invalid-url.com/api",
			expectError: true,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			dataChan := make(chan PokeTypeResponse)
			errChan := make(chan error)

			go GetSubsetSerialized(dataChan, errChan, test.url, 100, 0, 100)

			// Use a select with a timeout to wait for a result
			select {
			case result := <-dataChan:
				jsonString, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					t.Errorf("error marshalling result to json: %v", err)
				}
				fmt.Println(string(jsonString))
				if test.expectError {
					t.Errorf("expected an error, but got a result: %v", result)
				}
				if len(result.Results) == 0 {
					t.Errorf("expected results, but got an empty slice")
				}
			case err := <-errChan:
				if !test.expectError {
					t.Errorf("expected no error, but got: %v", err)
				}
			case <-time.After(5 * time.Second): // Timeout after 5 seconds
				t.Fatal("test timed out waiting for a response")
			}

			close(dataChan)
			close(errChan)

			
		})
	}
}
