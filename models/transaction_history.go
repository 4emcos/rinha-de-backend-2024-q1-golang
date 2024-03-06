package models

type TransactionHistory struct {
	Value       int64
	Type        string
	Description string
	Timestamp   *string
}
