package service

import (
	"flowFinance/internal/repository"
	"strings"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(newRepo *repository.TransactionRepository) TransactionService {
	return TransactionService{
		repo: newRepo,
	}
}

func (s *TransactionService) DetectCategory(desc string) string {
	desc = strings.ToLower(desc)

	switch {
	case strings.Contains(desc, "pizza"):
		return "food"

	case strings.Contains(desc, "uber"):
		return "transport"
	default:
		return "other"
	}
}
