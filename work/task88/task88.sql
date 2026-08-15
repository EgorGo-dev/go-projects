WITH category_revenue AS (
    SELECT
        i.store_id,
        c.name AS category_name,
        SUM(p.amount) AS revenue
    FROM payment p
    JOIN rental r ON p.rental_id = r.rental_id
    JOIN inventory i ON r.inventory_id = i.inventory_id
    JOIN film f ON i.film_id = f.film_id
    JOIN film_category fc ON f.film_id = fc.film_id
    JOIN category c ON fc.category_id = c.category_id
    GROUP BY i.store_id, c.category_id, c.name
)
SELECT
    store_id,
    category_name,
    ROUND(revenue, 2) AS revenue,
    ROUND(
        100.0 * revenue / SUM(revenue) OVER (PARTITION BY store_id),
        2
    ) AS share_percent
FROM category_revenue
ORDER BY store_id, revenue DESC;