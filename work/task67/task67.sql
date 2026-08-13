SELECT 
    store_id,
    COUNT(*) AS sales_count
FROM sales
GROUP BY store_id
HAVING COUNT(*) > (
    SELECT AVG(cnt) 
    FROM (
        SELECT COUNT(*) AS cnt 
        FROM sales 
        GROUP BY store_id
    ) AS store_counts
);