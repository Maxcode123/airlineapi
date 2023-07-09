package main

import "os"

// Configuration struct holding all configuration settings.
type Config struct {
	DuffelAPIToken string
	DuffelAPI      string
	APIPort        string
}

// Initializes the configuration by reading environment variables.
func InitConfig() {
	Conf.DuffelAPI = os.Getenv("DUFFEL_API")
	Conf.DuffelAPIToken = os.Getenv("DUFFEL_API_TOKEN")
	Conf.APIPort = os.Getenv("API_PORT")
}

// Configuration instance to be used throughout the project.
var Conf Config
