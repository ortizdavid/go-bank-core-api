-- view_customer_data
DROP VIEW IF EXISTS view_customer_data;
CREATE VIEW view_customer_data AS 
SELECT * FROM customers;


-- view_account_data
DROP VIEW IF EXISTS view_account_data;
CREATE VIEW view_account_data AS 
SELECT * FROM accounts;


-- view_transaction_data
DROP VIEW IF EXISTS view_transaction_data;
CREATE VIEW view_transaction_data AS 
SELECT 
    tr.transaction_id, tr.unique_id,
    tr.code, tr.amount,
    tr.balance_before,tr.balance_after,
    tr.currency, tr.description,
    tr.transaction_date,
    tr.created_at, tr.updated_at,
    acc.account_id, acc.account_number,
    acc.iban,
    cu.customer_name, cu.identification_number,
    tt.type_id, tt.type_name,
    ts.status_id, ts.status_name
FROM transactions tr
JOIN accounts acc ON(acc.account_id = tr.account_id)
JOIN customers cu ON(cu.customer_id = acc.customer_id)
JOIN transaction_type tt ON(tt.type_id = tr.transaction_type_id)
JOIN transaction_status ts ON(ts.status_id = tr.transaction_status_id)
ORDER BY tr.created_at DESC;


