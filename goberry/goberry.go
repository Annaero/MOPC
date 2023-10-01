package main

import (
	"log"

	"github.com/Annaero/MOPC/goberry/configs"
	"github.com/Annaero/MOPC/goberry/models/database"
	"github.com/Annaero/MOPC/goberry/router"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.Debug = true
	db, err := database.ConnectMongDB(configs.EnvMongoURI())
	if err != nil {
		log.Fatal(err)
	}
	router.GetRoutes(e, db)

	// Run server instance.
	e.Logger.Fatal(e.Start(":3000"))
}
