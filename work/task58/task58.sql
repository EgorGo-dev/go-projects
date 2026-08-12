SELECT 
    store_id,
    COUNT(*) AS sales_count,
    COUNT(DISTINCT customer_id) AS customers_count,
    CASE 
        WHEN COUNT(DISTINCT customer_id) <= 20 THEN 'Low'
        WHEN COUNT(DISTINCT customer_id) <= 50 THEN 'Medium'
        ELSE 'High'
    END AS store_level
FROM sales
GROUP BY store_id
HAVING COUNT(*) > 0
ORDER BY sales_count DESC;