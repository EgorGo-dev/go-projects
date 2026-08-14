WITH store_revenue AS (
    SELECT 
        s.store_id,
        SUM(si.unit_price * si.quantity - si.discount_amount) AS total_revenue_raw
    FROM sales s
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY s.store_id
)
SELECT 
    store_id,
    ROUND(total_revenue_raw, 2) AS total_store_amount,
    RANK() OVER (ORDER BY total_revenue_raw DESC) AS store_rank_by_revenue
FROM store_revenue
ORDER BY store_rank_by_revenue;