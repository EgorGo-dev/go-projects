--LENGTH(column_name)
SELECT 
    code AS method_code,
    name AS method_name,
    LENGTH(code) AS code_len,
    code || ' - ' || name AS label
FROM 
    payment_methods
ORDER BY 
    code_len DESC, 
    method_code ASC;