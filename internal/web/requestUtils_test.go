package web

import (
	"encoding/json"
	"fmt"
	"testing"

)

// make sure you use -test.v to see the stdout
func TestGetPrint(t *testing.T) {
	cases := []struct{
		url string
		expectedError bool
	}{
		{
			url: "https://pokeapi.co/api/v2/location-area/",
			expectedError: false,
		},
		{
			url: "https://pokeapi.co/api/v2/location-area/1/",
			expectedError: false,
		},
	}

	for _, test := range cases {
		t.Logf("testing url, %s", test.url)
		err := GetPrint(test.url)
		if test.expectedError && err == nil {
			t.Errorf("Expected error for get request at url (%s)", test.url)
		} else if (!test.expectedError && err != nil) {
			t.Errorf("Expected no error for get request at url (%s), but got error %s", test.url, err.Error())
		}
	}
}

func TestGetResponseBytes(t *testing.T) {
	cases := []struct{
		url string
		limit int
		offset int
		expectedError bool
	}{
		{
			url: "https://pokeapi.co/api/v2/location-area/",
			limit: 100,
			offset: 0,
			expectedError: false,
		},
		{
			url: "https://pokeapi.co/api/v2/location-area/",
			limit: 2000, // number greater than count
			offset: 0,
			expectedError: false,
		},
	}

	for _, test := range cases {
		t.Logf("testing url, %s", test.url)
		byteSlice, err := GetResponseBytes(test.url, test.limit, test.offset)
		if test.expectedError && err == nil {
			t.Errorf("Expected error for get request at url (%s)", test.url)
		} else if (!test.expectedError && err != nil) {
			t.Errorf("Expected no error for get request at url (%s), but got error %s", test.url, err.Error())
		}
		response := struct {
				Count    int                `json:"count"`
				Next     string             `json:"next"`
				Previous string             `json:"previous"`
				Results  []struct{
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"results"`
		}{}
		err = json.Unmarshal(byteSlice, &response)
		if err != nil {
			t.Errorf("Error unmarshalling location area response: %v", err)
		}
		fmt.Printf("count: %d, resultsCount: %d\n", response.Count, len(response.Results))
	}
}