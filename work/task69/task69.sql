WITH tx_totals AS (
    SELECT 
        account_id,
        SUM(points_delta) AS tx_sum
    FROM loyalty_point_transactions
    GROUP BY account_id
)
SELECT 
    la.id AS account_id,
    la.points_balance,
    tt.tx_sum,
    la.points_balance - tt.tx_sum AS difference
FROM loyalty_accounts la
INNER JOIN tx_totals tt ON la.id = tt.account_id
ORDER BY la.id;