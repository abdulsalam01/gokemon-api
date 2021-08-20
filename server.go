package main

import (
	"gokemon/abdulsalam/db"
	"gokemon/abdulsalam/route"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	server := echo.New()

	db.Init()
	route.InitServer(server).Init()

	server.Use(middleware.CORS())
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	server.Logger.Fatal(server.Start(":1323"))
}
