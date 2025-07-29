package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DefaultStorageLocation string
}

func (c *Config) GetRootDir() string {
	if c.DefaultStorageLocation == "" {
		c.Init()
	}
	return c.DefaultStorageLocation
}

func (c *Config) Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot initialize config object. Cannot identify users home directory. Exiting program. See error below:\n", err)
	}
	c.DefaultStorageLocation = homeDir + "/.pokemon-go"

	_, err = os.ReadDir(c.DefaultStorageLocation) 
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Root Program Directory does not exist. Creating directory...")
			err = os.Mkdir(c.DefaultStorageLocation, 0755)
			if err != nil {
				log.Fatal("Cannot initialize config object. Cannot create directory. Exiting program. See error below:\n", err)
			}
		} else {
			log.Fatal("Cannot initialize config object. Cannot read directory. Exiting program. See error below:\n", err)
		}
	}

}