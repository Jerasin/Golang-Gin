package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func GetEnv(name, setDefault string) string {
	result := os.Getenv(name)

	if result == "" && setDefault != "" {
		fmt.Println("default", setDefault)
		result = setDefault
	}

	return result
}

func EnvConfig() {
	loadConfig()
}
