SELECT DISTINCT 
    tx_type || ':' || points_delta AS profile,
    ABS(points_delta) AS abs_delta
FROM loyalty_point_transactions
ORDER BY abs_delta DESC, profile ASC
LIMIT 25;