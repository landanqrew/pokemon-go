package config

import (
	"fmt"
	"testing"
)


func TestGetRootDir(t *testing.T) {
	config := Config{}
	rootDir := config.GetRootDir()
	fmt.Printf("RootDir: %s", rootDir)
}