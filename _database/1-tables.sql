\c postgres;
DROP DATABASE IF EXISTS go_bank_core_api;
CREATE DATABASE go_bank_core_api;
\c go_bank_core_api;


DROP TABLE IF EXISTS account_status;
CREATE TABLE account_status (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(20) UNIQUE NOT NULL,
    description VARCHAR(150)
);


DROP TABLE IF EXISTS account_type;
CREATE TABLE account_type (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(20) UNIQUE NOT NULL,
    description VARCHAR(150)
);


DROP TABLE IF EXISTS transaction_status;
CREATE TABLE transaction_status (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(20) UNIQUE NOT NULL,
    description VARCHAR(150)
);


DROP TABLE IF EXISTS transaction_type;
CREATE TABLE transaction_type (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(20) UNIQUE NOT NULL,
    description VARCHAR(150)
);


DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    customer_name VARCHAR(150) NOT NULL,
    identification_number VARCHAR(30) UNIQUE NOT NULL,
    email VARCHAR(150) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    address VARCHAR(200),
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
    account_id SERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    account_type INT NOT NULL,
    account_status INT NOT NULL,
    account_number VARCHAR(18) UNIQUE NOT NULL,
    iban VARCHAR(31) UNIQUE NOT NULL,
    holder_name VARCHAR(150),
    balance DECIMAL(18, 2) NOT NULL DEFAULT 0.0,
    currency VARCHAR(3),
    unique_id UUID DEFAULT gen_random_uuid(),
	created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_account_type FOREIGN KEY(account_type) REFERENCES account_type(type_id),
    CONSTRAINT fk_account_status FOREIGN KEY(account_status) REFERENCES account_status(status_id),
    CONSTRAINT fk_customer_account FOREIGN KEY(customer_id) REFERENCES customers(customer_id)
);


DROP TABLE IF EXISTS transactions;
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    transaction_type_id INT NOT NULL,
    transaction_status_id INT NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    currency VARCHAR(3),
    balance_before DECIMAL(18, 2) NOT NULL,
    balance_after DECIMAL(18, 2) NOT NULL,
    description VARCHAR(150),
    transaction_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    unique_id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_transaction_type FOREIGN KEY(transaction_type_id) REFERENCES transaction_type(type_id),
    CONSTRAINT fk_transaction_status FOREIGN KEY(transaction_status_id) REFERENCES transaction_status(status_id)
);


