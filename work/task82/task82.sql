WITH sale_amounts AS (
    SELECT
        s.customer_id,
        s.id AS sale_id,
        ROUND(SUM(si.quantity * si.unit_price - si.discount_amount), 2) AS sale_amount
    FROM sales s
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY s.customer_id, s.id
),
ranked_sales AS (
    SELECT
        customer_id,
        sale_id,
        sale_amount,
        DENSE_RANK() OVER (PARTITION BY customer_id ORDER BY sale_amount DESC) AS rank_within_customer
    FROM sale_amounts
)
SELECT
    customer_id,
    sale_id,
    sale_amount,
    rank_within_customer
FROM ranked_sales
WHERE rank_within_customer = 1
ORDER BY customer_id, sale_id;
