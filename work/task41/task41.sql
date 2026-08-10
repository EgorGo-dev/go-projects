SELECT 
    name,
    REPLACE(name, ' ', '_') AS clean_name,
    CAST(base_price * 1.1 * 0.85 * 100 + 0.5 + 0.000001 AS INTEGER) / 100.0 AS final_price
FROM 
    products 
ORDER BY 
    base_price * 1.1 * 0.85 DESC
LIMIT 
    20;