package main

import (
	"binder/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "views/public")

	// initiate routes
	routes.InitWebRoutes(e)

	e.Logger.Fatal(e.Start(":8888"))
}
