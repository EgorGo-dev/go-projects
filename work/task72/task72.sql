WITH sale_amounts AS (
    SELECT 
        s.id AS sale_id,
        s.customer_id,
        SUM(si.unit_price * si.quantity - si.discount_amount) AS raw_sale_amount
    FROM sales s
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY s.id, s.customer_id
)
SELECT 
    customer_id,
    sale_id,
    ROUND(raw_sale_amount, 2) AS sale_amount,
    ROUND(AVG(raw_sale_amount) OVER (PARTITION BY customer_id), 2) AS avg_amount_per_customer,
    ROUND(raw_sale_amount - AVG(raw_sale_amount) OVER (PARTITION BY customer_id), 2) AS diff_from_avg
FROM sale_amounts
ORDER BY customer_id, sale_amount DESC;