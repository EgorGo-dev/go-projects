SELECT 
    email,
    LOWER(email) AS email_lower,
    LOWER(SUBSTR(email, INSTR(email, '@') + 1)) AS domain,
    LOWER(SUBSTR(email, 1, INSTR(email, '@') - 1)) AS login,
    REPLACE(LOWER(SUBSTR(email, 1, INSTR(email, '@') - 1)), '.', '_') 
        || '@' || LOWER(SUBSTR(email, INSTR(email, '@') + 1)) AS masked,
    LENGTH(SUBSTR(email, 1, INSTR(email, '@') - 1)) AS login_len
FROM 
    customers
WHERE 
    email IS NOT NULL AND email != ''
ORDER BY 
    domain ASC,
    login_len DESC;