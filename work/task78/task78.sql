SELECT
	id AS product_id,
	category_id,
	base_price,
    DENSE_RANK() OVER (PARTITION BY category_id ORDER BY base_price DESC) AS dense_price_rank
FROM products
ORDER BY category_id, dense_price_rank;