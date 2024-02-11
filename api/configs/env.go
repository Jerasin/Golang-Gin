package configs

import (
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

func GetEnv(name string) string {
	return os.Getenv(name)
}

func EnvConfig() {
	loadConfig()
}
