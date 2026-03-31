package handler

import (
	"flowFinance/internal/models"
	"flowFinance/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var transaction models.Transaction

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(400).SendString("invalid request")
	}

	transaction.Category = h.service.DetectCategory(transaction.Description)

	err := h.service.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return c.JSON(transaction)
}

func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	filter := models.TransactionFilters{
		Category:    c.Query("category"),
		Description: c.Query("description"),
		MinAmount:   c.QueryFloat("minAmount"),
		MaxAmount:   c.QueryFloat("maxAmount"),
	}

	transactions, err := h.service.GetAllTransactions(filter)
	if err != nil {
		return err
	}
	return c.JSON(transactions)
}
