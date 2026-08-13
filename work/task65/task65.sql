SELECT 
	c.full_name AS customer_name,
	lpt.sale_id,
	sl.sale_datetime,
	lpt.points_delta,
	lt.name AS tier_name
FROM loyalty_point_transactions lpt
INNER JOIN sales sl ON lpt.sale_id = sl.id
INNER JOIN loyalty_accounts la ON lpt.account_id = la.id
INNER JOIN customers c ON la.customer_id = c.id
INNER JOIN loyalty_tiers lt ON la.tier_id = lt.id
WHERE lpt.tx_type = 'EARN'
ORDER BY 
    lpt.points_delta DESC,
	sl.id ASC;
-- loyalty_point_transactions lpt
-- sales sl
-- customers c
-- loyalty_accounts la
-- loyalty_tiers lt