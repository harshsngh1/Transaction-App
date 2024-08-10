# Transaction-App

This is a Simple Transaction app to demo some APIs written in Golang using Echo framework

## Technologies Used
- Language : Golang
- Framework : Echo Framework
- Database : In-memory

## Project Structure

```
Transaction-App/
│
├── main.go
├── go.mod
├── go.sum
├── handlers/
│   └── addTransactionHandler.go
│   └── getTransactionsHandler.go
│   └── getTransactionByTypeHandler.go
│   └── sumTransactionHandler.go
├── routes/
│   └── routes.go
├── models/
│   └── transaction.go
└── tests/
    └── transaction_test.go
```
## Installation

1. Clone Repo
```
https://github.com/harshsngh1/Transaction-App
```

2. Run this command to start server at post 8080
```
go run main.go
```

## Testing cURL commands
### Testing PUT endpoint /transactionservice/transaction/:transaction_id 
```
curl -X PUT http://localhost:8080/transactionservice/transaction/1 -H "Content-Type: application/json" -d '{"amount": 1000, "type":"parent"}' 
curl -X PUT http://localhost:8080/transactionservice/transaction/2 -H "Content-Type: application/json" -d '{"amount": 2000, "type":"child", "parent_id": 1}'
curl -X PUT http://localhost:8080/transactionservice/transaction/3 -H "Content-Type: application/json" -d '{"amount": 3000, "type":"child", "parent_id": 1}'
curl -X PUT http://localhost:8080/transactionservice/transaction/4 -H "Content-Type: application/json" -d '{"amount": 1500, "type":"grandchild", "parent_id": 3}'
curl -X PUT http://localhost:8080/transactionservice/transaction/5 -H "Content-Type: application/json" -d '{"amount": 2500, "type":"grandchild", "parent_id": 3}'
```
### Testing GET endpoint /transactionservice/transaction/:transaction_id
```
curl -X GET http://localhost:8080/transactionservice/transaction/1
```
### Testing GET endpoint /transactionservice/types/:type
```
curl -X GET http://localhost:8080/transactionservice/parent
```
### Testing GET endpoint /transactionservice/sum/:transaction_id
```
curl -X GET http://localhost:8080/transactionservice/sum/1
```