SELECT 
	COUNT(name) AS products_total,
	COUNT(DISTINCT SUBSTR(name, 1, 1)) AS distinct_first_letters,
	ROUND(AVG(LENGTH(name)), 1) AS avg_name_length,
	MAX(LENGTH(name)) AS max_name_length
FROM 
	products;