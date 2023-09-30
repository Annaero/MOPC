package main

import (
	"github.com/Annaero/MOPC/goberry/configs"
	"github.com/Annaero/MOPC/goberry/models/database"
	"github.com/Annaero/MOPC/goberry/router"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	mongo_client := configs.ConnectDB()
	db := database.MongoDB{Client: mongo_client}
	router.GetRoutes(e, &db)

	// Run server instance.
	e.Logger.Fatal(e.Start(":3000"))
}
