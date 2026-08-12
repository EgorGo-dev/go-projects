SELECT 
    city,
    COUNT(*) AS total_active_stores,
    (SELECT COUNT(DISTINCT sale.store_id) 
     FROM sales sale 
     WHERE sale.store_id IN (SELECT id FROM stores WHERE is_active = 1 AND city = s.city)
    ) AS stores_with_sales,
    ROUND(
        (SELECT COUNT(DISTINCT sale.store_id) 
         FROM sales sale 
         WHERE sale.store_id IN (SELECT id FROM stores WHERE is_active = 1 AND city = s.city)
        ) * 1.0 / COUNT(*), 
        2
    ) AS activity_ratio
FROM stores s
WHERE s.is_active = 1
GROUP BY s.city
ORDER BY activity_ratio DESC;