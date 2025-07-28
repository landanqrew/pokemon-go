package web

import (
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