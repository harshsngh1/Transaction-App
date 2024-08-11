package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"transaction-app/handlers"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddTransaction(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/transactionservice/transaction/10", bytes.NewBufferString(`{"amount": 5000, "type":"cars"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetPath("/transactionservice/transaction/:transaction_id")
	c.SetParamNames("transaction_id")
	c.SetParamValues("10")
	if assert.NoError(t, handlers.AddTransaction(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		expected := `{"status":"ok"}`
		assert.JSONEq(t, expected, res.Body.String())
	}
}

func TestSumTransaction(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/transactionservice/sum/10", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetPath("/transactionservice/sum/:transaction_id")
	c.SetParamNames("transaction_id")
	c.SetParamValues("10")
	if assert.NoError(t, handlers.GetSum(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		expected := `{"sum":5000}`
		assert.JSONEq(t, expected, res.Body.String())
	}
}

func TestGetTransaction(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/transactionservice/transaction/10", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetPath("/transactionservice/transaction/:transaction_id")
	c.SetParamNames("transaction_id")
	c.SetParamValues("10")
	if assert.NoError(t, handlers.GetTransaction(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		expected := `{"amount":5000,"type":"cars"}`
		assert.JSONEq(t, expected, res.Body.String())
	}
}

func TestGetTransactionByTest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/transactionservice/types/cars", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetPath("/transactionservice/types/:type")
	c.SetParamNames("type")
	c.SetParamValues("cars")
	if assert.NoError(t, handlers.GetTransactionByType(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		expected := `[10]`
		assert.JSONEq(t, expected, res.Body.String())
	}
}
