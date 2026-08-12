SELECT 
	id,
    store_id,
    customer_id,
    payment_method_id,
    sale_datetime,
    receipt_no
FROM sales
WHERE 
	(UPPER(receipt_no) LIKE 'MSK%'
	OR UPPER(receipt_no) LIKE 'SPB%'
	OR UPPER(receipt_no) LIKE 'KZN%')
	AND sale_datetime BETWEEN '2025-01-01' AND '2025-12-31 23:59:59'
	AND UPPER(receipt_no) NOT LIKE '%TEST%';