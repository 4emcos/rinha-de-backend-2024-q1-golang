package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rinha-de-backend-2024-q1-golang/internal/database"
	"rinha-de-backend-2024-q1-golang/internal/repository"
	"rinha-de-backend-2024-q1-golang/models"
	"strconv"
)

func NewTransaction(c *gin.Context, db database.Pgx) {
	input := &models.TransactionRequest{}
	id := c.Param("id")

	if err := c.ShouldBindJSON(input); err != nil || id == "" {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	numberId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	transaction, err, status := repository.UpdateBalance(int32(numberId), *input, db)

	if !status {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, transaction)
}
