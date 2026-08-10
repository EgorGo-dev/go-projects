SELECT 
    name, 
    base_price, base_price * 1.1 AS price_with_markup
FROM 
    products
ORDER BY 
    price_with_markup DESC, name ASC
LIMIT 
    20;
