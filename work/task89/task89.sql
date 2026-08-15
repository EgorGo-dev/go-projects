WITH film_rentals AS (
    SELECT
        c.name AS category_name,
        f.title,
        COUNT(*) AS rentals_count
    FROM rental r
    JOIN inventory i ON r.inventory_id = i.inventory_id
    JOIN film f ON i.film_id = f.film_id
    JOIN film_category fc ON f.film_id = fc.film_id
    JOIN category c ON fc.category_id = c.category_id
    GROUP BY c.name, f.film_id, f.title
),
ranked AS (
    SELECT
        category_name,
        title,
        rentals_count,
        ROW_NUMBER() OVER (
            PARTITION BY category_name 
            ORDER BY rentals_count DESC, title ASC
        ) AS rank_in_category
    FROM film_rentals
)
SELECT
    category_name,
    title,
    rentals_count,
    rank_in_category
FROM ranked
WHERE rank_in_category <= 3
ORDER BY category_name, rank_in_category;