-- +goose Up
CREATE TABLE farmers
(
    name text NOT NULL,
    password text NOT NULL,
    email text NOT NULL,
    PRIMARY KEY (name,email)
);

-- +goose Down
DROP TABLE farmers;