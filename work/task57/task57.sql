SELECT 
    account_id,
    COUNT(*) AS tx_count,
    SUM(points_delta) AS total_delta
FROM loyalty_point_transactions
GROUP BY account_id
HAVING COUNT(*) >= 3
ORDER BY total_delta DESC;