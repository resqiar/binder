package main

import (
	"binder/configs"
	"binder/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	// database connection
	if err := configs.InitDBCon(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	store := configs.InitiateSession()
	e.Use(session.Middleware(*store))
	e.Static("/static", "views/public")

	// initiate configs
	configs.InitGoogleConfig()

	// initiate routes
	routes.InitWebRoutes(e)
	routes.InitAPIRoutes(e)

	e.Logger.Fatal(e.Start(":8888"))
}
