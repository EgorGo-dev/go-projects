SELECT 
    t.id,
    t.customer_id,
    t.sale_total
FROM (
    -- Этот подзапрос вычисляет сумму для каждой продажи
    SELECT 
        s.id,
        s.customer_id,
        ROUND(SUM(si.unit_price * si.quantity - si.discount_amount), 2) AS sale_total
    FROM sales s
    JOIN sale_items si ON s.id = si.sale_id
    GROUP BY s.id, s.customer_id
) AS t
WHERE NOT EXISTS (
    -- Проверяем, нет ли у этого же клиента другой продажи с большей суммой
    SELECT 1
    FROM sales s2
    JOIN sale_items si2 ON s2.id = si2.sale_id
    WHERE s2.customer_id = t.customer_id
    GROUP BY s2.id, s2.customer_id
    HAVING 
        ROUND(SUM(si2.unit_price * si2.quantity - si2.discount_amount), 2) > t.sale_total
        OR (
            ROUND(SUM(si2.unit_price * si2.quantity - si2.discount_amount), 2) = t.sale_total 
            AND s2.id < t.id
        )
);