SELECT 
	name,
	SUBSTR(name, 1, 12) AS title,
	ROUND(base_price, 2) AS base_price_rounded,
	ROUND(base_price * 1.2, 2) AS tax_price,
	ROUND(base_price * 1.2 * 0.93, 2) AS promo_price,
	SUBSTR(name, 1, 12) ||': '|| ROUND(base_price * 1.2 * 0.93, 2) AS card_label
FROM 
    products
ORDER BY 
	promo_price DESC 
LIMIT 
	15;