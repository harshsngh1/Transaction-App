package handlers

import (
	"net/http"
	"strconv"
	"transaction-app/models"

	"github.com/labstack/echo/v4"
)

func AddTransaction(c echo.Context) error {
	transactionId, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid transaction Id") //Use 400 Bad Request (http.StatusBadRequest) because the client provided an invalid input (the transaction ID).
	}
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	models.Transactions[int(transactionId)] = transaction
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
