SELECT 
    c.name AS category_name,
    ROUND(SUM(si.unit_price * si.quantity - si.discount_amount), 2) AS total_revenue
FROM sale_items si 
INNER JOIN products p ON si.product_id = p.id 
INNER JOIN categories c ON p.category_id = c.id
GROUP BY c.name
ORDER BY total_revenue DESC;