package db

import (
	"encoding/json"
	"os"
)

func Configuration(pathConfig string, config *Config) {
	file, err := os.Open(pathConfig)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
}
