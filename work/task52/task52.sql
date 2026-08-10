SELECT
    (SELECT ROUND(SUM(applied_amount), 2) FROM sale_coupons) AS coupon_discount_total,
    (SELECT ROUND(SUM(quantity * unit_price - discount_amount), 2) FROM sale_items) AS sales_net_total,
    ROUND(
        (SELECT SUM(applied_amount) FROM sale_coupons) * 100.0 /
        (SELECT SUM(quantity * unit_price - discount_amount) FROM sale_items), 2
    ) AS coupon_share_percent;