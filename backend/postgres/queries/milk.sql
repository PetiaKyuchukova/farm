-- name: GetMilkInTimeframe :many
SELECT * FROM milk
where date BETWEEN $1 AND $2; ;

-- name: UpsertMilk :exec
INSERT INTO milk(date,liters,price) VALUES (@date, @liters, @price)
    ON CONFLICT(date)
    DO UPDATE SET
         date = @date,
        liters = @liters,
        price = @price;


