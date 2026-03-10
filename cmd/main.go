package main

import (
	"flowFinance/internal/handler"
	"flowFinance/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Can't connect to config file: %v", err)
	}

	app := fiber.New()

	transactionService := &service.TransactionService{}
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app.Post("/transactions", transactionHandler.CreateTransaction)

	log.Fatal(app.Listen(":3000"))
}
