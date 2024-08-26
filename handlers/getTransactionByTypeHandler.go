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
	/*
	 If no transactions of the specified type are found, you might want to return an empty list with a 200 OK.
	 Alternatively, you could return 204 No Content if you wish to signify that the request was successful, but no data is available.
	*/
}
