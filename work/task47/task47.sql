SELECT 
	COUNT(*) AS total_products,
	COUNT(DISTINCT SUBSTR(name, 1, 1)) AS distinct_first_letters
FROM 
    products;