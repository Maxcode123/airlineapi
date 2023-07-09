package main

import (
	"errors"
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
func InitConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		return errors.New("could not load .env file")
	}
	Conf.DuffelAPI = os.Getenv("DUFFEL_API")
	Conf.DuffelAPIToken = os.Getenv("DUFFEL_API_TOKEN")
	Conf.APIPort = os.Getenv("API_PORT")
	return nil
}

// Configuration instance to be used throughout the project.
var Conf Config
