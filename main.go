package main

import (
	"asset-management/config"
	"asset-management/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Init Firebase
	firebaseApp := config.InitFirebase()

	// เรียกใช้ Route Module
	routes.SetupRoutes(app, firebaseApp)

	app.Listen(":8080")
}