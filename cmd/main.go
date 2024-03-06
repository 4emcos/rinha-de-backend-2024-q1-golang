package main

import (
	"github.com/gin-gonic/gin"
	"rinha-de-backend-2024-q1-golang/internal/database"
	"rinha-de-backend-2024-q1-golang/internal/service"
)

func main() {

	db := database.Connect()

	router := gin.Default()
	router.POST("/clientes/:id/transacoes", func(c *gin.Context) {
		service.NewTransaction(c, db)
	})

	router.GET("/clientes/:id/extrato", func(c *gin.Context) {
		service.Statement(c, db)
	})

	router.Run(":" + "8080")
}
