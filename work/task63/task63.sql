SELECT
    s.id AS sale_id,
    s.sale_datetime,
    ROUND(SUM(si.unit_price * si.quantity - si.discount_amount), 2) AS total_amount
FROM sales s
JOIN sale_items si ON s.id = si.sale_id
GROUP BY s.id, s.sale_datetime
ORDER BY s.sale_datetime DESC;