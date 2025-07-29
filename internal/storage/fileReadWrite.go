package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/landanqrew/pokemon-go/internal/config"
)


func WriteBytes(byteData []byte, fileName string) error {
	config := config.Config{}
	rootProgramDir := config.GetRootDir()
	storageDir := rootProgramDir + "/storage"

	_, err := os.ReadDir(storageDir) 
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Storage Directory does not exist. Creating directory...")
			err = os.Mkdir(storageDir, 0755)
			if err != nil {
				return fmt.Errorf("cannot make storage directory. %v", err)
			}
		} else {
			return fmt.Errorf("cannot read storage directory. %v", err)
		}
	}

	err = os.WriteFile(storageDir + "/" + fileName, byteData, 0666)
	if err != nil {
		return fmt.Errorf("cannot write file. %v", err)
	}

	return nil
}


func ReadBytes(fileName string) ([]byte, error) {
	config := config.Config{}
	rootProgramDir := config.GetRootDir()
	storageDir := rootProgramDir + "/storage"

	_, err := os.ReadDir(storageDir) 
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Storage Directory does not exist. Creating directory...")
			err = os.Mkdir(storageDir, 0755)
			if err != nil {
				return nil, fmt.Errorf("cannot make storage directory. %v", err)
			}
		} else {
			return nil, fmt.Errorf("cannot read storage directory. %v", err)
		}
	}

	byteData, err := os.ReadFile(storageDir + "/" + fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot read file. %v", err)
	}

	return byteData, nil
}


func DeserializeJsonArray[T any](byteData []byte) ([]T, error) {
	var jsonData []T
	err := json.Unmarshal(byteData, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal jsonData. %v", err)
	}

	return jsonData, nil
}

func DeserializeJsonObject[T any](byteData []byte) (T, error) {
	var jsonData T
	err := json.Unmarshal(byteData, &jsonData)
	if err != nil {
		return jsonData, fmt.Errorf("cannot unmarshal jsonData. %v", err)
	}

	return jsonData, nil
}

func StorageFileExists(fileName string) bool {
	config := config.Config{}
	rootProgramDir := config.GetRootDir()
	storageDir := rootProgramDir + "/storage"

	_, err := os.ReadDir(storageDir)
	if err != nil {
		return false
	}

	_, err = os.Stat(storageDir + "/" + fileName)
	return err == nil
}