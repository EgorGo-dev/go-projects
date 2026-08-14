WITH ranked_products AS (
    SELECT 
    	name AS product_name,
    	category_id,
    	base_price
    FROM products
)
SELECT 
	product_name,
	category_id,
	base_price,
	ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY base_price DESC)AS price_rank_in_category
FROM ranked_products
ORDER BY category_id, price_rank_in_category;