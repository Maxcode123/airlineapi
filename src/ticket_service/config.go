package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Configuration struct holding all configuration settings.
type Config struct {
	DuffelAPIToken string
	DuffelAPI      string
	APIPort        string
}

// Initializes the configuration by reading environment variables.
func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not load .env file")
	}
	Conf.DuffelAPI = os.Getenv("DUFFEL_API")
	Conf.DuffelAPIToken = os.Getenv("DUFFEL_API_TOKEN")
	Conf.APIPort = os.Getenv("API_PORT")
}

// Configuration instance to be used throughout the project.
var Conf Config
