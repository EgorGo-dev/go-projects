SELECT 
	COUNT(DISTINCT ROUND(base_price / 10) * 10) AS distinct_price_levels,
	MIN(ROUND(base_price/10)*10) AS min_level,
	MAX(ROUND(base_price/10)*10) AS max_level
FROM products
	