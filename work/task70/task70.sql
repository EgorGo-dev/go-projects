WITH first_sales AS (
    SELECT 
    	customer_id,
    	MIN(sale_datetime) AS first_sale_datetime
    FROM sales
    GROUP BY customer_id
)
SELECT 
    fs.customer_id,
	fs.first_sale_datetime,
	s.store_id
FROM first_sales fs
JOIN sales s 
    ON fs.customer_id = s.customer_id
	AND fs.first_sale_datetime = s.sale_datetime
ORDER BY fs.customer_id