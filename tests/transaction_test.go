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
		expected := `{"sum":0}`
		assert.JSONEq(t, expected, res.Body.String())
	}
}
