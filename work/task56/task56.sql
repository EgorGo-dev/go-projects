SELECT 
    id,
    account_id,
    sale_id,
    tx_type,
    points_delta,
    tx_datetime,
    comment
FROM loyalty_point_transactions
WHERE 
	(UPPER(tx_type) LIKE 'ADJUST' OR UPPER(tx_type) LIKE 'EXPIRE')
	AND sale_id IS NULL
	AND (comment LIKE '%Корректировка баллов%' OR comment LIKE '%Сгорание баллов%');