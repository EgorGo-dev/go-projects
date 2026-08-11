SELECT 
    id,
    code,
    coupon_type,
    value,
    start_at,
    end_at,
    ROUND(value * 1.2, 2) AS value_with_tax,
    LENGTH(code) AS code_len,
    code || ' (' || value || '%)' AS label
FROM coupons 
WHERE 
    status IN ('ISSUED', 'USED')
    AND end_at BETWEEN '2025-01-01' AND '2025-12-31'
    AND value BETWEEN 5 AND 50
ORDER BY 
    end_at ASC,
    value DESC
LIMIT 20;