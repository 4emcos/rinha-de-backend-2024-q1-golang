package models

type TransactionRequest struct {
	Value       int64  `json:"valor" binding:"required,min=1"`
	Type        string `json:"tipo" binding:"required,max=1,oneof=d c"`
	Description string `json:"descricao" binding:"required,max=10"`
}
