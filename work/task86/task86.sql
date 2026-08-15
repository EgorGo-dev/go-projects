SELECT
    customer_id,
    COUNT(*) AS rentals_count
FROM rental
GROUP BY customer_id
ORDER BY customer_id;

SELECT
    AVG(cnt)
FROM (
    SELECT COUNT(*) AS cnt
    FROM rental
    GROUP BY customer_id
) AS t;