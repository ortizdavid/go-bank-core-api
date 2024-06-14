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

- Create Account with Customer
    ```http
    POST /api/accounts/create-with-customer
    ```
    ```json
    {
        "customer_name": "Anna Maria",
        "identification_number": "05455T5F644",
        "email": "ana@gmail.com",
        "phone": "+294678902348",
        "address": "Luanda, Angola",
        "account_type_id": 1,
        "currency": "USD"
    }
    ```

- Create Account And Associate Customer
    ```http
    POST /api/accounts
    ```
    ```json
    {
        "customer_id": 5,
        "account_type": 1,
        "currency": "USD"
    }
    ```

- Change Account Status
    ```http
    PUT /api/accounts/change-status
    ```
    ```json
    {
        "account_number": "8792529764",
        "account_status": 6
    }
    ```

- Change Account Type
    ```http
    PUT /api/accounts/change-type
    ```
    ```json
    {
        "account_number": "8792529764",
        "account_type": 1
    }
    ```