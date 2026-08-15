WITH client_category_rentals AS (
    SELECT 
        r.customer_id,
        c.category_id,
        c.name AS category_name,
        COUNT(*) AS rentals_in_category
    FROM rental r
    JOIN inventory i ON r.inventory_id = i.inventory_id
    JOIN film f ON i.film_id = f.film_id
    JOIN film_category fc ON f.film_id = fc.film_id
    JOIN category c ON fc.category_id = c.category_id
    GROUP BY r.customer_id, c.category_id, c.name
),
favorite_category AS (
    SELECT 
        customer_id,
        category_name,
        rentals_in_category
    FROM (
        SELECT 
            customer_id,
            category_name,
            rentals_in_category,
            ROW_NUMBER() OVER (
                PARTITION BY customer_id 
                ORDER BY rentals_in_category DESC, category_name ASC
            ) AS rn
        FROM client_category_rentals
    ) ranked
    WHERE rn = 1
),
payment_stats AS (
    SELECT 
        customer_id,
        ROUND(SUM(amount), 2) AS total_payments,
        ROUND(AVG(amount), 2) AS avg_payment
    FROM payment
    GROUP BY customer_id
),
total_rentals AS (
    SELECT 
        customer_id,
        COUNT(*) AS total_rentals_count
    FROM rental
    GROUP BY customer_id
)
SELECT
    c.customer_id,
    c.first_name,
    c.last_name,
    fc.category_name AS favorite_category,
    fc.rentals_in_category AS rentals_in_favorite,
    ps.total_payments,
    ps.avg_payment,
    RANK() OVER (ORDER BY ps.total_payments DESC) AS rank_by_payments
FROM customer c
JOIN favorite_category fc ON c.customer_id = fc.customer_id
JOIN payment_stats ps ON c.customer_id = ps.customer_id
JOIN total_rentals tr ON c.customer_id = tr.customer_id
WHERE tr.total_rentals_count > 10
ORDER BY rank_by_payments;