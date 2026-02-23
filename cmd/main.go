package main

import (
	"flowFinance/internal/handler"
	"flowFinance/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	transactionService := &service.TransactionService{}
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app.Post("/transactions", transactionHandler.CreateTransaction)

	app.Listen(":3000")
}
