package routes

import (
	"transaction-app/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.PUT("/transactionservice/transaction/:transaction_id", handlers.AddTransaction)
	e.GET("/transactionservice/transaction/:transaction_id", handlers.GetTransaction)
	e.GET("/transactionservice/types/:type", handlers.GetTransactionByType)
	e.GET("/transactionservice/sum/:transaction_id", handlers.GetSum)
}
