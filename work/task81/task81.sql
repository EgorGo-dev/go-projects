WITH city_customer_revenue AS (
    SELECT 
        st.city,
        s.customer_id,
        ROUND(SUM(si.quantity * si.unit_price - si.discount_amount), 2) AS total_revenue
    FROM sales s
    JOIN stores st ON s.store_id = st.id
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY st.city, s.customer_id
)
SELECT 
    city,
    customer_id,
    total_revenue,
    DENSE_RANK() OVER (PARTITION BY city ORDER BY total_revenue DESC) AS city_rank
FROM city_customer_revenue
ORDER BY city, city_rank;