-- +goose Up
CREATE TABLE milk
(
    date text not null ,
    liters timestamp,
    price double precision,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE milk;