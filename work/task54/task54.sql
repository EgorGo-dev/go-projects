SELECT
    id,
    sale_id,
    product_id,
    quantity,
    unit_price,
    discount_amount,
    ROUND(quantity * unit_price, 1) AS gross_total,
    ROUND(quantity * unit_price - discount_amount, 1) AS net_total,
    ROUND(discount_amount * 100.0 / (quantity * unit_price), 1) AS discount_percent,
    'SALE ' || sale_id || ': ' || ROUND(discount_amount * 100.0 / (quantity * unit_price), 1) || '%' AS label
FROM
    sale_items
WHERE
    (quantity * unit_price) BETWEEN 600 AND 1300
    AND discount_amount >= 0.15 * (quantity * unit_price)
    AND unit_price IS NOT NULL
    AND quantity > 0
ORDER BY
    discount_percent DESC,
    gross_total DESC
LIMIT 15;