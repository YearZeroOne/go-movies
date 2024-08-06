package main

import (
	"fmt"
	"github.com/YearZeroOne/go-movies/config"
	"github.com/YearZeroOne/go-movies/database"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/contrib/jwt"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	UserRoutes(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: config.JwtSecret},
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		SuccessHandler: func(c *fiber.Ctx) error {
			fmt.Println("Hello World!")
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	MovieRoutes(app)

	app.Listen(":3000")
}
