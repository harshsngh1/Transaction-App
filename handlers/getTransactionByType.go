package handlers

import (
	"net/http"
	"transaction-app/models"

	"github.com/labstack/echo/v4"
)

func GetTransactionByType(c echo.Context) error {
	transactionType := c.Param("type")

	var transactionIds []int

	for key, value := range models.Transactions {
		if value.Type == transactionType {
			transactionIds = append(transactionIds, key)
		}
	}
	return c.JSON(http.StatusOK, transactionIds)
}
