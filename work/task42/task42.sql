SELECT 
    email,
	LOWER(email) AS email_lower,
	LOWER(SUBSTR(email, 1, INSTR(email, '@') - 1)) AS login
FROM 
    customers
ORDER BY
	login ASC;
