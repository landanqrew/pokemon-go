package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func FetchAndSerializeArray[T any](url string) ([]T, error) {
	output := []T{}
	res, err := http.Get(url)
	if err != nil {
		return output, fmt.Errorf("failed to fetch data from url (%s): %v", url, err.Error())
	}
	// i always forget this part
	defer res.Body.Close()

	byteSlice, err := io.ReadAll(res.Body)
	if err != nil {
		return output, fmt.Errorf("could not read bytes from response body, %w", err)
	}

	err = json.Unmarshal(byteSlice, &output)
	if err != nil {
		return output, fmt.Errorf("cannot unmarshal response to designated type: %w", err)
	}

	return output, nil
}

func FetchAndSerializeStruct[T any](url string) (T, int, error) {
	outputPtr := new(T)
	output := *outputPtr
	res, err := http.Get(url)
	if err != nil {
		return output, res.StatusCode, fmt.Errorf("failed to fetch data from url (%s): %v", url, err.Error())
	}
	// i always forget this part
	defer res.Body.Close()

	byteSlice, err := io.ReadAll(res.Body)
	if err != nil {
		return output, res.StatusCode, fmt.Errorf("could not read bytes from response body, %w", err)
	}

	err = json.Unmarshal(byteSlice, &output)
	if err != nil {
		return output, res.StatusCode, fmt.Errorf("cannot unmarshal response to designated type: %w", err)
	}

	return output, res.StatusCode, nil
}


func GetPrint(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return  fmt.Errorf("failed to fetch data from url (%s): %v", url, err.Error())
	}

	byteSlice, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("could not read bytes from response body, %w", err)
	}

	fmt.Println("response:\n",string(byteSlice))
	return nil
}

