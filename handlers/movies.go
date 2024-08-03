package handlers

import (
	"time"

	"github.com/YearZeroOne/go-movies/database"
	"github.com/YearZeroOne/go-movies/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetMovies(c *fiber.Ctx) error {
	movies := []models.Movie{}
	database.DB.Db.Find(&movies)

	return c.Status(200).JSON(movies)
}

func GetMovieById(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movie

	database.DB.Db.First(&movie, id)
	if movie.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Movie not found"})
	}
	return c.Status(200).JSON(movie)
}

func DeleteMovieById(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movie
	database.DB.Db.First(&movie, id)
	if movie.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Movie not found"})
	}
	database.DB.Db.Delete(&movie)
	return c.SendStatus(204)
}

func CreateMovie(c *fiber.Ctx) error {
	var movie models.Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if movie.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Title required"})
	}
	if movie.Genre == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Genre required"})
	}

	database.DB.Db.Create(&movie)
	return c.Status(200).JSON(movie)
}

func EditMovie(c *fiber.Ctx) error {
	id := c.Params("id")

	var movie models.Movie
	if err := database.DB.Db.First(&movie, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Movie not found"})
	}

	var updatedData models.Movie
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if updatedData.Title != "" {
		movie.Title = updatedData.Title
	}
	if updatedData.ImageURL != nil {
		movie.ImageURL = updatedData.ImageURL
	}
	if updatedData.Genre != "" {
		movie.Genre = updatedData.Genre
	}
	if updatedData.Release != nil {
		movie.Release = updatedData.Release
	}

	if err := database.DB.Db.Save(&movie).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update movie"})
	}

	return c.Status(200).JSON(movie)
}

var jwtSecret = []byte("your_secret_key")

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Cannot parse JSON"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Password: string(password),
	}

	database.DB.Db.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Cannot parse JSON"})
	}

	var user models.User
	database.DB.Db.Where("username = ?", data["username"]).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Could not log in"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}
