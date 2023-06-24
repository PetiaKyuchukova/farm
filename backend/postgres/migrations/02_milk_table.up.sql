-- +goose Up
CREATE TABLE milk
(
    date text not null ,
    liters timestamp,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE milk;