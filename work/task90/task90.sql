WITH monthly AS (
    SELECT
        strftime('%Y-%m', payment_date) AS month,
        ROUND(SUM(amount), 2) AS total_revenue
    FROM payment
    GROUP BY month
)
SELECT
    month,
    total_revenue,
    ROUND(
        total_revenue - LAG(total_revenue) OVER (ORDER BY month),
        2
    ) AS growth
FROM monthly
ORDER BY month;