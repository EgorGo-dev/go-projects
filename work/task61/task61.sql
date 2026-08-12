SELECT 
    c.id AS customer_id,
    c.full_name AS customer_name,
    COUNT(s.id) AS sales_count
FROM customers c
INNER JOIN sales s ON c.id = s.customer_id
GROUP BY c.id, c.full_name
ORDER BY sales_count DESC, customer_id ASC;