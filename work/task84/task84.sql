SELECT
    film.title,
    COUNT(*) AS rentals_count
FROM rental
JOIN inventory ON rental.inventory_id = inventory.inventory_id
JOIN film ON inventory.film_id = film.film_id
GROUP BY film.film_id, film.title
ORDER BY rentals_count DESC, film.title DESC
LIMIT 10;