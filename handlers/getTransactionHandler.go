package handlers

import (
	"net/http"
	"strconv"
	"transaction-app/models"

	"github.com/labstack/echo/v4"
)

func GetTransaction(c echo.Context) error {
	transactionId, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid transaction Id")
	}
	transaction, exists := models.Transactions[int(transactionId)]
	if !exists {
		return c.JSON(http.StatusNotFound, "Transaction not found")
	}
	return c.JSON(http.StatusOK, transaction)
}
