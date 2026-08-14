SELECT
    customer_id,
    id AS sale_id,
    sale_datetime,
    ROW_NUMBER() OVER (PARTITION BY customer_id ORDER BY sale_datetime) AS purchase_number
FROM sales
ORDER BY customer_id, sale_datetime;