package main

import (
	"transaction-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))

}
