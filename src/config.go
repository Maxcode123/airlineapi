package main

import "os"

type Config struct {
	DuffelAPIToken string
	DuffelAPI      string
}

func InitConfig() {
	Conf.DuffelAPI = os.Getenv("DUFFEL_API")
	Conf.DuffelAPIToken = os.Getenv("DUFFEL_API_TOKEN")
}

var Conf Config
