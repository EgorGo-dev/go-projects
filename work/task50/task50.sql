SELECT 
    COUNT(*) AS items_count,
    ROUND(SUM(quantity * unit_price), 2) AS total_gross,
    ROUND(SUM(discount_amount), 2) AS total_discount,
    ROUND(SUM(quantity * unit_price - discount_amount), 2) AS total_net,
    ROUND(AVG(quantity * unit_price - discount_amount), 2) AS avg_line_net
FROM sale_items;