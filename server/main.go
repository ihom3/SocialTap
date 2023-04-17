package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//connect to DB
	database.Connect()

	//Create app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:4200",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Language",
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
