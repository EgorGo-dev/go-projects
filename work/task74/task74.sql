WITH sale_amounts AS (
    SELECT 
        s.customer_id,
        s.id AS sale_id,
        SUM(si.unit_price * si.quantity - si.discount_amount) AS sale_amount
    FROM sales s
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY s.customer_id, s.id
)
SELECT 
    customer_id,
    sale_id,
    sale_amount,
    SUM(sale_amount) OVER (PARTITION BY customer_id) AS customer_total_amount,
    ROUND(
        sale_amount * 100.0 / SUM(sale_amount) OVER (PARTITION BY customer_id), 
        2
    ) AS sale_share_percent
FROM sale_amounts
ORDER BY customer_id, sale_amount DESC;