WITH sale_totals AS (
    SELECT
        s.customer_id,
        s.id AS sale_id,
        s.sale_datetime,
        SUM(si.quantity * si.unit_price - si.discount_amount) AS sale_amount_raw
    FROM sales s
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY s.customer_id, s.id, s.sale_datetime
)
SELECT
    customer_id,
    sale_id,
    ROUND(sale_amount_raw, 2) AS sale_amount,
    ROUND(LAG(sale_amount_raw) OVER (PARTITION BY customer_id ORDER BY sale_datetime), 2) AS previous_sale_amount,
    ROUND(sale_amount_raw - LAG(sale_amount_raw) OVER (PARTITION BY customer_id ORDER BY sale_datetime), 2) AS change_from_previous
FROM sale_totals
ORDER BY customer_id, sale_datetime;