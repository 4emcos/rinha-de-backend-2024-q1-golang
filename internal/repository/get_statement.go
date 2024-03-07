package repository

import (
	"context"
	"rinha-de-backend-2024-q1-golang/internal/database"
	"rinha-de-backend-2024-q1-golang/models"
	"time"
)

const (
	getUserBalanceQuery      = "SELECT u.initial_balance, u.limit_in_cents FROM rinha.users u WHERE u.id = $1"
	getUserTransactionsQuery = "SELECT h.value, h.type, h.description, h.do_at FROM rinha.history h WHERE h.user_id = $1 ORDER BY id DESC LIMIT 10"
)

func GetStatement(id int32, db database.Pgx) (models.Statement, error) {
	var balance models.Balance
	transactions := make([]models.Transaction, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := db.QueryRow(ctx, getUserBalanceQuery, id).Scan(&balance.Total, &balance.Limit)
	if err != nil {
		return models.Statement{}, err
	}

	balance.StatementDate = time.Now()

	rows, err := db.Query(ctx, getUserTransactionsQuery, id)
	if err != nil {
		return models.Statement{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction = models.Transaction{}
		err = rows.Scan(&transaction.Value, &transaction.Type, &transaction.Description, &transaction.TransactionDate)
		if err != nil {
			return models.Statement{}, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return models.Statement{}, err
	}

	return models.Statement{
		Balance:      balance,
		Transactions: transactions,
	}, nil
}
