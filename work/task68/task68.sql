WITH store_revenue AS (
    SELECT 
        s.id AS store_id,
        s.name AS store_name,
        ROUND(SUM(si.unit_price * si.quantity - si.discount_amount)) AS total_revenue
    FROM stores s
    JOIN sales sl ON s.id = sl.store_id
    JOIN sale_items si ON sl.id = si.sale_id
    GROUP BY s.id, s.name
),
avg_revenue AS (
    SELECT AVG(total_revenue) AS avg_revenue
    FROM store_revenue
)
SELECT 
    sr.store_id,
    sr.store_name,
    sr.total_revenue
FROM store_revenue sr
WHERE sr.total_revenue > (SELECT avg_revenue FROM avg_revenue);