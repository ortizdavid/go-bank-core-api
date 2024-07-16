-- view_customer_data
DROP VIEW IF EXISTS view_customer_data;
CREATE VIEW view_customer_data AS 
SELECT 
    cu.customer_id, cu.unique_id,
    cu.customer_name, cu.identification_number,
    cu.gender, cu.birth_date,
    cu.email, cu.phone,
    cu.address,
    cu.created_at, cu.updated_at,
    ct.type_id, ct.type_name,
    cs.status_id, cs.status_name
FROM customers cu
JOIN customer_type ct ON (ct.type_id = cu.customer_type)
JOIN customer_status cs ON (cs.status_id = cu.customer_status);


-- view_account_data
DROP VIEW IF EXISTS view_account_data;
CREATE VIEW view_account_data AS 
SELECT 
    acc.account_id, acc.unique_id,
    acc.account_number, acc.iban,
    acc.balance, acc.currency,
    acc.created_at, acc.updated_at,
    cu.customer_id, cu.customer_name, 
    cu.identification_number,
    act.type_id, act.type_name,
    acs.status_id, acs.status_name
FROM accounts acc
JOIN customers cu ON (cu.customer_id = acc.customer_id)
JOIN account_status acs ON (acs.status_id = acc.account_status)
JOIN account_type act ON (act.type_id = acc.account_type)
ORDER BY acc.created_at DESC;


-- view_transaction_data
DROP VIEW IF EXISTS view_transaction_data;
CREATE VIEW view_transaction_data AS 
SELECT 
    tr.transaction_id, tr.unique_id,
    tr.code, tr.amount,
    tr.balance_before, tr.balance_after,
    tr.currency, tr.description,
    tr.transaction_date,
    tr.created_at, tr.updated_at,
    src.account_id AS source_account_id, 
    src.account_number AS source_account_number, 
    src.iban AS source_iban,
    dst.account_id AS destination_account_id, 
    dst.account_number AS destination_account_number, 
    dst.iban AS destination_iban,
    cu.customer_name, cu.identification_number,
    tt.type_id, tt.type_name,
    ts.status_id, ts.status_name
FROM transactions tr
JOIN accounts src ON(src.account_id = tr.source_id)
JOIN accounts dst ON(dst.account_id = tr.destination_id)
JOIN customers cu ON(cu.customer_id = src.customer_id)
JOIN transaction_type tt ON(tt.type_id = tr.transaction_type_id)
JOIN transaction_status ts ON(ts.status_id = tr.transaction_status_id)
ORDER BY tr.created_at DESC;

