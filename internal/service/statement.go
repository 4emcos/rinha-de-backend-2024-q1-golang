package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rinha-de-backend-2024-q1-golang/internal/database"
	"rinha-de-backend-2024-q1-golang/internal/repository"
	"strconv"
)

func Statement(c *gin.Context, db database.Pgx) {
	id := c.Param("id")
	numberId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	statement, err := repository.GetStatement(int32(numberId), db)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, statement)
}
