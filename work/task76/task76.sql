WITH sale_totals AS (
    SELECT
        s.store_id,
        s.id AS sale_id,
        ROUND(SUM(si.quantity * si.unit_price - si.discount_amount), 2) AS sale_amount
    FROM sales s
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY s.store_id, s.id
)
SELECT
    store_id,
    sale_id,
    sale_amount,
    RANK() OVER (PARTITION BY store_id ORDER BY sale_amount DESC) AS rank_in_store
FROM sale_totals
ORDER BY store_id, rank_in_store;