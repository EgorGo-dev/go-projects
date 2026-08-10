SELECT 
    ROUND(SUM(LENGTH(name) * base_price) / 
    SUM(base_price), 2) AS weighted_avg_name_length
FROM 
    products;