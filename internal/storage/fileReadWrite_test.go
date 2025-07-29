package storage

import (
	"bytes"
	"os"
	"testing"

	"github.com/landanqrew/pokemon-go/internal/config"
)


func TestWriteBytes(t *testing.T) {
	jsonBytes := []byte(`{"name": "test"}`)
	err := WriteBytes(jsonBytes, "test.json")
	if err != nil {
		t.Errorf("Error writing bytes: %v", err)
	}
}

func TestReadBytes(t *testing.T) {
	expected := []byte(`{"name": "test"}`)
	actual, err := ReadBytes("test.json")
	if err != nil {
		t.Errorf("Error reading bytes: %v", err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	config := config.Config{}
	rootDir := config.GetRootDir()
	storageDir := rootDir + "/storage"

	err = os.Remove(storageDir + "/test.json")
	if err != nil {
		t.Errorf("Error removing file: %v", err)
	}
}