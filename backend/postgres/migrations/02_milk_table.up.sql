-- +goose Up
CREATE TABLE milk
(
    date timestamp not null ,
    liters double precision,
    price double precision,
    PRIMARY KEY (date)
);

-- +goose Down
DROP TABLE milk;