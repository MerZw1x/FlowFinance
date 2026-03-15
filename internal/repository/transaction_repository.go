package repository

import (
	"context"
	"flowFinance/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(newDb *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		db: newDb,
	}
}

func (tp *TransactionRepository) CreateTransaction(transaction models.Transaction) error {
	sqlStr := `INSERT INTO transactions(amount, description,
			category) VALUES ($1, $2, $3)`
	tp.db.Exec(context.Background(), sqlStr, transaction.Amount, transaction.Description, transaction.Category)
}
