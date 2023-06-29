package main

import (
	controller "demoproject/Controllers"
	TournamentController "demoproject/Controllers/TournamentController"
	UserController "demoproject/Controllers/UserController"
	migrate "demoproject/Migrations"

	"github.com/labstack/echo"
)

func main() {

	if controller.ConnectionError != nil {
		panic("Faild to connect database.")
	}

	migrate.Migrate(controller.Connection)
	migrate.Seeder(controller.Connection)

	e := echo.New()

	{
		e.GET("/users", UserController.Index)
		e.GET("/user/:id", UserController.Show)
		e.POST("/user", UserController.Create)
		e.PUT("/user/:id", UserController.Update)
		e.DELETE("/user/:id", UserController.Delete)
	}

	{
		e.GET("/tournaments", TournamentController.Index)
		e.GET("/tournament/:id", TournamentController.Show)
		e.POST("/tournament", TournamentController.Create)
		e.PUT("/tournament/:id", TournamentController.Update)
		e.DELETE("/tournament/:id", TournamentController.Delete)
	}

	e.Logger.Fatal(e.Start(":1324"))
}
