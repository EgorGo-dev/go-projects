SELECT 
    c.full_name AS customer_name,
    lt.name AS tier_name,
    SUM(CASE WHEN lpt.tx_type = 'EARN' THEN lpt.points_delta ELSE 0 END) AS total_earned_points,
    SUM(CASE WHEN lpt.tx_type = 'SPEND' THEN ABS(lpt.points_delta) ELSE 0 END) AS total_spent_points
FROM loyalty_point_transactions lpt
INNER JOIN loyalty_accounts la ON lpt.account_id = la.id
INNER JOIN customers c ON la.customer_id = c.id
INNER JOIN loyalty_tiers lt ON la.tier_id = lt.id
WHERE lpt.tx_type IN ('EARN', 'SPEND')
GROUP BY c.id, c.full_name, lt.name
ORDER BY total_earned_points DESC, c.id ASC;