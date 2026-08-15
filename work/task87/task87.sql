SELECT
    i.store_id,
    COUNT(*) AS total_rentals,
    COUNT(DISTINCT r.customer_id) AS unique_customers
FROM rental r
JOIN inventory i ON r.inventory_id = i.inventory_id
GROUP BY i.store_id
ORDER BY total_rentals DESC;