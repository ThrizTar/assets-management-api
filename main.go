package main

import (
	"assets-management-api/config"
	"assets-management-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Init Firebase
	firebaseApp := config.InitFirebase()

	// เรียกใช้ Route Module
	routes.SetupRoutes(app, firebaseApp)

	app.Listen(":8080")
}
