package models

import "time"

type Statement struct {
	Balance      Balance       `json:"saldo"`
	Transactions []Transaction `json:"ultimas_transacoes"`
}

type Balance struct {
	Total         int64     `json:"total"`
	StatementDate time.Time `json:"data_extrato"`
	Limit         int64     `json:"limite"`
}

type Transaction struct {
	Value           int64     `json:"valor"`
	Type            string    `json:"tipo"`
	Description     string    `json:"descricao"`
	TransactionDate time.Time `json:"realizada_em"`
}
