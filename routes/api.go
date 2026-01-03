package routes

import (
	"assets-management-api/controllers"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, firebaseApp *firebase.App) {
	stockCtrl := controllers.NewStockController(firebaseApp)

	api := app.Group("/api/v1")

	stocks := api.Group("/stocks")
	stocks.Post("/", stockCtrl.SaveStocks) // บันทึกหุ้น
}
