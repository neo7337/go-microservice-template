package main

import (
	"os"

	"github.com/neo7337/go-microservice-template/internal/app"
	"github.com/neo7337/go-microservice-template/internal/config"
	"oss.nandlabs.io/golly/l3"
)

var logger = l3.Get()

func main() {
	// Determine the application environment (default to "local" if not set)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}
	logger.InfoF("Starting application in %s environment", appEnv)
	// Build the configuration file name based on the environment
	configFileName := "config-" + appEnv + ".yml"
	logger.InfoF("Using configuration file: %s", configFileName)
	// Load the configuration file
	_, err := config.LoadConfig(configFileName)
	if err != nil {
		logger.ErrorF("Error loading config:", err)
		return
	}
	app.Start()
}
