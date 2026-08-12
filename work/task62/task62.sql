SELECT
    c.id AS customer_id,
    c.full_name AS customer_name
FROM customers c
LEFT JOIN sales s ON c.id = s.customer_id
WHERE s.customer_id IS NULL;