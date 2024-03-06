package repository

import (
	"context"
	"rinha-de-backend-2024-q1-golang/internal/database"
	"rinha-de-backend-2024-q1-golang/models"
	"time"
)

const (
	debitQuery  = "SELECT * FROM rinha.debit($1, $2, $3)"
	creditQuery = "SELECT * FROM rinha.credit($1, $2, $3)"
)

func UpdateBalance(id int32, transaction models.TransactionRequest, db database.Pgx) (models.TransactionSuccessResponse, error, bool) {
	var newLimit int64
	var balance int64
	var success bool

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error
	if transaction.Type == "d" {
		err = db.QueryRow(ctx, debitQuery, id, transaction.Value, transaction.Description).Scan(&newLimit, &success, &balance)
	} else {
		err = db.QueryRow(ctx, creditQuery, id, transaction.Value, transaction.Description).Scan(&newLimit, &success, &balance)
	}

	if err != nil {
		//if errors.Is(err, pgx.ErrNoRows) {
		//	return models.TransactionSuccessResponse{}, nil, 100
		//}
		return models.TransactionSuccessResponse{}, err, false
	}

	if !success {
		return models.TransactionSuccessResponse{}, nil, false
	}

	return models.TransactionSuccessResponse{
		Limit:   balance,
		Balance: newLimit,
	}, nil, true
}
