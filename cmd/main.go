package main

import (
	"flowFinance/internal/database"
	"flowFinance/internal/handler"
	"flowFinance/internal/repository"
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

	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Can't create database: %w", err)
	}

	transactionRepository := repository.NewTransactionRepository(db)

	transactionService := service.NewTransactionService(transactionRepository)

	transactionHandler := handler.NewTransactionHandler(transactionService)

	app.Post("/transactions", transactionHandler.CreateTransaction)

	log.Fatal(app.Listen(":3000"))
}
