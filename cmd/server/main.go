package main

import (
	"fmt"

	"github.com/balasl342/go-gin-starter-template/config"
	"github.com/balasl342/go-gin-starter-template/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// To set the gin mode to release.
	gin.SetMode(gin.ReleaseMode)

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	app := gin.Default()

	// Setup routes
	routes.SetupRoutes(app)

	// Get the port from configuration
	port := viper.GetString("port")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	//serve
	app.Run(fmt.Sprintf(":%s", port))
}
