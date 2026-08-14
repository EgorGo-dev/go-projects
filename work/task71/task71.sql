WITH unused_coupons AS (
    SELECT
    	c.id AS coupon_id,
    	c.code
FROM coupons c
WHERE c.status = 'ISSUED' 
    AND NOT EXISTS (
    	SELECT 1
    	FROM sale_coupons sc
    	WHERE sc.coupon_id = c.id
    ) 
)
SELECT 
	coupon_id,
	code
FROM unused_coupons
ORDER BY coupon_id;