SELECT
	tx_type,
	SUM(CASE 
    	WHEN comment LIKE '%#%' 
    	OR comment LIKE '%баллов%'
    	OR comment LIKE '%награду%' 
    	THEN 1
    END) AS tx_count 
FROM loyalty_point_transactions
GROUP BY tx_type
HAVING tx_count IS NOT NULL
ORDER BY tx_count DESC;