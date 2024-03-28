package main

import (
	"binder/configs"
	"binder/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	e := echo.New()
	e.Static("/static", "views/public")

	// initiate configs
	configs.InitGoogleConfig()

	// initiate routes
	routes.InitWebRoutes(e)
	routes.InitAPIRoutes(e)

	e.Logger.Fatal(e.Start(":8888"))
}
