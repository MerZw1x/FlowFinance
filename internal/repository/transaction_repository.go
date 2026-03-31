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

	_, err := tp.db.Exec(context.Background(),
		sqlStr, transaction.Amount, transaction.Description, transaction.Category)

	return err
}

func (tp *TransactionRepository) GetAllTransactions(filter models.TransactionFiltets) ([]models.Transaction, error) {
	sqlStr := "SELECT amount, description, category FROM transactions"

	rows, err := tp.db.Query(context.Background(), sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Transactions []models.Transaction

	for rows.Next() {
		var subtr models.Transaction

		err := rows.Scan(&subtr.Amount, &subtr.Description, &subtr.Category)
		if err != nil {
			return nil, err
		}

		Transactions = append(Transactions, subtr)

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return Transactions, nil
}
