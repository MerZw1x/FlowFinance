package service

import "strings"

type TransactionService struct{}

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
