SELECT 
    name,
	LENGTH(name) AS name_len,
	SUBSTR(name, 1, 3) AS short_code,
	SUBSTR(name, 1, 3) || ' (' || LENGTH(name) || ')' AS label
FROM 
    categories
ORDER BY 
	name_len DESC,
    name ASC;