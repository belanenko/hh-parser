package clickconfig

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbUsername string `json:"DB_USERNAME"`
	DbPassword string `json:"DB_PASSWORD"`
	DbHost     string `json:"DB_HOST"`
	DbName     string `json:"DB_NAME"`
	DbCertPath string `json:"DB_CERT_PATH"`
}

func ReadConfig(path string) *Config {
	text, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	var outConfig Config
	json.Unmarshal(text, &outConfig)
	if outConfig.DbName == "" {
		log.Fatalln("empty config")
	}
	return &outConfig
}
