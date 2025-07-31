package pokecache

import (
	"reflect"
	"slices"
	"testing"
	"time"
)

func TestAddAndReap(t *testing.T) {
	cache := NewCache(time.Second * 5)

	cache.Add("test", []byte("test"))

	cache.Add("test", []byte("test"))

	cache.Add("test2", []byte("test2"))

	cache.Add("test3", []byte("test3"))

	cache.Add("test4", []byte("test4"))

	cache.Add("test5", []byte("test5"))

	expectedKeys := []string{"test", "test2", "test3", "test4", "test5"}

	actualKeys := []string{}

	for key := range cache.Entries {
		actualKeys = append(actualKeys, key)
	}

	slices.Sort(expectedKeys)
	slices.Sort(actualKeys)
	if !slices.Equal(expectedKeys, actualKeys) {
		t.Errorf("expected keys %v, got %v", expectedKeys, actualKeys)
	}
	expectedValues := []string{"test", "test2", "test3", "test4", "test5"}

	actualValues := []string{}

	for _, value := range cache.Entries {
		actualValues = append(actualValues, string(value.Value))
	}

	slices.Sort(expectedValues)
	slices.Sort(actualValues)
	if !reflect.DeepEqual(expectedValues, actualValues) {
		t.Errorf("expected values %v, got %v", expectedValues, actualValues)
	}

	// allow time to reap
	time.Sleep(time.Second * 6)

	actualKeys = []string{}

	for key := range cache.Entries {
		actualKeys = append(actualKeys, key)
	}

	if !slices.Equal([]string{}, actualKeys) {
		t.Errorf("expected no keys, got %v", actualKeys)
	}
}


func TestGet(t *testing.T) {
	cache := NewCache(time.Second * 10)

	cache.Add("test", []byte("test"))

	cache.Add("test2", []byte("test2"))

	cache.Add("test3", []byte("test3"))

	actual, err := cache.Get("test")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if string(actual) != "test" {
		t.Errorf("expected %v, got %v", "test", string(actual))
	}

	// should error because of invalid url
	_, err = cache.Get("test4")
	if err == nil {
		t.Error("expected error, got none")
	}
	
}