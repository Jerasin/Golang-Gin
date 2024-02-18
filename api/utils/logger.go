package utils

import (
	"fmt"

	"github.com/Jerasin/configs"
)

type ConfigLogger struct {
	Environment string
}

func checkEnv(env string) bool {
	fmt.Println("INFO Environment: ", env)

	if env == "dev" {
		return true
	} else {
		return false
	}

}

func LoggerInfo(info string) {
	ENVIRONMENT := configs.GetEnv("ENVIRONMENT")

	if checkEnv(ENVIRONMENT) {
		fmt.Println("------------------------------------------------------")
		fmt.Println("------------------------------------------------------")
		fmt.Println("[INFO]: ", info)
		fmt.Println("------------------------------------------------------")
		fmt.Println("------------------------------------------------------")
	}
}
