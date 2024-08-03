package main

import (
	"github.com/YearZeroOne/go-movies/handlers"
	"github.com/gofiber/fiber/v2"
)

func MovieRoutes(app *fiber.App) {
	app.Get("/movies", handlers.GetMovies)
	app.Get("/movies/:id", handlers.GetMovieById)
	app.Post("/movies", handlers.CreateMovie)
	app.Patch("/movies/:id", handlers.EditMovie)
	app.Delete("/movies/:id", handlers.DeleteMovieById)
}

func UserRoutes(app *fiber.App) {
    app.Post("/register", handlers.Register)
    app.Post("/login", handlers.Login)
}
