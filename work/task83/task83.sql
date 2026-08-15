WITH film_rentals AS (
    SELECT
        f.film_id,
        f.title,
        f.release_year,
        COUNT(*) AS rentals_count
    FROM rental r
    JOIN inventory i ON r.inventory_id = i.inventory_id
    JOIN film f ON i.film_id = f.film_id
    GROUP BY f.film_id, f.title, f.release_year
),
ranked_films AS (
    SELECT
        title,
        release_year,
        rentals_count,
        ROW_NUMBER() OVER (
            PARTITION BY release_year
            ORDER BY rentals_count DESC, title ASC
        ) AS rank_in_year
    FROM film_rentals
)
SELECT
    title,
    release_year,
    rentals_count,
    rank_in_year
FROM ranked_films
WHERE rank_in_year <= 3
ORDER BY release_year, rank_in_year;