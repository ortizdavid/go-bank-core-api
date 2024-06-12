# go-bank-core-api
REST API with Banking System written in golang


## Tools
- PostgreSQL
- net/http
- Go


## Features
- Customers
    - [ ] Create Customer
    - [ ] Get All Customers
    - [ ] Get Account
- Accounts
    - [ ] Create Account with Customer
    - [ ] Create and Associate an Account
    - [ ] Get All Accounts
    - [ ] Get Account
    - [ ] Change Account Status
    - [ ] Change Account Type
- Transactions
    - [ ] Deposit
    - [ ] Withdrawal
    - [ ] Transfer By Account Number and Iban
    - [ ] Create Transaction
    - [ ] Get All Transactions
    - [ ] Get Account Transactions
- Reports
    - [ ] Customers Report
    - [ ] Account Reports
    - [ ] Transactions Report
- Statistics
    - [ ] Customers Statistics
    - [ ] Account Statistics
    - [ ] Transactions Report


## How to run
- Download or clone repository: `git clone https://github.com/ortizdavid/go-bank-core-api`
- Copy database scripts from [_database folder](_database) to Postgres
- Change databaase connection from [.env](.env) file
- Import Postman Collections from [_api_collections](_api_colletions)
- Download Packages `go mod tidy`
- Run Application: `go run main.go`


## Example of endpoints

- Deposit
    ```http
    POST /api/transactions/deposit
    ```
    ```json
    {
        "account_number": "8792529764",
        "amount": 127000.10,
        "currency": "USD"
    }
    ```

- Withdraw
    ```http
    POST /api/transactions/withdraw
    ```
    ```json
    {
        "account_number": "8792529764",
        "amount": 1000.95,
        "currency": "USD"
    }
    ```

- Transfer by Account Number
    ```http
    POST /api/transactions/transfer
    ```
    ```json
    {
        "source_number": "8792529764",
        "destination_number": "7840163431",
        "amount": 529.98,
        "currency": "USD"
    }
    ```

- Transfer by Iban
    ```http
    POST /api/transactions/transfer
    ```
    ```json
    {
        "source_iban": "XX741234565111434546",
        "destination_iban": "XX481234569540236342",
        "amount": 12300.98,
        "currency": "USD"
    }
    ```