package repository

import (
	"context"
	"flowFinance/internal/models"
	"strconv"

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

func (tp *TransactionRepository) GetAllTransactions(transactionsPtr *[]models.Transaction) error {
	sqlStr := "SELECT amount, description, category FROM transactions"
	i := 0

	rows, err := tp.db.Query(context.Background(), sqlStr)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var amountBytes []byte
		var descriptionBytes []byte
		var categoryBytes []byte

		err := rows.Scan(&amountBytes, &descriptionBytes, &categoryBytes)
		if err != nil {
			return err
		}

		amountStr := string(amountBytes)
		(*transactionsPtr)[i].Amount, err = strconv.ParseFloat(amountStr, 64)

		(*transactionsPtr)[i].Description = string(descriptionBytes)

		(*transactionsPtr)[i].Category = string(categoryBytes)
		i++

	}

	err = rows.Err()
	return err
}
