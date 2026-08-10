SELECT
    sale_id,
    product_id,
    quantity,
    unit_price,
    discount_amount,
    quantity * unit_price AS gross_total,
    (quantity * unit_price) - discount_amount AS net_total,
    ROUND((discount_amount / (quantity * unit_price)) * 100) AS discount_percent
FROM 
    sale_items 
ORDER BY 
    discount_percent DESC, 
    net_total DESC
LIMIT 
    20 OFFSET 20;