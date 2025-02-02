package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadConfig() {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "development"
	}

	var envFile string
	if env == "production" {
		envFile = ".env.prod"
	} else {
		envFile = ".env"
	}

	fmt.Printf("Loading config file %s\n", envFile)

	err := godotenv.Load(envFile)
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
