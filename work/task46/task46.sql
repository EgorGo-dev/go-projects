SELECT 
	COUNT(*) AS total_products,
	MIN(base_price) AS min_price,
	MAX(base_price) AS max_price,
	ROUND(AVG(base_price), 2) AS avg_price,
	ROUND(AVG(base_price*1.2), 2) AS avg_price_with_tax
FROM 
	products;
