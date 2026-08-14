WITH next_dates AS (
    SELECT
        customer_id,
        id AS sale_id,
        sale_datetime,
        LEAD(sale_datetime) OVER (PARTITION BY customer_id ORDER BY sale_datetime) AS next_sale_datetime
    FROM sales
)
SELECT
    customer_id,
    sale_id,
    sale_datetime,
    next_sale_datetime,
    ROUND(JULIANDAY(next_sale_datetime) - JULIANDAY(sale_datetime), 0) AS days_to_next_purchase
FROM next_dates
ORDER BY customer_id, sale_datetime;