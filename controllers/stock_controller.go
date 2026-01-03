package controllers

import (
	"assets-management-api/models"
	"assets-management-api/repository"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

type StockController struct {
	FirebaseApp *firebase.App
}

func NewStockController(app *firebase.App) *StockController {
	return &StockController{FirebaseApp: app}
}

func (ctrl *StockController) SaveStocks(c *fiber.Ctx) error {
	var req models.SaveStocksRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ข้อมูล JSON ไม่ถูกต้อง",
		})
	}

	err := repository.SaveStocksToFirestore(ctrl.FirebaseApp, req.Stocks)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "ไม่สามารถบันทึกลงฐานข้อมูลได้",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "บันทึกข้อมูลเรียบร้อยแล้ว",
	})
}
