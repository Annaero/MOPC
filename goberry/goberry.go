package main

import (
	"github.com/Annaero/MOPC/goberry/router"
	"github.com/labstack/echo/v4"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

func main() {

	e := echo.New()

	// Get router with all routes.
	router.GetRoutes(e)

	// Run server instance.
	e.Logger.Fatal(e.Start(":3000"))
}
