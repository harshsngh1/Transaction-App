package handlers

import (
	"net/http"
	"strconv"
	"transaction-app/models"

	"github.com/labstack/echo/v4"
)

func GetSum(c echo.Context) error {

	transactionId, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid transaction ID"})
	}

	var sum float64
	var queue []int64
	processed := make(map[int64]bool)
	queue = append(queue, transactionId)

	for len(queue) > 0 {
		currentId := queue[0]
		queue = queue[1:]
		if processed[currentId] {
			continue
		}

		processed[currentId] = true

		transaction, exists := models.Transactions[int(currentId)]
		if !exists {
			continue
		}
		sum += transaction.Amount

		for key, value := range models.Transactions {
			if value.ParentId != nil && *value.ParentId == currentId {
				queue = append(queue, int64(key))
			}
		}
	}
	return c.JSON(http.StatusOK, map[string]float64{"sum": sum})
	/*
		Consideration: If a transaction ID doesn’t exist, the function currently just continues without processing it.
		Depending on your business logic, you might want to return an error if the initial transaction ID doesn’t exist, like 404 Not Found.
		currently we are returning sum = 0, which is not appropriate.
	*/

}
